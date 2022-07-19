package ops

import (
	"github.com/go-board/std/iterator"
)

func Fold[T, B any](iter iterator.Iterator[T], init B, accum func(T, B) B) B {
	for {
		a := iter.Next()
		if a.IsNone() {
			return init
		}
		init = accum(a.Value(), init)
	}
}
