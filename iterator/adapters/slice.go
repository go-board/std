package adapters

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/slices"
)

func OfSlice[T any](elems ...T) iterator.Iterator[T] {
	return slices.Iter(elems)
}
