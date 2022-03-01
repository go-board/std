package internal

import (
	"github.com/go-board/std/clone"
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type cloneIter[T any] struct {
	iter  iterator.Iterator[T]
	clone delegate.Function1[optional.Optional[T], optional.Optional[T]]
}

func (self *cloneIter[T]) Next() optional.Optional[T] {
	return self.clone(self.iter.Next())
}

func CloneBy[T any](
	iter iterator.Iterator[T],
	clone delegate.Function1[optional.Optional[T], optional.Optional[T]],
) iterator.Iterator[T] {
	return &cloneIter[T]{iter, clone}
}

func Clone[T clone.Cloneable[T]](iter iterator.Iterator[T]) iterator.Iterator[T] {
	return CloneBy(iter, optional.Clone[T])
}
