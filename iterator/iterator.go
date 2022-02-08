package iterator

import (
	"github.com/go-board/std/optional"
)

type Iterator[T any] interface {
	Next() optional.Optional[T]
}

type DoubleEndedIterator[T any] interface {
	Iterator[T]
	Prev() optional.Optional[T]
}

type SizedIterator[T any] interface {
	Iterator[T]
	SizeHint() (uint, optional.Optional[uint])
}
