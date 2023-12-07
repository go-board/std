package maps

import (
	"github.com/go-board/std/iter"
	"github.com/go-board/std/tuple"
)

// MapEntry persist k-v pair of a map.
type MapEntry[K, V any] struct{ inner tuple.Pair[K, V] }

func (e MapEntry[K, V]) Key() K   { return e.inner.First }
func (e MapEntry[K, V]) Value() V { return e.inner.Second }

func entry[K, V any](key K, value V) MapEntry[K, V] {
	return MapEntry[K, V]{inner: tuple.PairOf(key, value)}
}

// Entries returns all entry of a map.
func Entries[K comparable, V any, M ~map[K]V](m M) iter.Seq[MapEntry[K, V]] {
	return func(yield func(MapEntry[K, V]) bool) {
		for k, v := range m {
			if !yield(entry(k, v)) {
				break
			}
		}
	}
}

// Keys return key slice of a map.
func Keys[K comparable, V any, M ~map[K]V](m M) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range m {
			if !yield(k) {
				break
			}
		}
	}
}

// Values returns value slice of a map.
func Values[K comparable, V any, M ~map[K]V](m M) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range m {
			if !yield(v) {
				break
			}
		}
	}
}

func Collect[K comparable, V any](s iter.Seq[MapEntry[K, V]]) map[K]V {
	m := make(map[K]V)
	CollectInto(s, m)
	return m
}

func CollectInto[K comparable, V any, M ~map[K]V](s iter.Seq[MapEntry[K, V]], m M) {
	iter.CollectFunc(s, func(x MapEntry[K, V]) bool { m[x.Key()] = x.Value(); return true })
}

// ForEach iter over the map, and call the udf on each k-v pair.
func ForEach[K comparable, V any, M ~map[K]V](m M, f func(K, V)) {
	iter.ForEach(Entries(m), func(x MapEntry[K, V]) {
		f(x.Key(), x.Value())
	})
}

// Map call f on each k-v pair and maps to x-y pair into a new map.
func Map[K, X comparable, V, Y any, M ~map[K]V](m M, f func(K, V) (X, Y)) map[X]Y {
	return Collect(iter.Map(Entries(m), func(e MapEntry[K, V]) MapEntry[X, Y] {
		return entry(f(e.Key(), e.Value()))
	}))
}

func MapKey[K, X comparable, V any, M ~map[K]V](m M, f func(K, V) X) map[X]V {
	return Map(m, func(k K, v V) (X, V) { return f(k, v), v })
}

func MapValue[K comparable, V, X any, M ~map[K]V](m M, f func(K, V) X) map[K]X {
	return Map(m, func(k K, v V) (K, X) { return k, f(k, v) })
}

func Retain[K comparable, V any, M ~map[K]V](m M, f func(K, V) bool) {
	for k, v := range m {
		if !f(k, v) {
			delete(m, k)
		}
	}
}

// Filter keep those elements which match the given predicate function.
func Filter[K comparable, V any, M ~map[K]V](m M, f func(K, V) bool) M {
	return Collect(Entries(m).Filter(func(me MapEntry[K, V]) bool { return f(me.Key(), me.Value()) }))
}

// FilterMap keep those elements which match the given predicate function and map to new type elements.
func FilterMap[K comparable, V any, M ~map[K]V, X comparable, Y any, N ~map[X]Y](m M, f func(K, V) (X, Y, bool)) N {
	return Collect(iter.FilterMap(Entries(m), func(e MapEntry[K, V]) (MapEntry[X, Y], bool) {
		if x, y, ok := f(e.Key(), e.Value()); ok {
			return entry(x, y), true
		}
		return MapEntry[X, Y]{}, false
	}))
}

// MergeKeep merge many maps and keep first value when conflict occurred.
func MergeKeep[K comparable, V any, M ~map[K]V](ms iter.Seq[M]) M {
	return MergeFunc(ms, func(key K, prev V, current V) V { return prev })
}

// MergeOverwrite merge many maps and keep last value when conflict occurred.
func MergeOverwrite[K comparable, V any, M ~map[K]V](ms iter.Seq[M]) M {
	return MergeFunc(ms, func(key K, prev V, current V) V { return current })
}

// MergeFunc merge many maps and solve conflict use an udf.
//
// UDF has signature `func(V, V) V`, first param is previous element,
// second is visit element, return element will be used.
func MergeFunc[K comparable, V any, M ~map[K]V](ms iter.Seq[M], onConflict func(key K, prev V, current V) V) M {
	x := make(M)
	iter.ForEach(ms, func(m M) { ForEach(m, func(k K, v V) { x[k] = onConflict(k, x[k], v) }) })
	return x
}

// Invert maps k-v to v-k, when key conflict, the back element will overwrite the previous one.
func Invert[K, V comparable, M1 ~map[K]V, M2 ~map[V]K](m M1) M2 {
	m2 := make(M2)
	for k, v := range m {
		m2[v] = k
	}
	return m2
}
