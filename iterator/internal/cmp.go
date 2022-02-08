package internal

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
)

func EqualBy[T any](iter iterator.Iterator[T], iterable iterator.Iterable[T], cmp delegate.Equal[T]) bool {
	return false
}

func Equal[T comparable](iter iterator.Iterator[T], iterable iterator.Iterable[T]) bool {
	return EqualBy(iter, iterable, func(t1, t2 T) bool { return t1 == t2 })
}

func IsSorted[T any](iter iterator.Iterator[T], cmp delegate.Comparison[T]) bool {
	return false
}
