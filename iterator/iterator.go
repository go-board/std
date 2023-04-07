package iterator

import (
	"github.com/go-board/std/optional"
)

// Iterator is the interface that wraps the basic methods for iterating over a collection.
//
// 💣 Important: Iterators usually contains state, so the behavior of reuse is undefined.
//
// For a short variant, use [Iter] instead.
// In the future, [Iterator] will be deprecated and removed in favor of [Iter].
type Iterator[T any] interface {
	// Next advances the iterator and returns the next value.
	//
	// Returns [optional.None] when iteration is finished.
	//
	// Example:
	// 	iter := slices.Iter(1, 2, 3)
	//  iter.Next() // Some(1)
	//  iter.Next() // Some(2)
	//  iter.Next() // Some(3)
	//  iter.Next() // None
	Next() optional.Optional[T]
}

// DoubleEndedIterator is the interface that wraps the basic methods for iterating over a collection in both directions.
type DoubleEndedIterator[T any] interface {
	Iterator[T]
	Prev() optional.Optional[T]
}

// Iter is short alias for [Iterator]
type Iter[T any] interface {
	Iterator[T]
}

// IterFunc is a function type that implements [Iter].
type IterFunc[T any] func() optional.Optional[T]

func (fn IterFunc[T]) Next() optional.Optional[T] { return fn() }

var _ Iter[any] = (IterFunc[any])(nil)
