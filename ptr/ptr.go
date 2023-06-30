package ptr

import "github.com/go-board/std/cmp"

func zero[T any]() (v T) { return }

// Ref return reference of value
func Ref[T any](t T) *T { return &t }

// Default return default value of type
func Default[T any]() *T {
	return Ref(zero[T]())
}

// ValueOr return value of pointer if not nil, else return default value.
func ValueOr[T any](v *T, d T) T {
	if v == nil {
		return d
	}
	return *v
}

// ValueOrZero return value of pointer if not nil, else return zero value.
func ValueOrZero[T any](v *T) T {
	return ValueOr(v, zero[T]())
}

// Compare compares two pointer. If both non-nil, compare underlying data,
// if both nil, return 0, non-nil pointer is always greater than nil pointer.
func Compare[T cmp.Ordered](l, r *T) int {
	if l != nil && r != nil {
		return cmp.Compare(*l, *r)
	}
	if l == nil && r == nil {
		return 0
	}
	if l != nil {
		return +1
	}
	return -1
}

// Equal test whether two pointer are equal. If both non-nil, test underlying data,
// if both nil, return true, else return false
func Equal[T comparable](l, r *T) bool {
	if l != nil && r != nil {
		return *l == *r
	} else if l == nil && r == nil {
		return true
	}
	return false
}
