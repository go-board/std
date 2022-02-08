package internal

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
)

func All[T any](iter iterator.Iterator[T], predicate delegate.Predicate[T]) bool {
	for s := iter.Next(); s.IsSome(); s = iter.Next() {
		if !predicate(s.Value()) {
			return false
		}
	}
	return true
}

func Any[T any](iter iterator.Iterator[T], predicate delegate.Predicate[T]) bool {
	for s := iter.Next(); s.IsSome(); s = iter.Next() {
		if predicate(s.Value()) {
			return true
		}
	}
	return false
}

func None[T any](iter iterator.Iterator[T], predicate delegate.Predicate[T]) bool {
	for s := iter.Next(); s.IsSome(); s = iter.Next() {
		if predicate(s.Value()) {
			return false
		}
	}
	return true
}
