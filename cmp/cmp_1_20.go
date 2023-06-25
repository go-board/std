//go:build !go1.21

package cmp

import (
	"github.com/go-board/std/core"
)

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
func isNaN[T Ordered](x T) bool {
	return x != x
}

type Ordered = core.Ordered

func Compare[T Ordered](lhs T, rhs T) int {
	xNaN := isNaN(lhs)
	yNaN := isNaN(rhs)
	if xNaN && yNaN {
		return 0
	}
	if xNaN || lhs < rhs {
		return -1
	}
	if yNaN || lhs > rhs {
		return +1
	}
	return 0
}

func Less[T Ordered](lhs T, rhs T) bool {
	return Compare(lhs, rhs) < 0
}

func LessThan[T Ordered](lhs T, rhs T) bool {
	return Compare(lhs, rhs) <= 0
}

func Greater[T Ordered](lhs T, rhs T) bool {
	return Compare(lhs, rhs) > 0
}

func GreaterThan[T Ordered](lhs T, rhs T) bool {
	return Compare(lhs, rhs) >= 0
}

func Equal[T comparable](lhs T, rhs T) bool { return lhs == rhs }

func NotEqual[T comparable](lhs T, rhs T) bool { return lhs != rhs }
