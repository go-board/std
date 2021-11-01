package maps

type HashMap[K comparable, V any] struct{ store map[K]V }

func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{store: map[K]V{}}
}

func (m *HashMap[K, V]) Put(key K, value V) {
	m.store[key] = value
}

func (m *HashMap[K, V]) Remove(keys ...K) {
	for _, k := range keys {
		delete(m.store, k)
	}
}

func (m *HashMap[K, V]) Iter(f func(K, V)) {
	for k, v := range m.store {
		f(k, v)
	}
}
