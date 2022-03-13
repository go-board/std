package internal

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type chunkIter[T any, TS ~[]T] struct {
	iter      iterator.Iterator[T]
	chunkSize uint
}

func (self *chunkIter[T, TS]) Next() optional.Optional[TS] {
	cnt := uint(0)
	chunk := make(TS, 0, self.chunkSize)

	for s := self.iter.Next(); s.IsSome(); s = self.iter.Next() {
		cnt++
		chunk = append(chunk, s.Value())
		if cnt == self.chunkSize {
			return optional.Some(chunk)
		}
	}
	if len(chunk) != 0 {
		return optional.Some(chunk)
	}
	return optional.None[TS]()
}

func Chunk[T any](iter iterator.Iterator[T], n uint) iterator.Iterator[[]T] {
	return &chunkIter[T, []T]{iter: iter, chunkSize: n}
}
