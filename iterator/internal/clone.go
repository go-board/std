package internal

import "github.com/go-board/std/iterator"

func Clone[T any](iter iterator.Iterator[T]) iterator.Iterator[T] {
	return iter
}
