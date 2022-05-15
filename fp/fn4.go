package fp

// Function4 is a function that takes four arguments.
type Function4[A, B, C, D, R any] interface {
	Apply(A, B, C, D) R
	Curry() func(A) func(B) func(C) func(D) R
}

// Fn4 is a function that takes four arguments.
type Fn4[A, B, C, D, R any] func(A, B, C, D) R

func (f Fn4[A, B, C, D, R]) Apply(a A, b B, c C, d D) R { return f(a, b, c, d) }

func (f Fn4[A, B, C, D, R]) Curry() func(A) func(B) func(C) func(D) R {
	return func(a A) func(B) func(C) func(D) R {
		return func(b B) func(C) func(D) R {
			return func(c C) func(D) R { return func(d D) R { return f.Apply(a, b, c, d) } }
		}
	}
}

// Apply4 is a function that takes four arguments.
func Apply4[A, B, C, D, R any](f func(A, B, C, D) R, a A, b B, c C, d D) R {
	return Fn4[A, B, C, D, R](f).Apply(a, b, c, d)
}

// Curry4 is a function that curry four argument function.
func Curry4[A, B, C, D, R any](f func(A, B, C, D) R) func(A) func(B) func(C) func(D) R {
	return Fn4[A, B, C, D, R](f).Curry()
}

// Uncurry4 is a function that uncurry four argument function.
func Uncurry4[A, B, C, D, R any](f func(A) func(B) func(C) func(D) R) func(A, B, C, D) R {
	return func(a A, b B, c C, d D) R { return f(a)(b)(c)(d) }
}
