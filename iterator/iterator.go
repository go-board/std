package iterator

import (
	"github.com/go-board/std/optional"
)

// Iterator is the interface that wraps the basic methods for iterating over a collection.
type Iterator[T any] interface {
	Next() optional.Optional[T]
}

// DoubleEndedIterator is the interface that wraps the basic methods for iterating over a collection in both directions.
type DoubleEndedIterator[T any] interface {
	Iterator[T]
	Prev() optional.Optional[T]
}

// SizedIterator is the interface that wraps the basic methods for iterating over a collection that also provides a size.
type SizedIterator[T any] interface {
	Iterator[T]
	SizeHint() (uint, optional.Optional[uint])
}
