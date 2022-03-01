package adapters

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type sliceIter[T any] struct {
	elements  []T
	iterIndex int
}

func (self *sliceIter[T]) Next() optional.Optional[T] {
	if self.iterIndex < len(self.elements) {
		self.iterIndex++
		return optional.Some(self.elements[self.iterIndex-1])
	}
	return optional.None[T]()
}

func OfSlice[T any](eles ...T) iterator.Iterator[T] {
	return &sliceIter[T]{elements: eles}
}