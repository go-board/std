package fp

// Function1 is a function that takes one argument.
type Function1[A, R any] interface {
	Apply(A) R
	Curry() func(A) R
}

// Fn1 is a function that takes one argument.
type Fn1[A, R any] func(A) R

func (f Fn1[A, R]) Apply(a A) R { return f(a) }

func (f Fn1[A, R]) Curry() func(A) R { return Fn1[A, R](func(a A) R { return f.Apply(a) }) }

// Apply1 is a function that apply one argument to a function.
func Apply1[A, R any](f func(A) R, a A) R {
	return Fn1[A, R](f).Apply(a)
}

// Curry1 is a function that curry one argument function.
func Curry1[A, R any](f func(A) R) func(A) R {
	return Fn1[A, R](f).Curry()
}

// Uncurry1 is a function that uncurry one argument function.
func Uncurry1[A, R any](f func(A) R) func(A) R {
	return func(a A) R { return f(a) }
}
