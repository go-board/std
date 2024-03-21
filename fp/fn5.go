package fp

// Fn5 is a function that takes five arguments.
type Fn5[A, B, C, D, E, R any] func(A, B, C, D, E) R

func (f Fn5[A, B, C, D, E, R]) Apply(a A, b B, c C, d D, e E) R { return f(a, b, c, d, e) }

func (f Fn5[A, B, C, D, E, R]) Curry() Fn1[A, Fn1[B, Fn1[C, Fn1[D, Fn1[E, R]]]]] {
	return func(a A) Fn1[B, Fn1[C, Fn1[D, Fn1[E, R]]]] {
		return func(b B) Fn1[C, Fn1[D, Fn1[E, R]]] {
			return func(c C) Fn1[D, Fn1[E, R]] {
				return func(d D) Fn1[E, R] {
					return func(e E) R {
						return f.Apply(a, b, c, d, e)
					}
				}
			}
		}
	}
}

func (f Fn5[A, B, C, D, E, R]) Partial1(a A) Fn4[B, C, D, E, R] {
	return func(b B, c C, d D, e E) R {
		return f.Apply(a, b, c, d, e)
	}
}

func (f Fn5[A, B, C, D, E, R]) Partial2(a A, b B) Fn3[C, D, E, R] {
	return func(c C, d D, e E) R { return f.Apply(a, b, c, d, e) }
}

func (f Fn5[A, B, C, D, E, R]) Partial3(a A, b B, c C) Fn2[D, E, R] {
	return func(d D, e E) R { return f.Apply(a, b, c, d, e) }
}

func (f Fn5[A, B, C, D, E, R]) Partial4(a A, b B, c C, d D) Fn1[E, R] {
	return func(e E) R { return f.Apply(a, b, c, d, e) }
}

func MakeFn5[A, B, C, D, E, R any](f func(A, B, C, D, E) R) Fn5[A, B, C, D, E, R] {
	return f
}

// Uncurry5 is a function that uncurry five argument function.
func Uncurry5[A, B, C, D, E, R any](f func(A) func(B) func(C) func(D) func(E) R) Fn5[A, B, C, D, E, R] {
	return func(a A, b B, c C, d D, e E) R { return f(a)(b)(c)(d)(e) }
}
