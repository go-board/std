package collections

import (
	"github.com/go-board/std/iter"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/tuple"
)

type MapEntry[K, V any] struct {
	inner tuple.Pair[K, V]
}

func (e MapEntry[K, V]) Key() K   { return e.inner.First }
func (e MapEntry[K, V]) Value() V { return e.inner.Second }

func MapEntryOf[K, V any](k K, v V) MapEntry[K, V] {
	return MapEntry[K, V]{inner: tuple.PairOf(k, v)}
}

type Map[K, V any] interface {
	Insert(k K, v K) optional.Optional[V]
	Get(k K) optional.Optional[V]
	Remove(k K) optional.Optional[V]
	Contains(k K) bool
	Keys() []K
	Values() []V

	iter.Iterable[MapEntry[K, V]]
	KeyIter() iter.Iter[K]
	ValueIter() iter.Iter[V]
}

func MapOf[K, V any](entries ...MapEntry[K, V]) Map[K, V] {
	return nil
}

type SortedMap[K, V any] interface {
	Map[K, V]

	First() optional.Optional[MapEntry[K, V]]
	Last() optional.Optional[MapEntry[K, V]]
	PopFirst() optional.Optional[MapEntry[K, V]]
	PopLast() optional.Optional[MapEntry[K, V]]
}

type Set[T any] interface {
	Insert(elem T)
	InsertMany(elems ...T)
	Contains(elem T) bool
	Remove(elem T) bool

	iter.Iterable[T]
}

type SortedSet[T any] interface {
	Set[T]

	First() optional.Optional[T]
	Last() optional.Optional[T]
	PopFirst() optional.Optional[T]
	PopLast() optional.Optional[T]
}

type Range[E comparable] struct {
	start     E
	end       E
	current   E
	successor func(E) E
}

func (r *Range[E]) Next() optional.Optional[E] {
	e := r.successor(r.current)
	if e == r.end {
		return optional.Optional[E]{}
	}
	r.current = e
	return optional.Some(e)
}

func NewRange[E comparable](start E, end E, successor func(E) E) {
}
