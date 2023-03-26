package sync

import (
	"sync"

	"github.com/go-board/std/optional"
)

// Map is generic version of sync.Map.
type Map[K any, V any] struct{ inner sync.Map }

func (self *Map[K, V]) Insert(key K, v V) {
	self.inner.Store(key, v)
}

func (self *Map[K, V]) Get(key K) (V, bool) {
	val, ok := self.inner.Load(key)
	if !ok {
		var t V
		return t, false
	}
	return val.(V), true
}

func (self *Map[K, V]) GetOptional(key K) optional.Optional[V] {
	return optional.FromPair(self.Get(key))
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
