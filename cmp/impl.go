package cmp

func ternary[T any](ok bool, lhs T, rhs T) T {
	if ok {
		return lhs
	}
	return rhs
}

// MaxBy calculates the maximum value of two values by a given function.
func MaxBy[A any, F ~func(A, A) Order](cmp F, lhs, rhs A) A {
	return ternary(cmp(lhs, rhs).IsGt(), lhs, rhs)
}

func MaxFunc[A any](cmp func(A, A) int, lhs, rhs A) A {
	return ternary(cmp(lhs, rhs) > 0, lhs, rhs)
}

// MaxByKey calculates the maximum value of two values by a key function.
func MaxByKey[A any, F ~func(A) K, K Ordered](key F, lhs, rhs A) A {
	return ternary(key(lhs) > key(rhs), lhs, rhs)
}

// MaxOrdered calculates the maximum value of two ordered values.
func MaxOrdered[A Ordered](lhs, rhs A) A {
	return ternary(lhs > rhs, lhs, rhs)
}

// MinBy calculates the minimum value of two values by a given function.
func MinBy[A any, F ~func(A, A) Order](cmp F, lhs, rhs A) A {
	return ternary(cmp(lhs, rhs).IsLt(), lhs, rhs)
}

func MinFunc[A any](cmp func(A, A) int, lhs, rhs A) A {
	return ternary(cmp(lhs, rhs) < 0, lhs, rhs)
}

// MinByKey calculates the minimum value of two values by a key function.
func MinByKey[A any, F ~func(A) K, K Ordered](key F, lhs, rhs A) A {
	return ternary(key(lhs) < key(rhs), lhs, rhs)
}

// MinOrdered calculates the minimum value of two ordered values.
func MinOrdered[A Ordered](lhs, rhs A) A {
	return ternary(lhs < rhs, lhs, rhs)
}

func Or[E comparable](elems ...E) E {
	var empty E
	for _, x := range elems {
		if empty != x {
			return x
		}
	}
	return empty
}
