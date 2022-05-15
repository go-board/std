package fp

// Function2 is a function that takes two arguments.
type Function2[A, B, R any] interface {
	Apply(A, B) R
	Curry() func(A) func(B) R
}

// Fn2 is a function that takes two arguments.
type Fn2[A, B, R any] func(A, B) R

func (f Fn2[A, B, R]) Apply(a A, b B) R { return f(a, b) }

func (f Fn2[A, B, R]) Curry() func(A) func(B) R {
	return func(a A) func(B) R { return func(b B) R { return f.Apply(a, b) } }
}

// Apply2 is a function that takes two arguments.
func Apply2[A, B, R any](f func(A, B) R, a A, b B) R {
	return Fn2[A, B, R](f).Apply(a, b)
}

// Curry2 is a function that curry two arguments function.
func Curry2[A, B, R any](f func(A, B) R) func(A) func(B) R {
	return Fn2[A, B, R](f).Curry()
}

// Uncurry2 is a function that uncurry two arguments function.
func Uncurry2[A, B, R any](f func(A) func(B) R) func(A, B) R {
	return func(a A, b B) R { return f(a)(b) }
}
