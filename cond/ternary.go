package cond

// Ternary operator with a condition and two possible results.
// The condition is evaluated and the first result is returned if it is true,
// otherwise the second result is returned.
//
//	cond.Ternary(true, "yes", "no") // "yes"
//	cond.Ternary(false, "yes", "no") // "no"
func Ternary[T any](condition bool, ifTrue, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}
