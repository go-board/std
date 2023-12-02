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

func Collect[K comparable, V any](s iter.Seq[MapEntry[K, V]]) map[K]V {
	m := make(map[K]V)
	iter.ForEach(s, func(e MapEntry[K, V]) { m[e.Key()] = e.Value() })
	return m
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

// ForEach iter over the map, and call the udf on each k-v pair.
func ForEach[K comparable, V any, M ~map[K]V](m M, f func(K, V)) {
	iter.ForEach(Entries(m), func(x MapEntry[K, V]) {
		f(x.Key(), x.Value())
	})
}

// Map map each k-v pair to x-y pair into a new map.
func Map[K, X comparable, V, Y any, M ~map[K]V](m M, f func(K, V) (X, Y)) map[X]Y {
	n := make(map[X]Y)
	for k, v := range m {
		x, y := f(k, v)
		n[x] = y
	}
	return n
}

func MapKey[K, X comparable, V any, M ~map[K]V](m M, f func(K) X) map[X]V {
	n := make(map[X]V)
	for k, v := range m {
		n[f(k)] = v
	}
	return n
}

func MapValue[K comparable, V, X any, M ~map[K]V](m M, f func(V) X) map[K]X {
	n := make(map[K]X)
	for k, v := range m {
		n[k] = f(v)
	}
	return n
}

// Filter keep those elements which match the given predicate function.
func Filter[K comparable, V any, M ~map[K]V](m M, f func(K, V) bool) M {
	return Collect(iter.Filter(Entries(m), func(me MapEntry[K, V]) bool {
		return f(me.Key(), me.Value())
	}))
}

// FilterMap keep those elements which match the given predicate function and map to new type elements.
func FilterMap[K comparable, V any, M ~map[K]V, X comparable, Y any, N ~map[X]Y](m M, f func(K, V) (X, Y, bool)) N {
	r := N{}
	for k, v := range m {
		if x, y, ok := f(k, v); ok {
			r[x] = y
		}
	}
	return r
}

// MergeKeep merge many maps and keep first value when conflict occrured.
func MergeKeep[K comparable, V any, M ~map[K]V](ms iter.Seq[M]) M {
	x := M{}
	iter.ForEach(ms, func(m M) {
		ForEach(m, func(k K, v V) {
			if _, ok := x[k]; !ok {
				x[k] = v
			}
		})
	})
	return x
}

// MergeOverwrite merge many maps and keep last value when conflict occrured.
func MergeOverwrite[K comparable, V any, M ~map[K]V](ms iter.Seq[M]) M {
	x := M{}
	iter.ForEach(ms, func(m M) { ForEach(m, func(k K, v V) { x[k] = v }) })
	return x
}

// MergeFunc merge many maps and solve conflict use a udf.
//
// UDF has signature `func(V, V) V`, first param is previous element,
// second is visit element, return element will be used.
func MergeFunc[K comparable, V any, M ~map[K]V](ms iter.Seq[M], onConflict func(V, V) V) M {
	x := make(M)
	iter.ForEach(ms, func(m M) { ForEach(m, func(k K, v V) { x[k] = onConflict(x[k], v) }) })
	return x
}

// Invert maps k-v to v-k, when conflict, the back element will overwrite the previous one.
func Invert[K comparable, V comparable, M1 ~map[K]V, M2 ~map[V]K](m M1) M2 {
	m2 := make(M2)
	for k, v := range m {
		m2[v] = k
	}
	return m2
}
