package internal

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
)

func Fold[T, B any](iter iterator.Iterator[T], init B, accum delegate.Function2[T, B, B]) B {
	return init
}
