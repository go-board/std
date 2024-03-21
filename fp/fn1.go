package fp

// Fn1 is a function that takes one argument.
type Fn1[A, R any] func(A) R

func (f Fn1[A, R]) Apply(a A) R { return f(a) }

func (f Fn1[A, R]) Curry() Fn1[A, R] { return func(a A) R { return f.Apply(a) } }

func MakeFn1[A, R any](f func(A) R) Fn1[A, R] { return f }

func Curry1[A, R any](f func(A) R) Fn1[A, R] { return f }

// Uncurry1 is a function that uncurry one argument function.
func Uncurry1[A, R any](f func(A) R) func(A) R {
	return func(a A) R { return f(a) }
}
