//go:build go1.21

package cmp

import "cmp"

type Ordered = cmp.Ordered

func Compare[T Ordered](lhs T, rhs T) int {
	return cmp.Compare(lhs, rhs)
}

func Less[T Ordered](lhs T, rhs T) bool {
	return cmp.Less(lhs, rhs)
}

func LessThan[T Ordered](lhs T, rhs T) bool {
	return cmp.Compare(lhs, rhs) <= 0
}

func Greater[T Ordered](lhs T, rhs T) bool {
	return cmp.Compare(lhs, rhs) > 0
}

func GreaterThan[T Ordered](lhs T, rhs T) bool {
	return cmp.Compare(lhs, rhs) >= 0
}

func Equal[T comparable](lhs T, rhs T) bool { return lhs == rhs }

func NotEqual[T comparable](lhs T, rhs T) bool { return lhs != rhs }
