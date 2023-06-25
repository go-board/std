//go:build go1.21

package maps

import "maps"

func Entries[K comparable, V any, M ~map[K]V](m M) []MapEntry[K, V] {
	es := make([]MapEntry[K, V], 0, len(m))
	for k, v := range m {
		es = append(es, entry(k, v))
	}
	return es
}

func Keys[K comparable, V any, M ~map[K]V](m M) []K {
	return maps.Keys(m)
}

func Values[K comparable, V any, M ~map[K]V](m M) []V {
	return maps.Values(m)
}

func Equal[K comparable, V comparable, M1 ~map[K]V, M2 ~map[K]V](lhs M1, rhs M2) bool {
	return maps.Equal(lhs, rhs)
}

func EqualBy[K comparable, V1, V2 any, M1 ~map[K]V1, M2 ~map[K]V2](lhs M1, rhs M2, equal func(V1, V2) bool) bool {
	return maps.EqualFunc(lhs, rhs, equal)
}

func Clone[K comparable, V any, M ~map[K]V](m M) M {
	return maps.Clone(m)
}

func DeleteBy[K comparable, V any](x map[K]V, f func(K, V) bool) {
	maps.DeleteFunc(x, f)
}
