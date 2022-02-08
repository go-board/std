package internal

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
)

func ForEach[T any](iter iterator.Iterator[T], consumer delegate.Consumer1[T]) {
	for e := iter.Next(); e.IsSome(); e = iter.Next() {
		consumer(e.Value())
	}
}
