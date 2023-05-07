package adapters

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

// IntIter returns an iterator that iterates over the integers from 0 to upperBound.
func IntIter(upperBound int) iterator.Iter[int] {
	return IntRanger(0, upperBound)
}

func IntRanger(lowerBound, upperBound int) iterator.Iter[int] {
	if lowerBound > upperBound {
		panic("lowerBound must less than upperBound")
	}
	current := lowerBound - 1
	return iterator.IterFunc[int](func() optional.Optional[int] {
		if current < upperBound {
			current++
			return optional.Some(current)
		}
		return optional.None[int]()
	})
}
