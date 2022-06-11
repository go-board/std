package stream

import (
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/iterator/internal"
	"github.com/go-board/std/optional"
)

type streamImpl[T any] struct {
	iter iterator.Iterator[T]
}

var _ Stream[any] = (*streamImpl[any])(nil)

func FromIterator[T any](iter iterator.Iterator[T]) Stream[T] {
	return &streamImpl[T]{iter: iter}
}

func FromIterable[T any](iterable iterator.Iterable[T]) Stream[T] {
	return FromIterator(iterable.Iter())
}

func (self *streamImpl[T]) Iter() iterator.Iterator[T] { return self.iter }

func (self *streamImpl[T]) All(predicate func(T) bool) bool {
	return internal.All(self.iter, predicate)
}

func (self *streamImpl[T]) Any(predicate func(T) bool) bool {
	return internal.Any(self.iter, predicate)
}

func (self *streamImpl[T]) Once(predicate func(T) bool) bool {
	return internal.Once(self.iter, predicate)
}

func (self *streamImpl[T]) None(predicate func(T) bool) bool {
	return internal.None(self.iter, predicate)
}

func (self *streamImpl[T]) Chain(o iterator.Iterable[T]) Stream[T] {
	return &streamImpl[T]{iter: internal.Chain[T](self.iter, o)}
}

func (self *streamImpl[T]) Map(transformer func(T) T) Stream[T] {

	return &streamImpl[T]{iter: internal.Map(self.iter, transformer)}
}

func (self *streamImpl[T]) Filter(predicate func(T) bool) Stream[T] {
	return &streamImpl[T]{iter: internal.Filter(self.iter, predicate)}
}

func (self *streamImpl[T]) Reduce(reduce func(T, T) T) optional.Optional[T] {
	return internal.Reduce(self.iter, reduce)
}

func (self *streamImpl[T]) Flatten(flatten func(T) iterator.Iterator[T]) Stream[T] {
	return &streamImpl[T]{iter: internal.Flatten(self.iter, flatten)}
}

func (self *streamImpl[T]) Inspect(inspect func(T)) Stream[T] {
	return &streamImpl[T]{iter: internal.Inspect(self.iter, inspect)}
}

func (self *streamImpl[T]) MaxBy(ord cmp.OrdFunc[T]) optional.Optional[T] {
	return internal.MaxBy(self.iter, ord)
}

func (self *streamImpl[T]) MinBy(ord cmp.OrdFunc[T]) optional.Optional[T] {
	return internal.MinBy(self.iter, ord)
}

func (self *streamImpl[T]) Equal(o iterator.Iterable[T], ord cmp.EqFunc[T]) bool {
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

func (self *streamImpl[T]) TakeWhile(predicate func(T) bool) Stream[T] {
	return &streamImpl[T]{iter: internal.TakeWhile(self.iter, predicate)}
}

func (self *streamImpl[T]) SkipWhile(predicate func(T) bool) Stream[T] {
	return &streamImpl[T]{iter: internal.SkipWhile(self.iter, predicate)}
}

func (self *streamImpl[T]) ForEach(consumer func(T)) {
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

func (self *streamImpl[T]) IsSorted(ord cmp.OrdFunc[T]) bool {
	return internal.IsSorted(self.iter, ord)
}
