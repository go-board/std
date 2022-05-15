package fp

// Function3 is a function that takes three arguments.
type Function3[A, B, C, R any] interface {
	Apply(A, B, C) R
	Curry() func(A) func(B) func(C) R
}

// Fn3 is a function that takes three arguments.
type Fn3[A, B, C, R any] func(A, B, C) R

func (f Fn3[A, B, C, R]) Apply(a A, b B, c C) R { return f(a, b, c) }

func (f Fn3[A, B, C, R]) Curry() func(A) func(B) func(C) R {
	return func(a A) func(B) func(C) R {
		return func(b B) func(C) R { return func(c C) R { return f.Apply(a, b, c) } }
	}
}

// Apply3 is a function that takes three arguments.
func Apply3[A, B, C, R any](f func(A, B, C) R, a A, b B, c C) R {
	return Fn3[A, B, C, R](f).Apply(a, b, c)
}

// Curry3 is a function that curry three arguments function.
func Curry3[A, B, C, R any](f func(A, B, C) R) func(A) func(B) func(C) R {
	return Fn3[A, B, C, R](f).Curry()
}

// Uncurry3 is a function that uncurry three arguments function.
func Uncurry3[A, B, C, R any](f func(A) func(B) func(C) R) func(A, B, C) R {
	return func(a A, b B, c C) R { return f(a)(b)(c) }
}
