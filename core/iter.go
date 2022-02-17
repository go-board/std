package core

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type intIter struct {
	upperBound Int
	current    Int
}

func (self *intIter) Next() optional.Optional[Int] {
	if self.current < self.upperBound {
		self.current++
		return optional.Some(self.current)
	}
	return optional.None[Int]()
}

func (self Int) Iter() iterator.Iterator[Int] {
	return &intIter{upperBound: self}
}

var _ iterator.Iterable[Int] = Int(0)

type uintIter struct {
	upperBound Uint
	current    Uint
}

func (self *uintIter) Next() optional.Optional[Uint] {
	if self.current < self.upperBound {
		self.current++
		return optional.Some(self.current)
	}
	return optional.None[Uint]()
}

func (self Uint) Iter() iterator.Iterator[Uint] {
	return &uintIter{upperBound: self}
}

var _ iterator.Iterable[Uint] = Uint(0)
