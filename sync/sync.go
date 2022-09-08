package sync

import "sync"

// Map is generic version of sync.Map.
type Map[K any, V any] struct{ inner sync.Map }

// Put puts value into the map.
func (self *Map[K, V]) Put(key K, v V) {
	self.inner.Store(key, v)
}

// Get returns the value associated with key, or nil if there is none.
func (self *Map[K, V]) Get(key K) (V, bool) {
	val, ok := self.inner.Load(key)
	if !ok {
		var t V
		return t, false
	}
	return val.(V), true
}

// Contains return true if map contains a key.
//
// Example:
//
//	var m Map[int]int
//	m.Put(1, 2)
//	ok1 := m.Contains(1) // true
//	ok2 := m.Contains(2) // false
func (self *Map[K, V]) Contains(k K) bool {
	_, ok := self.inner.Load(k)
	return ok
}

// GetOrPut returns a value from the map if it exists and then puts it in the map.
func (self *Map[K, V]) GetOrPut(key K, val V) (V, bool) {
	replaced, ok := self.inner.LoadOrStore(key, val)
	if !ok {
		var t V
		return t, false
	}
	return replaced.(V), true
}

// Range iterate over all elements in the map.
func (self *Map[K, V]) Range(fn func(K, V) bool) {
	self.inner.Range(func(key, value any) bool { return fn(key.(K), value.(V)) })
}

// Delete remove specific key from the map if present.
func (self *Map[K, V]) Delete(key K) {
	self.inner.Delete(key)
}

// Pool is generic version of sync.Pool
type Pool[T any, P *T] struct{ inner sync.Pool }

// NewPool creates a new Pool which can provide new or existing instace of T.
func NewPool[T any, P *T](fn func() P) *Pool[T, P] {
	return &Pool[T, P]{inner: sync.Pool{New: func() any { return fn() }}}
}

// Put puts a value into the pool and returns the value.
func (self *Pool[T, P]) Put(val P) { self.inner.Put(val) }

// Get gets a value from the pool and returns the value.
func (self *Pool[T, P]) Get() P { return self.inner.Get().(P) }
