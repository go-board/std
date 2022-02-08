package internal

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type inspectIter[T any] struct {
	iter    iterator.Iterator[T]
	inspect delegate.Consumer1[T]
}

func (i *inspectIter[T]) Next() optional.Optional[T] {
	o := i.iter.Next()
	if o.IsSome() {
		i.inspect(o.Value())
	}
	return o
}

func Inspect[T any](iter iterator.Iterator[T], inspect delegate.Consumer1[T]) iterator.Iterator[T] {
	return &inspectIter[T]{iter, inspect}
}
