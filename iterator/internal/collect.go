package internal

import "github.com/go-board/std/iterator"

func Collect[T any](iter iterator.Iterator[T]) []T {
	return []T{}
}
