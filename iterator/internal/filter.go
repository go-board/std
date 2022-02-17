package internal

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type filterIter[T any] struct {
	iter   iterator.Iterator[T]
	filter delegate.Predicate[T]
}

func (i *filterIter[T]) Next() optional.Optional[T] {
	for s := i.iter.Next(); s.IsSome(); s = i.iter.Next() {
		if i.filter(s.Value()) {
			return s
		}
	}
	return optional.None[T]()
}

func Filter[T any](iter iterator.Iterator[T], filter delegate.Predicate[T]) iterator.Iterator[T] {
	return &filterIter[T]{iter, filter}
}

type mapFilterIter[T, U any] struct {
	iter      iterator.Iterator[T]
	mapFilter func(T) optional.Optional[U]
}

func (i *mapFilterIter[T, U]) Next() optional.Optional[U] {
	for s := i.iter.Next(); s.IsSome(); s = i.iter.Next() {
		if m := i.mapFilter(s.Value()); m.IsSome() {
			return m
		}
	}
	return optional.None[U]()
}

func MapFilter[T, U any](iter iterator.Iterator[T], mapFilter func(T) optional.Optional[U]) iterator.Iterator[U] {
	return &mapFilterIter[T, U]{iter, mapFilter}
}

type takeIter[T any] struct {
	iter      iterator.Iterator[T]
	n         uint
	iterIndex uint
}

func (i *takeIter[T]) Next() optional.Optional[T] {
	if i.iterIndex < i.n {
		i.iterIndex++
		return i.iter.Next()
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

func (i *skipIter[T]) Next() optional.Optional[T] {
	for i.iterIndex < i.n {
		i.iterIndex++
		i.iter.Next()
	}
	return i.iter.Next()
}

func Skip[T any](iter iterator.Iterator[T], n uint) iterator.Iterator[T] {
	return &skipIter[T]{iter, n, 0}
}

type takeWhileIter[T any] struct {
	iter      iterator.Iterator[T]
	predicate delegate.Predicate[T]
}

func (i *takeWhileIter[T]) Next() optional.Optional[T] {
	for s := i.iter.Next(); s.IsSome(); s = i.iter.Next() {
		if i.predicate(s.Value()) {
			return s
		}
	}
	return optional.None[T]()
}

func TakeWhile[T any](iter iterator.Iterator[T], predicate delegate.Predicate[T]) iterator.Iterator[T] {
	return &takeWhileIter[T]{iter, predicate}
}

type skipWhileIter[T any] struct {
	iter      iterator.Iterator[T]
	predicate delegate.Predicate[T]
}

func (i *skipWhileIter[T]) Next() optional.Optional[T] {
	for s := i.iter.Next(); s.IsSome(); s = i.iter.Next() {
		if !i.predicate(s.Value()) {
			return s
		}
	}
	return optional.None[T]()
}
func SkipWhile[T any](iter iterator.Iterator[T], predicate delegate.Predicate[T]) iterator.Iterator[T] {
	return &skipWhileIter[T]{iter, predicate}
}

type stepByIter[T any] struct {
	iter      iterator.Iterator[T]
	n         uint
	iterIndex uint
}

func (i *stepByIter[T]) Next() optional.Optional[T] {
	for s := i.iter.Next(); s.IsSome(); s = i.iter.Next() {
		i.iterIndex++
		if i.iterIndex%i.n == 0 {
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
