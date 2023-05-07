package ops

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/operator"
)

// EqualBy
func EqualBy[T any](
	iter iterator.Iter[T],
	iterable iterator.Iterable[T],
	eq func(T, T) bool,
) bool {
	iterB := iterable.Iter()
	for {
		a := iter.Next()
		b := iterB.Next()
		if a.IsNone() && b.IsNone() {
			return true
		}

		if a.IsSome() && b.IsSome() {
			if !eq(a.Value(), b.Value()) {
				return false
			}
			continue
		}
		return false
	}
}

func Equal[T comparable](
	iter iterator.Iter[T],
	iterable iterator.Iterable[T],
) bool {
	return EqualBy(iter, iterable, operator.Eq[T])
}
