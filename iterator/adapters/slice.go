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
	self.iterIndex++
	if self.iterIndex < len(self.elements) {
		return optional.Some(self.elements[self.iterIndex])
	}
	return optional.None[T]()
}

func (self *sliceIter[T]) SizeHint() (uint, optional.Optional[uint]) {
	return 0, optional.Some(uint(len(self.elements)))
}

func OfSlice[T any](elems ...T) iterator.Iterator[T] {
	return &sliceIter[T]{elements: elems, iterIndex: -1}
}
