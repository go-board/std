package stream

import (
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/iterator/ops"
	"github.com/go-board/std/optional"
)

type streamImpl[T any] struct{ iter iterator.Iterator[T] }

var _ Stream[any] = (*streamImpl[any])(nil)

// FromIterator returns a Stream[T] from an iterator.Iterator[T].
func FromIterator[T any](iter iterator.Iterator[T]) Stream[T] {
	return &streamImpl[T]{iter: iter}
}

// FromIterable returns a Stream[T] from an iterator.Iterable[T].
func FromIterable[T any](iterable iterator.Iterable[T]) Stream[T] {
	return FromIterator(iterable.Iter())
}

func (self *streamImpl[T]) Iter() iterator.Iterator[T] { return self.iter }

func (self *streamImpl[T]) All(predicate func(T) bool) bool {
	return ops.All(self.iter, predicate)
}

func (self *streamImpl[T]) Any(predicate func(T) bool) bool {
	return ops.Any(self.iter, predicate)
}

func (self *streamImpl[T]) Once(predicate func(T) bool) bool {
	return ops.Once(self.iter, predicate)
}

func (self *streamImpl[T]) None(predicate func(T) bool) bool {
	return ops.None(self.iter, predicate)
}

func (self *streamImpl[T]) Chain(o iterator.Iterable[T]) Stream[T] {
	return &streamImpl[T]{iter: ops.Chain[T](self.iter, o)}
}

func (self *streamImpl[T]) Map(transformer func(T) T) Stream[T] {
	return &streamImpl[T]{iter: ops.Map(self.iter, transformer)}
}

func (self *streamImpl[T]) Filter(predicate func(T) bool) Stream[T] {
	return &streamImpl[T]{iter: ops.Filter(self.iter, predicate)}
}

func (self *streamImpl[T]) Fold(initial T, accumulator func(T, T) T) T {
	return ops.Fold(self.iter, initial, accumulator)
}

func (self *streamImpl[T]) Reduce(reduce func(T, T) T) optional.Optional[T] {
	return ops.Reduce(self.iter, reduce)
}

func (self *streamImpl[T]) Flatten(flatten func(T) iterator.Iterator[T]) Stream[T] {
	return &streamImpl[T]{iter: ops.Flatten(self.iter, flatten)}
}

func (self *streamImpl[T]) Chunk(size int) Stream[Stream[T]] {
	panic("implement me")
}

func (self *streamImpl[T]) Inspect(inspect func(T)) Stream[T] {
	return &streamImpl[T]{iter: ops.Inspect(self.iter, inspect)}
}

func (self *streamImpl[T]) MaxBy(ord cmp.OrdFunc[T]) optional.Optional[T] {
	return ops.MaxBy(self.iter, ord)
}

func (self *streamImpl[T]) MinBy(ord cmp.OrdFunc[T]) optional.Optional[T] {
	return ops.MinBy(self.iter, ord)
}

func (self *streamImpl[T]) EqualBy(o iterator.Iterable[T], ord cmp.EqFunc[T]) bool {
	return ops.EqualBy(self.iter, o, ord)
}

func (self *streamImpl[T]) Collect() []T {
	return ops.Collect(self.iter)
}

func (self *streamImpl[T]) Nth(n uint) optional.Optional[T] {
	return ops.Nth(self.iter, n)
}

func (self *streamImpl[T]) Take(n uint) Stream[T] {
	return &streamImpl[T]{iter: ops.Take(self.iter, n)}
}

func (self *streamImpl[T]) Skip(n uint) Stream[T] {
	return &streamImpl[T]{iter: ops.Skip(self.iter, n)}
}

func (self *streamImpl[T]) TakeWhile(predicate func(T) bool) Stream[T] {
	return &streamImpl[T]{iter: ops.TakeWhile(self.iter, predicate)}
}

func (self *streamImpl[T]) SkipWhile(predicate func(T) bool) Stream[T] {
	return &streamImpl[T]{iter: ops.SkipWhile(self.iter, predicate)}
}

func (self *streamImpl[T]) ForEach(consumer func(T)) {
	ops.ForEach(self.iter, consumer)
}

func (self *streamImpl[T]) StepBy(step uint) Stream[T] {
	return &streamImpl[T]{iter: ops.StepBy(self.iter, step)}
}

func (self *streamImpl[T]) First() optional.Optional[T] {
	return self.iter.Next()
}

func (self *streamImpl[T]) Last() optional.Optional[T] {
	return ops.Last(self.iter)
}

func (self *streamImpl[T]) Advancing(step uint) Stream[T] {
	return &streamImpl[T]{iter: ops.Advancing(self.iter, step)}
}

func (self *streamImpl[T]) IsSorted(ord cmp.OrdFunc[T]) bool {
	return ops.IsSorted(self.iter, ord)
}
