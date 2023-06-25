package cond

// Ternary operator with a condition and two possible results.
// The condition is evaluated and the first result is returned if it is true,
// otherwise the second result is returned.
//
//	cond.Ternary(true, "yes", "no") // "yes"
//	cond.Ternary(false, "yes", "no") // "no"
//
// 💣 Since go doesn't support lazy evaluation, so dereference a pointer will cause panic.
func Ternary[T any](condition bool, ifTrue, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}

//go:inline
func TernaryFunc[T any](condition bool, ifTrue, ifFalse func() T) T {
	if condition {
		return ifTrue()
	}
	return ifFalse()
}
