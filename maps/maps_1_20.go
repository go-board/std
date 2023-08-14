//go:build !go1.21

package maps

import (
	"github.com/go-board/std/operator"
)

func Clone[K comparable, V any, M ~map[K]V](m M) M {
	r := make(M)
	for k, v := range m {
		r[k] = v
	}
	return r
}

func EqualBy[K comparable, V1, V2 any, M1 ~map[K]V1, M2 ~map[K]V2](lhs M1, rhs M2, equal func(V1, V2) bool) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for k, v1 := range lhs {
		v2, ok := rhs[k]
		if !ok || !equal(v1, v2) {
			return false
		}
	}
	return true
}

func Equal[K, V comparable, M1 ~map[K]V, M2 ~map[K]V](lhs M1, rhs M2) bool {
	return EqualBy(lhs, rhs, operator.Eq[V])
}
