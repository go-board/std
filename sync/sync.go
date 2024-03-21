package sync

import (
	"sync"

	"github.com/go-board/std/iter"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/tuple"
)

// Map is generic version of sync.Map.
type Map[K any, V any] struct{ inner sync.Map }

func (self *Map[K, V]) Insert(key K, v V) {
	self.inner.Store(key, v)
}

func (self *Map[K, V]) Get(key K) (V, bool) {
	return self.GetOptional(key).Get()
}

func (self *Map[K, V]) GetOptional(key K) optional.Optional[V] {
	return optional.Map(optional.FromPair(self.inner.Load(key)), func(x any) V { return x.(V) })
}

func (self *Map[K, V]) GetOrPut(key K, val V) (V, bool) {
	replaced, ok := self.inner.LoadOrStore(key, val)
	if !ok {
		var t V
		return t, false
	}
	return replaced.(V), true
}

func (self *Map[K, V]) Range(fn func(K, V) bool) {
	self.inner.Range(func(key, value any) bool { return fn(key.(K), value.(V)) })
}

func (self *Map[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		self.Range(func(k K, v V) bool { return yield(k) })
	}
}

func (self *Map[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		self.Range(func(k K, v V) bool { return yield(v) })
	}
}

func (self *Map[K, V]) Entries() iter.Seq[tuple.Pair[K, V]] {
	return func(yield func(tuple.Pair[K, V]) bool) {
		self.Range(func(k K, v V) bool { return yield(tuple.MakePair(k, v)) })
	}
}

func (self *Map[K, V]) Delete(key K) {
	self.inner.Delete(key)
}

// Pool is generic version of sync.Pool
type Pool[T any] struct{ inner sync.Pool }

func NewPool[T any](fn func() *T) *Pool[T] {
	return &Pool[T]{inner: sync.Pool{New: func() any { return fn() }}}
}

func (self *Pool[T]) Put(val *T) { self.inner.Put(val) }

func (self *Pool[T]) Get() *T { return self.inner.Get().(*T) }
