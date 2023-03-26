package ops

import (
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/iterator"
)

func EqualBy[
	T any,
	I iterator.Iterator[T],
	IA iterator.Iterable[T],
	F cmp.EqFunc[T],
](iter I, iterable IA, eq F) bool {
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
