package fp

// Function5 is a function that takes five arguments.
type Function5[A, B, C, D, E, R any] interface {
	Apply(A, B, C, D, E) R
	Curry() func(A) func(B) func(C) func(D) func(E) R
}

// Fn5 is a function that takes five arguments.
type Fn5[A, B, C, D, E, R any] func(A, B, C, D, E) R

func (f Fn5[A, B, C, D, E, R]) Apply(a A, b B, c C, d D, e E) R { return f(a, b, c, d, e) }

func (f Fn5[A, B, C, D, E, R]) Curry() func(A) func(B) func(C) func(D) func(E) R {
	return func(a A) func(B) func(C) func(D) func(E) R {
		return func(b B) func(C) func(D) func(E) R {
			return func(c C) func(D) func(E) R {
				return func(d D) func(E) R { return func(e E) R { return f.Apply(a, b, c, d, e) } }
			}
		}
	}
}

// Apply5 is a function that takes five arguments.
func Apply5[A, B, C, D, E, R any](f func(A, B, C, D, E) R, a A, b B, c C, d D, e E) R {
	return Fn5[A, B, C, D, E, R](f).Apply(a, b, c, d, e)
}

// Curry5 is a function that curry five argument function.
func Curry5[A, B, C, D, E, R any](f func(A, B, C, D, E) R) func(A) func(B) func(C) func(D) func(E) R {
	return Fn5[A, B, C, D, E, R](f).Curry()
}

// Uncurry5 is a function that uncurry five argument function.
func Uncurry5[A, B, C, D, E, R any](f func(A) func(B) func(C) func(D) func(E) R) func(A, B, C, D, E) R {
	return func(a A, b B, c C, d D, e E) R { return f(a)(b)(c)(d)(e) }
}
