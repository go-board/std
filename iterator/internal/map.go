package internal

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type mapIter[T, U any] struct {
	iter        iterator.Iterator[T]
	transformer delegate.Transform[T, U]
}

func (i *mapIter[T, U]) Next() optional.Optional[U] {
	e := i.iter.Next()
	if e.IsSome() {
		return optional.Some(i.transformer(e.Value()))
	}
	return optional.None[U]()
}

func Map[T, U any](iter iterator.Iterator[T], transformer delegate.Transform[T, U]) iterator.Iterator[U] {
	return &mapIter[T, U]{iter, transformer}
}
