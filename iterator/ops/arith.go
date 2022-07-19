package ops

import (
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/cond"
	"github.com/go-board/std/core"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

func MaxBy[T any](iter iterator.Iterator[T], ord cmp.OrdFunc[T]) optional.Optional[T] {
	return Reduce(iter, func(a, b T) T { return cond.Ternary(ord(a, b).IsGe(), a, b) })
}

func MinBy[T any](iter iterator.Iterator[T], ord cmp.OrdFunc[T]) optional.Optional[T] {
	return Reduce(iter, func(a, b T) T { return cond.Ternary(ord(a, b).IsLe(), a, b) })
}

func SumBy[T any](iter iterator.Iterator[T], sum func(T, T) T) optional.Optional[T] {
	return Reduce(iter, sum)
}

func ProductBy[T any](iter iterator.Iterator[T], product func(T, T) T) optional.Optional[T] {
	return Reduce(iter, product)
}

func ord[T core.Ordered](lhs, rhs T) cmp.Ordering {
	if lhs > rhs {
		return cmp.Greater
	} else if lhs < rhs {
		return cmp.Less
	}
	return cmp.Equal
}

func Max[T core.Ordered](iter iterator.Iterator[T]) optional.Optional[T] {
	return MaxBy(iter, ord[T])
}

func Min[T core.Ordered](iter iterator.Iterator[T]) optional.Optional[T] {
	return MinBy(iter, ord[T])
}

func Sum[T core.Number](iter iterator.Iterator[T]) optional.Optional[T] {
	return SumBy(iter, func(t1 T, t2 T) T { return t1 + t2 })
}

func Product[T core.Number](iter iterator.Iterator[T]) optional.Optional[T] {
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
