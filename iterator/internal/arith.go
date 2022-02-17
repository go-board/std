package internal

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

func MaxBy[T any](iter iterator.Iterator[T], ord delegate.Ord[T]) optional.Optional[T] {
	return Reduce(iter, func(a, b T) T {
		if ord(a, b) > 0 {
			return a
		}
		return b
	})
}

func MinBy[T any](iter iterator.Iterator[T], ord delegate.Ord[T]) optional.Optional[T] {
	return Reduce(iter, func(a, b T) T {
		if ord(a, b) < 0 {
			return a
		}
		return b
	})
}

func SumBy[T any](iter iterator.Iterator[T], sum delegate.Function2[T, T, T]) optional.Optional[T] {
	return Reduce(iter, sum)
}

func ProductBy[T any](iter iterator.Iterator[T], product delegate.Function2[T, T, T]) optional.Optional[T] {
	return Reduce(iter, product)
}

type Numbric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

func Max[T Numbric](iter iterator.Iterator[T]) optional.Optional[T] {
	return MaxBy(iter, func(t1, t2 T) int {
		if t1 > t2 {
			return 1
		} else if t1 < t2 {
			return -1
		}
		return 0
	})
}

func Min[T Numbric](iter iterator.Iterator[T]) optional.Optional[T] {
	return MinBy(iter, func(t1, t2 T) int {
		if t1 > t2 {
			return 1
		} else if t1 < t2 {
			return -1
		}
		return 0
	})
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
