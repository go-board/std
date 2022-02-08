package internal

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type flattenIter[T any] struct {
	iter    iterator.Iterator[T]
	flatten delegate.Transform[T, iterator.Iterator[T]]
}

func (i *flattenIter[T]) Next() optional.Optional[T] {
	return optional.None[T]()
}

func Flatten[T any](iter iterator.Iterator[T], flatten delegate.Transform[T, iterator.Iterator[T]]) iterator.Iterator[T] {
	return &flattenIter[T]{iter: iter, flatten: flatten}
}
