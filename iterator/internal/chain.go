package internal

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type chainIter[T any] struct {
	iter  iterator.Iterator[T]
	chain iterator.Iterator[T]
}

func (self *chainIter[T]) Next() optional.Optional[T] {
	return self.iter.Next().Or(self.chain.Next())
}

func Chain[T any, I iterator.Iterator[T], IA iterator.Iterable[T]](
	iter I,
	iterable IA,
) iterator.Iterator[T] {
	return &chainIter[T]{iter, iterable.Iter()}
}
