package maps

type Entry[K, V any] interface {
	GetKey() K
	GetValue() V
}

type mapEntry[K, V any] struct {
	key   K
	value V
}

func (e *mapEntry[K, V]) GetKey() K { return e.key }

func (e *mapEntry[K, V]) GetValue() V { return e.value }

func NewEntry[K comparable, V any](k K, v V) Entry[K, V] {
	return &mapEntry[K, V]{key: k, value: v}
}
