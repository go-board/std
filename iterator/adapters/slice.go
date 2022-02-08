package adapters

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type sliceIter[T any] struct {
	elements  []T
	iterIndex int
}

func (i *sliceIter[T]) Next() optional.Optional[T] {
	if i.iterIndex < len(i.elements) {
		i.iterIndex++
		return optional.Some(i.elements[i.iterIndex-1])
	}
	return optional.None[T]()
}

func OfSlice[T any](eles ...T) iterator.Iterator[T] {
	return &sliceIter[T]{elements: eles}
}
