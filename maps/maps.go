package maps

import "github.com/go-board/std/tuple"

type MapEntry[K, V any] struct{ inner tuple.Pair[K, V] }

func (e MapEntry[K, V]) Key() K   { return e.inner.First }
func (e MapEntry[K, V]) Value() V { return e.inner.Second }

func entry[K, V any](key K, value V) MapEntry[K, V] {
	return MapEntry[K, V]{inner: tuple.PairOf(key, value)}
}

func Entries[K comparable, V any, M ~map[K]V](m M) []MapEntry[K, V] {
	entries := make([]MapEntry[K, V], 0, len(m))
	for k, v := range m {
		entries = append(entries, entry(k, v))
	}
	return entries
}

func Keys[K comparable, V any, M ~map[K]V](m M) []K {
	ks := make([]K, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

func Values[K comparable, V any, M ~map[K]V](m M) []V {
	vs := make([]V, 0, len(m))
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}
