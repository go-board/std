package adapters

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type intIter struct {
	upperBound int
	current    int
}

// IntIterator returns an iterator that iterates over the integers from 0 to upperBound.
func IntIterator(upperBound int) iterator.Iterator[int] {
	return &intIter{upperBound: upperBound}
}

func (self *intIter) Next() optional.Optional[int] {
	if self.current < self.upperBound {
		self.current++
		return optional.Some(self.current)
	}
	return optional.None[int]()
}

type uintIter struct {
	upperBound uint
	current    uint
}

// UintIterator returns an iterator that iterates over the unsigned integers from 0 to upperBound.
func UintIterator(upperBound uint) iterator.Iterator[uint] {
	return &uintIter{upperBound: upperBound}
}

func (self *uintIter) Next() optional.Optional[uint] {
	if self.current < self.upperBound {
		self.current++
		return optional.Some(self.current)
	}
	return optional.None[uint]()
}
