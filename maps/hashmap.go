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

func (m *HashMap[K, V]) Keys() []K {
	ks := make([]K, 0)
	for k := range m.store {
		ks = append(ks, k)
	}
	return ks
}

func (m *HashMap[K, V]) Entries() []Entry[K, V] {
	es := make([]Entry[K, V], 0)
	for k, v := range m.store {
		es = append(es, &mapEntry[K, V]{key: k, value: v})
	}
	return es
}
