package stream

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/iterator/internal"
	"github.com/go-board/std/optional"
)

type streamImpl[T any] struct {
	iter iterator.Iterator[T]
}

func FromIterator[T any](iter iterator.Iterator[T]) Stream[T] {
	return &streamImpl[T]{iter: iter}
}

func FromIterable[T any](iterable iterator.Iterable[T]) Stream[T] {
	return FromIterator(iterable.Iter())
}

func (self *streamImpl[T]) Iter() iterator.Iterator[T] { return self.iter }

func (self *streamImpl[T]) All(predicate delegate.Predicate[T]) bool {
	return internal.All(self.iter, predicate)
}

func (self *streamImpl[T]) Any(predicate delegate.Predicate[T]) bool {
	return internal.Any(self.iter, predicate)
}

func (self *streamImpl[T]) Once(predicate delegate.Predicate[T]) bool {
	return internal.Once(self.iter, predicate)
}

func (self *streamImpl[T]) None(predicate delegate.Predicate[T]) bool {
	return internal.None(self.iter, predicate)
}

func (self *streamImpl[T]) Chain(o iterator.Iterable[T]) Stream[T] {
	return &streamImpl[T]{iter: internal.Chain[T](self.iter, o)}
}

func (self *streamImpl[T]) Filter(predicate delegate.Predicate[T]) Stream[T] {
	return &streamImpl[T]{iter: internal.Filter(self.iter, predicate)}
}

func (self *streamImpl[T]) Reduce(reduce delegate.Function2[T, T, T]) optional.Optional[T] {
	return internal.Reduce(self.iter, reduce)
}

func (self *streamImpl[T]) Flatten(flatten delegate.Transform[T, iterator.Iterator[T]]) Stream[T] {
	return &streamImpl[T]{iter: internal.Flatten(self.iter, flatten)}
}

func (self *streamImpl[T]) Inspect(inspect delegate.Consumer1[T]) Stream[T] {
	return &streamImpl[T]{iter: internal.Inspect(self.iter, inspect)}
}

func (self *streamImpl[T]) MaxBy(ord delegate.Ord[T]) optional.Optional[T] {
	return internal.MaxBy(self.iter, ord)
}

func (self *streamImpl[T]) MinBy(ord delegate.Ord[T]) optional.Optional[T] {
	return internal.MinBy(self.iter, ord)
}

func (self *streamImpl[T]) Equal(o iterator.Iterable[T], ord delegate.Equal[T]) bool {
	return internal.EqualBy(self.iter, o, ord)
}

func (self *streamImpl[T]) Collect() []T {
	return internal.Collect(self.iter)
}

func (self *streamImpl[T]) Nth(n uint) optional.Optional[T] {
	return internal.Nth(self.iter, n)
}

func (self *streamImpl[T]) Take(n uint) Stream[T] {
	return &streamImpl[T]{iter: internal.Take(self.iter, n)}
}

func (self *streamImpl[T]) Skip(n uint) Stream[T] {
	return &streamImpl[T]{iter: internal.Skip(self.iter, n)}
}

func (self *streamImpl[T]) TakeWhile(predicate delegate.Predicate[T]) Stream[T] {
	return &streamImpl[T]{iter: internal.TakeWhile(self.iter, predicate)}
}

func (self *streamImpl[T]) SkipWhile(predicate delegate.Predicate[T]) Stream[T] {
	return &streamImpl[T]{iter: internal.SkipWhile(self.iter, predicate)}
}

func (self *streamImpl[T]) ForEach(consumer delegate.Consumer1[T]) {
	internal.ForEach(self.iter, consumer)
}

func (self *streamImpl[T]) StepBy(step uint) Stream[T] {
	return &streamImpl[T]{iter: internal.StepBy(self.iter, step)}
}

func (self *streamImpl[T]) Last() optional.Optional[T] {
	return internal.Last(self.iter)
}

func (self *streamImpl[T]) Advancing(step uint) Stream[T] {
	return &streamImpl[T]{iter: internal.Advancing(self.iter, step)}
}

func (self *streamImpl[T]) IsSorted(ord delegate.Ord[T]) bool {
	return internal.IsSorted(self.iter, ord)
}

type parStreamImpl[T any] struct {
	iter iterator.Iterator[T]
}

func ParFromIterator[T any](iter iterator.Iterator[T]) Stream[T] {
	return &streamImpl[T]{iter: iter}
}

func ParFromIterable[T any](iterable iterator.Iterable[T]) Stream[T] {
	return ParFromIterator(iterable.Iter())
}
