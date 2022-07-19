package ops

import "github.com/go-board/std/iterator"

func Collect[T any](iter iterator.Iterator[T]) []T {
	slice := make([]T, 0)
	for s := iter.Next(); s.IsSome(); s = iter.Next() {
		slice = append(slice, s.Value())
	}
	return slice
}
