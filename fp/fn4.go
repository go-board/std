package fp

// Fn4 is a function that takes four arguments.
type Fn4[A, B, C, D, R any] func(A, B, C, D) R

func (f Fn4[A, B, C, D, R]) Apply(a A, b B, c C, d D) R { return f(a, b, c, d) }

func (f Fn4[A, B, C, D, R]) Curry() Fn1[A, Fn1[B, Fn1[C, Fn1[D, R]]]] {
	return func(a A) Fn1[B, Fn1[C, Fn1[D, R]]] {
		return func(b B) Fn1[C, Fn1[D, R]] {
			return func(c C) Fn1[D, R] {
				return func(d D) R {
					return f.Apply(a, b, c, d)
				}
			}
		}
	}
}

func (f Fn4[A, B, C, D, R]) Partial1(a A) Fn3[B, C, D, R] {
	return func(b B, c C, d D) R {
		return f.Apply(a, b, c, d)
	}
}

func (f Fn4[A, B, C, D, R]) Partial2(a A, b B) Fn2[C, D, R] {
	return func(c C, d D) R { return f.Apply(a, b, c, d) }
}

func (f Fn4[A, B, C, D, R]) Partial3(a A, b B, c C) Fn1[D, R] {
	return func(d D) R { return f.Apply(a, b, c, d) }
}

func MakeFn4[A, B, C, D, R any](f func(A, B, C, D) R) Fn4[A, B, C, D, R] {
	return f
}

func Curry4[A, B, C, D, R any](f func(A, B, C, D) R) Fn1[A, Fn1[B, Fn1[C, Fn1[D, R]]]] {
	return MakeFn4(f).Curry()
}

// Uncurry4 is a function that uncurry four argument function.
func Uncurry4[A, B, C, D, R any](f func(A) func(B) func(C) func(D) R) Fn4[A, B, C, D, R] {
	return func(a A, b B, c C, d D) R { return f(a)(b)(c)(d) }
}
