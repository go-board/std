package maps

type Entry[K, V any] interface {
	GetKey() K
	GetValue() V
}

type mapEntry[K, V any] struct {
	key   K
	value V
}

func (self *mapEntry[K, V]) GetKey() K { return self.key }

func (self *mapEntry[K, V]) GetValue() V { return self.value }

func NewEntry[K comparable, V any](k K, v V) Entry[K, V] {
	return &mapEntry[K, V]{key: k, value: v}
}
