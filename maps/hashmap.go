package maps

type HashMap[K comparable, V any] struct{ store map[K]V }

func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{store: map[K]V{}}
}

func (self *HashMap[K, V]) Put(key K, value V) {
	self.store[key] = value
}

func (self *HashMap[K, V]) Remove(keys ...K) {
	for _, k := range keys {
		delete(self.store, k)
	}
}

func (self *HashMap[K, V]) Iter(f func(K, V)) {
	for k, v := range self.store {
		f(k, v)
	}
}

func (self *HashMap[K, V]) Keys() []K {
	ks := make([]K, 0)
	for k := range self.store {
		ks = append(ks, k)
	}
	return ks
}

func (self *HashMap[K, V]) Entries() []Entry[K, V] {
	es := make([]Entry[K, V], 0)
	for k, v := range self.store {
		es = append(es, &mapEntry[K, V]{key: k, value: v})
	}
	return es
}
