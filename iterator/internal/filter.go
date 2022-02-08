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
	for {
		e := i.iter.Next()
		if e.IsSome() {
			if i.filter(e.Value()) {
				return e
			} else {
				continue
			}
		} else {
			return optional.None[T]()
		}
	}
}

func Filter[T any](iter iterator.Iterator[T], filter delegate.Predicate[T]) iterator.Iterator[T] {
	return &filterIter[T]{iter, filter}
}

type mapFilterIter[T, U any] struct {
	iter      iterator.Iterator[T]
	mapFilter func(T) optional.Optional[U]
}

func (i *mapFilterIter[T, U]) Next() optional.Optional[U] {
	for {
		e := i.iter.Next()
		if e.IsSome() {
			m := i.mapFilter(e.Value())
			if m.IsSome() {
				return m
			} else {
				continue
			}
		} else {
			return optional.None[U]()
		}
	}
}

func MapFilter[T, U any](iter iterator.Iterator[T], mapFilter func(T) optional.Optional[U]) iterator.Iterator[U] {
	return &mapFilterIter[T, U]{iter, mapFilter}
}

func Take[T any](iter iterator.Iterator[T], n uint) iterator.Iterator[T] { return iter }
func Skip[T any](iter iterator.Iterator[T], n uint) iterator.Iterator[T] { return iter }

func TakeWhile[T any](iter iterator.Iterator[T], predicate delegate.Predicate[T]) iterator.Iterator[T] {
	return iter
}
func SkipWhile[T any](iter iterator.Iterator[T], predicate delegate.Predicate[T]) iterator.Iterator[T] {
	return iter
}

func StepBy[T any](iter iterator.Iterator[T], n uint) iterator.Iterator[T] {
	return iter
}

func Advancing[T any](iter iterator.Iterator[T], n uint) iterator.Iterator[T] {
	return iter
}
