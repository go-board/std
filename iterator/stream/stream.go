package stream

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type Stream[T any] interface {
	iterator.Iterable[T]
	All(predicate delegate.Predicate[T]) bool
	Any(predicate delegate.Predicate[T]) bool
	Once(predicate delegate.Predicate[T]) bool
	None(predicate delegate.Predicate[T]) bool
	Chain(o iterator.Iterable[T]) Stream[T]
	Map(transformer delegate.Transform[T, T]) Stream[T]
	Filter(predicate delegate.Predicate[T]) Stream[T]
	Reduce(reduce delegate.Function2[T, T, T]) optional.Optional[T]
	Flatten(flatten delegate.Transform[T, iterator.Iterator[T]]) Stream[T]
	Inspect(inspect delegate.Consumer1[T]) Stream[T]
	MaxBy(ord delegate.Ord[T]) optional.Optional[T]
	MinBy(ord delegate.Ord[T]) optional.Optional[T]
	Equal(o iterator.Iterable[T], ord delegate.Equal[T]) bool
	Collect() []T
	Nth(n uint) optional.Optional[T]
	Take(n uint) Stream[T]
	Skip(n uint) Stream[T]
	TakeWhile(predicate delegate.Predicate[T]) Stream[T]
	SkipWhile(predicate delegate.Predicate[T]) Stream[T]
	ForEach(consumer delegate.Consumer1[T])
	StepBy(step uint) Stream[T]
	Last() optional.Optional[T]
	Advancing(step uint) Stream[T]
	IsSorted(ord delegate.Ord[T]) bool
}
