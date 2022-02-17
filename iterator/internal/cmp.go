package internal

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
)

func EqualBy[
	T any,
	I iterator.Iterator[T],
	IA iterator.Iterable[T],
	F delegate.Equal[T],
](iter I, iterable IA, eq F) bool {
	return false
}

func IsSorted[T any](iter iterator.Iterator[T], ord delegate.Ord[T]) bool {
	return false
}

func CmpBy[T any](iter iterator.Iterator[T], iterable iterator.Iterable[T], ord delegate.Ord[T]) bool {
	return false
}
