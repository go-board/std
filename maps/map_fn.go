package maps

import "github.com/go-board/std/cmp"

// Keys returns a map's key list without order promise.
func Keys[K comparable, V any, M ~map[K]V](m M) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a map's value list without order promise.
func Values[K comparable, V any, M ~map[K]V](m M) []V {
	values := make([]V, 0, len(m))
	for _, val := range m {
		values = append(values, val)
	}
	return values
}

func ForEach[K comparable, V any, M ~map[K]V](m M, fn func(k K, v V) bool) {
	for k, v := range m {
		if !fn(k, v) {
			return
		}
	}
}

func EqualBy[K comparable, V any, M ~map[K]V](lhs, rhs M, eq cmp.EqFunc[V]) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for k, v := range lhs {
		if rv, ok := rhs[k]; !ok || !eq(v, rv) {
			return false
		}
	}
	return true
}

func Equal[K comparable, V comparable, M ~map[K]V](lhs M, rhs M) bool {
	return EqualBy(lhs, rhs, func(l, r V) bool { return l == r })
}

func Reverse[K, V comparable, M ~map[K]V](m M) map[V]K {
	reversed := make(map[V]K)
	for k, v := range m {
		reversed[v] = k
	}
	return reversed
}

func MutateUnion[K comparable, V any, M ~map[K]V](force bool, lhs M, rhs ...M) {
	for _, m := range rhs {
		for k, v := range m {
			if _, ok := lhs[k]; !ok || force {
				lhs[k] = v
			}
		}
	}
}

func Union[K comparable, V any, M ~map[K]V](force bool, ms ...M) map[K]V {
	resultMap := make(map[K]V)
	MutateUnion[K, V, M](force, resultMap, ms...)
	return resultMap
}
