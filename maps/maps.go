package maps

// DefaultHashMap is a hash map that uses a factory to generate default value for missing keys.
// And it is not thread safe, for concurrent access use with a sync.Mutex.
type DefaultHashMap[K comparable, V any] struct {
	factory func(K) V
	inner   map[K]V
}

// NewDefaultHashMap returns a new DefaultHashMap.
func NewDefaultHashMap[K comparable, V any](factory func(K) V) *DefaultHashMap[K, V] {
	return &DefaultHashMap[K, V]{factory: factory, inner: map[K]V{}}
}

// Has returns true if the map contains the key.
func (m *DefaultHashMap[K, V]) Has(key K) bool {
	_, ok := m.inner[key]
	return ok
}

// Get returns the value for the given key.
func (m *DefaultHashMap[K, V]) Get(key K) V {
	v, ok := m.inner[key]
	if !ok {
		v = m.factory(key)
		m.inner[key] = v
	}
	return v
}

// Set sets the value for the given key.
func (m *DefaultHashMap[K, V]) Set(key K, val V) {
	m.inner[key] = val
}

// Del deletes the value for the given key.
func (m *DefaultHashMap[K, V]) Del(key K) {
	delete(m.inner, key)
}

// Range calls the given function for each key-value pair in the map.
func (m *DefaultHashMap[K, V]) Range(f func(K, V)) {
	for k, v := range m.inner {
		f(k, v)
	}
}
