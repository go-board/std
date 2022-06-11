package stream

import (
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type Stream[T any] interface {
	iterator.Iterable[T]
	All(predicate func(T) bool) bool
	Any(predicate func(T) bool) bool
	Once(predicate func(T) bool) bool
	None(predicate func(T) bool) bool
	Chain(o iterator.Iterable[T]) Stream[T]
	Map(transformer func(T) T) Stream[T]
	Filter(predicate func(T) bool) Stream[T]
	Reduce(reduce func(T, T) T) optional.Optional[T]
	Flatten(flatten func(T) iterator.Iterator[T]) Stream[T]
	Inspect(inspect func(T)) Stream[T]
	MaxBy(ord cmp.OrdFunc[T]) optional.Optional[T]
	MinBy(ord cmp.OrdFunc[T]) optional.Optional[T]
	Equal(o iterator.Iterable[T], ord cmp.EqFunc[T]) bool
	Collect() []T
	Nth(n uint) optional.Optional[T]
	Take(n uint) Stream[T]
	Skip(n uint) Stream[T]
	TakeWhile(predicate func(T) bool) Stream[T]
	SkipWhile(predicate func(T) bool) Stream[T]
	ForEach(consumer func(T))
	StepBy(step uint) Stream[T]
	Last() optional.Optional[T]
	Advancing(step uint) Stream[T]
	IsSorted(ord cmp.OrdFunc[T]) bool
}
