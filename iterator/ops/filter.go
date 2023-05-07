package ops

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

func Filter[T any](it iterator.Iter[T], filter func(T) bool) iterator.Iter[T] {
	return iterator.IterFunc[T](func() optional.Optional[T] {
		for s := it.Next(); s.IsSome(); s = it.Next() {
			if filter(s.Value()) {
				return s
			}
		}
		return optional.None[T]()
	})
}

func MapFilter[T, U any](iter iterator.Iterator[T], mapFilter func(T) optional.Optional[U]) iterator.Iterator[U] {
	return iterator.IterFunc[U](func() optional.Optional[U] {
		for s := iter.Next(); s.IsSome(); s = iter.Next() {
			if m := mapFilter(s.Value()); m.IsSome() {
				return m
			}
		}
		return optional.None[U]()
	})
}

type takeIter[T any] struct {
	iter      iterator.Iterator[T]
	n         uint
	iterIndex uint
}

func (self *takeIter[T]) Next() optional.Optional[T] {
	if self.iterIndex < self.n {
		self.iterIndex++
		return self.iter.Next()
	}
	return optional.None[T]()
}

func Take[T any](iter iterator.Iterator[T], n uint) iterator.Iterator[T] {
	return &takeIter[T]{iter, n, 0}
}

type skipIter[T any] struct {
	iter      iterator.Iterator[T]
	n         uint
	iterIndex uint
}

func (self *skipIter[T]) Next() optional.Optional[T] {
	for self.iterIndex < self.n {
		self.iterIndex++
		self.iter.Next()
	}
	return self.iter.Next()
}

func Skip[T any](iter iterator.Iterator[T], n uint) iterator.Iterator[T] {
	return &skipIter[T]{iter, n, 0}
}

type takeWhileIter[T any] struct {
	iter      iterator.Iterator[T]
	predicate func(T) bool
}

func (self *takeWhileIter[T]) Next() optional.Optional[T] {
	for s := self.iter.Next(); s.IsSome(); s = self.iter.Next() {
		if self.predicate(s.Value()) {
			return s
		}
	}
	return optional.None[T]()
}

func TakeWhile[T any](iter iterator.Iterator[T], predicate func(T) bool) iterator.Iterator[T] {
	return &takeWhileIter[T]{iter, predicate}
}

type skipWhileIter[T any] struct {
	iter      iterator.Iterator[T]
	predicate func(T) bool
}

func (self *skipWhileIter[T]) Next() optional.Optional[T] {
	for s := self.iter.Next(); s.IsSome(); s = self.iter.Next() {
		if !self.predicate(s.Value()) {
			return s
		}
	}
	return optional.None[T]()
}

func SkipWhile[T any](iter iterator.Iterator[T], predicate func(T) bool) iterator.Iterator[T] {
	return &skipWhileIter[T]{iter, predicate}
}

type stepByIter[T any] struct {
	iter      iterator.Iterator[T]
	n         uint
	iterIndex uint
}

func (self *stepByIter[T]) Next() optional.Optional[T] {
	for s := self.iter.Next(); s.IsSome(); s = self.iter.Next() {
		self.iterIndex++
		if self.iterIndex%self.n == 0 {
			return s
		}
	}
	return optional.None[T]()
}

func StepBy[T any](iter iterator.Iterator[T], n uint) iterator.Iterator[T] {
	return &stepByIter[T]{iter, n, 0}
}

func Advancing[T any](iter iterator.Iterator[T], n uint) iterator.Iterator[T] {
	for i := uint(0); i < n; i++ {
		iter.Next()
	}
	return iter
}
