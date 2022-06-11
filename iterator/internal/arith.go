package internal

import (
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

func MaxBy[T any](iter iterator.Iterator[T], ord cmp.OrdFunc[T]) optional.Optional[T] {
	return Reduce(iter, func(a, b T) T {
		if ord(a, b).IsLe() {
			return a
		}
		return b
	})
}

func MinBy[T any](iter iterator.Iterator[T], ord cmp.OrdFunc[T]) optional.Optional[T] {
	return Reduce(iter, func(a, b T) T {
		if ord(a, b).IsLe() {
			return a
		}
		return b
	})
}

func SumBy[T any](iter iterator.Iterator[T], sum func(T, T) T) optional.Optional[T] {
	return Reduce(iter, sum)
}

func ProductBy[T any](iter iterator.Iterator[T], product func(T, T) T) optional.Optional[T] {
	return Reduce(iter, product)
}

type Numbric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

func ord[T Numbric](lhs, rhs T) cmp.Ordering {
	if lhs > rhs {
		return cmp.Greater
	} else if lhs < rhs {
		return cmp.Less
	}
	return cmp.Equal
}

func Max[T Numbric](iter iterator.Iterator[T]) optional.Optional[T] {
	return MaxBy(iter, ord[T])
}

func Min[T Numbric](iter iterator.Iterator[T]) optional.Optional[T] {
	return MinBy(iter, ord[T])
}

func Sum[T Numbric](iter iterator.Iterator[T]) optional.Optional[T] {
	return SumBy(iter, func(t1 T, t2 T) T { return t1 + t2 })
}

func Product[T Numbric](iter iterator.Iterator[T]) optional.Optional[T] {
	return ProductBy(iter, func(t1 T, t2 T) T { return t1 * t2 })
}

func Nth[T any](iter iterator.Iterator[T], n uint) optional.Optional[T] {
	iterIndex := uint(0)
	for s := iter.Next(); s.IsSome(); s = iter.Next() {
		iterIndex++
		if iterIndex == n {
			return s
		}
	}
	return optional.None[T]()
}
