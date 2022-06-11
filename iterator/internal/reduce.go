package internal

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

func Reduce[T any](iter iterator.Iterator[T], reduce func(T, T) T) optional.Optional[T] {
	first := iter.Next()
	if first.IsNone() {
		return optional.None[T]()
	}
	return optional.Some(Fold(iter, first.Value(), reduce))
}

func Last[T any](iter iterator.Iterator[T]) optional.Optional[T] {
	return Reduce(iter, func(a, b T) T { return b })
}
