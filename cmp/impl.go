package cmp

import (
	"github.com/go-board/std/cond"
	"github.com/go-board/std/core"
)

// MaxBy calculates the maximum value of two values by a given function.
func MaxBy[A any, F func(A, A) Ordering](compare F, lhs, rhs A) A {
	return cond.Ternary(compare(lhs, rhs).IsGt(), lhs, rhs)
}

// MaxByKey calculates the maximum value of two values by a key function.
func MaxByKey[A any, F func(A) K, K Ord[K]](f F, lhs, rhs A) A {
	keyLhs, keyRhs := f(lhs), f(rhs)
	return cond.Ternary(keyLhs.Cmp(keyRhs).IsGe(), lhs, rhs)
}

// Max calculates the maximum value of two values.
func Max[A Ord[A]](lhs A, rhs A) A {
	return MaxBy(func(lhs, rhs A) Ordering { return lhs.Cmp(rhs) }, lhs, rhs)
}

// MaxOrdered calculates the maximum value of two ordered values.
func MaxOrdered[A core.Ordered](lhs, rhs A) A {
	return cond.Ternary(lhs > rhs, lhs, rhs)
}

// MinBy calculates the minimum value of two values by a given function.
func MinBy[A any, F func(A, A) Ordering](compare F, lhs, rhs A) A {
	return cond.Ternary(compare(lhs, rhs).IsLt(), lhs, rhs)
}

// MinByKey calculates the minimum value of two values by a key function.
func MinByKey[A any, F func(A) K, K Ord[K]](f F, lhs, rhs A) A {
	keyLhs, keyRhs := f(lhs), f(rhs)
	return cond.Ternary(keyLhs.Cmp(keyRhs).IsLt(), lhs, rhs)
}

// Min calculates the minimum value of two values.
func Min[A Ord[A]](lhs, rhs A) A {
	return MinBy(func(lhs, rhs A) Ordering { return lhs.Cmp(rhs) }, lhs, rhs)
}

// MinOrdered calculates the minimum value of two ordered values.
func MinOrdered[A core.Ordered](lhs, rhs A) A {
	return cond.Ternary(lhs < rhs, lhs, rhs)
}

// EqByKey calculates the equality of two values by a key function.
func EqByKey[A any, F func(A) K, K Eq[K]](f F, lhs, rhs A) bool {
	lhsKey, rhsKey := f(lhs), f(rhs)
	return lhsKey.Eq(rhsKey)
}

// CmpByKey calculates the comparison of two values by a key function.
func CmpByKey[A any, F func(A) K, K Ord[K]](f F, lhs, rhs A) Ordering {
	lhsKey, rhsKey := f(lhs), f(rhs)
	return lhsKey.Cmp(rhsKey)
}

// Order is a function that returns an Ordering of two Ordering.
//  result table is:
//		| lhs\rhs | Less    |  Equal  | Greater |
//		|---------|---------|---------|---------+
//		| Less    | Equal   | Less    | Less    |
//		| Equal   | Greater | Equal   | Less    |
//		| Greater | Greater | Greater | Equal   |
func Order(lhs, rhs Ordering) Ordering {
	lhsInner, rhsInner := lhs.(order), rhs.(order)
	return cond.Ternary(lhsInner == rhsInner, Equal, cond.Ternary(lhsInner < rhsInner, Less, Greater))
}
