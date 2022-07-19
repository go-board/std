package stream

import (
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

// Stream is a lazy sequence of values of type T.
// 	It wraps an iterator.Iterator[T] and provides a lazy access to its elements.
// 	It also provides a way to chain operations on the stream.
type Stream[T any] interface {
	iterator.Iterable[T]
	// All returns true if all elements of the stream satisfy the predicate.
	All(predicate func(T) bool) bool
	// Any returns true if any element of the stream satisfies the predicate.
	Any(predicate func(T) bool) bool
	// Once returns true if the predicate is true for exactly one element of the stream satisfy the predicate.
	Once(predicate func(T) bool) bool
	// None returns true if none of the elements of the stream satisfy the predicate.
	None(predicate func(T) bool) bool
	// Chain returns a new stream that is the result of chaining the current stream with the given iterable.
	Chain(o iterator.Iterable[T]) Stream[T]
	// ForEach calls the given function for each element of the stream.
	ForEach(consumer func(T))
	// Map returns a new stream that is the result of applying the given function to each element of the stream.
	Map(transformer func(T) T) Stream[T]
	// Filter returns a new stream that is the result of applying the given predicate to each element of the stream.
	Filter(predicate func(T) bool) Stream[T]
	// Reduce returns an optional value of type T that is the result of applying the given function to each element of the stream.
	Reduce(reduce func(T, T) T) optional.Optional[T]
	// Fold returns a value of type T that is the result of applying the given function to each element of the stream with the given initial value.
	Fold(initial T, accumulator func(T, T) T) T
	// Flatten returns a new stream that is the result of flattening each element in the current stream.
	Flatten(flatten func(T) iterator.Iterator[T]) Stream[T]
	// Inspect does something with each element of the stream. Different from ForEach, this will return the stream itself.
	Inspect(inspect func(T)) Stream[T]
	// MaxBy returns an optional value of type T that is the maximum element of the stream according to the given comparator.
	MaxBy(ord cmp.OrdFunc[T]) optional.Optional[T]
	// MinBy returns an optional value of type T that is the minimum element of the stream according to the given comparator.
	MinBy(ord cmp.OrdFunc[T]) optional.Optional[T]
	// EqualBy returns true if the two streams are equal according to the given comparator.
	EqualBy(o iterator.Iterable[T], ord cmp.EqFunc[T]) bool
	// Collect returns a slice of all the elements of the stream.
	Collect() []T
	// Nth returns an optional value of type T that is the nth element of the stream.
	Nth(n uint) optional.Optional[T]
	// Take returns a new stream that is the result of taking the given number of elements from the current stream.
	Take(n uint) Stream[T]
	// Skip returns a new stream that is the result of skipping the given number of elements from the current stream.
	Skip(n uint) Stream[T]
	// TakeWhile returns a new stream that is the result of taking elements from the current stream while the predicate is true.
	TakeWhile(predicate func(T) bool) Stream[T]
	// SkipWhile returns a new stream that is the result of skipping elements from the current stream while the predicate is true.
	SkipWhile(predicate func(T) bool) Stream[T]
	// StepBy returns a new stream that is the result of stepping by the given number of elements from the current stream.
	StepBy(step uint) Stream[T]
	// First returns an optional value of type T that is the first element of the stream.
	First() optional.Optional[T]
	// Last returns an optional value of type T that is the last element of the stream.
	Last() optional.Optional[T]
	// Advancing returns a new stream that is the result of advancing the current stream by the given number of elements.
	Advancing(step uint) Stream[T]
	// IsSorted returns true if the stream is sorted according to the given comparator.
	IsSorted(ord cmp.OrdFunc[T]) bool
}
