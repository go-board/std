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

func Of[T any](iter iterator.Iterator[T]) Stream[T] {
	return &streamImpl[T]{iter: iter}
}

func OfIterable[T any](iterable iterator.Iterable[T]) Stream[T] {
	return Of(iterable.Iter())
}

func (s *streamImpl[T]) Iter() iterator.Iterator[T] {
	return s.iter
}

func (s *streamImpl[T]) All(predicate delegate.Predicate[T]) bool {
	return internal.All(s.iter, predicate)
}

func (s *streamImpl[T]) Any(predicate delegate.Predicate[T]) bool {
	return internal.Any(s.iter, predicate)
}

func (s *streamImpl[T]) None(predicate delegate.Predicate[T]) bool {
	return internal.None(s.iter, predicate)
}

func (s *streamImpl[T]) Chain(o iterator.Iterable[T]) Stream[T] {
	return &streamImpl[T]{iter: internal.Chain[T](s.iter, o)}
}

func (s *streamImpl[T]) Filter(predicate delegate.Predicate[T]) Stream[T] {
	return &streamImpl[T]{iter: internal.Filter(s.iter, predicate)}
}

func (s *streamImpl[T]) Reduce(reduce delegate.Add[T, T, T]) optional.Optional[T] {
	return internal.Reduce(s.iter, reduce)
}

func (s *streamImpl[T]) Flatten(flatten delegate.Transform[T, iterator.Iterator[T]]) Stream[T] {
	return &streamImpl[T]{iter: internal.Flatten(s.iter, flatten)}
}

func (s *streamImpl[T]) Inspect(inspect delegate.Consumer1[T]) Stream[T] {
	return &streamImpl[T]{iter: internal.Inspect(s.iter, inspect)}
}

func (s *streamImpl[T]) MaxBy(cmp delegate.Comparison[T]) optional.Optional[T] {
	return internal.MaxBy(s.iter, cmp)
}

func (s *streamImpl[T]) MinBy(cmp delegate.Comparison[T]) optional.Optional[T] {
	return internal.MinBy(s.iter, cmp)
}

func (s *streamImpl[T]) SumBy(sum delegate.Add[T, T, T]) optional.Optional[T] {
	return internal.SumBy(s.iter, sum)
}

func (s *streamImpl[T]) Equal(o iterator.Iterable[T], cmp delegate.Equal[T]) bool {
	return internal.EqualBy(s.iter, o, cmp)
}

func (s *streamImpl[T]) Collect() []T {
	return internal.Collect(s.iter)
}

func (s *streamImpl[T]) Nth(n uint) optional.Optional[T] {
	return internal.Nth(s.iter, n)
}

func (s *streamImpl[T]) Take(n uint) Stream[T] {
	return &streamImpl[T]{iter: internal.Take(s.iter, n)}
}

func (s *streamImpl[T]) Skip(n uint) Stream[T] {
	return &streamImpl[T]{iter: internal.Skip(s.iter, n)}
}

func (s *streamImpl[T]) TakeWhile(predicate delegate.Predicate[T]) Stream[T] {
	return &streamImpl[T]{iter: internal.TakeWhile(s.iter, predicate)}
}

func (s *streamImpl[T]) SkipWhile(predicate delegate.Predicate[T]) Stream[T] {
	return &streamImpl[T]{iter: internal.SkipWhile(s.iter, predicate)}
}

func (s *streamImpl[T]) ForEach(consumer delegate.Consumer1[T]) {
	internal.ForEach(s.iter, consumer)
}

func (s *streamImpl[T]) StepBy(step uint) Stream[T] {
	return &streamImpl[T]{iter: internal.StepBy(s.iter, step)}
}

func (s *streamImpl[T]) Last() optional.Optional[T] {
	return internal.Last(s.iter)
}

func (s *streamImpl[T]) Advancing(step uint) Stream[T] {
	return &streamImpl[T]{iter: internal.Advancing(s.iter, step)}
}

func (s *streamImpl[T]) IsSorted(cmp delegate.Comparison[T]) bool {
	return internal.IsSorted(s.iter, cmp)
}
