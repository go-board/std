package ops

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type cloneIter[T any] struct {
	iter  iterator.Iterator[T]
	clone func(optional.Optional[T]) optional.Optional[T]
}

func (self *cloneIter[T]) Next() optional.Optional[T] {
	return self.clone(self.iter.Next())
}

func CloneBy[T any](
	iter iterator.Iterator[T],
	clone func(optional.Optional[T]) optional.Optional[T],
) iterator.Iterator[T] {
	return &cloneIter[T]{iter, clone}
}
