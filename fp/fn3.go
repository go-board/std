package fp

// Fn3 is a function that takes three arguments.
type Fn3[A, B, C, R any] func(A, B, C) R

func (f Fn3[A, B, C, R]) Apply(a A, b B, c C) R { return f(a, b, c) }

func (f Fn3[A, B, C, R]) Curry() Fn1[A, Fn1[B, Fn1[C, R]]] {
	return func(a A) Fn1[B, Fn1[C, R]] {
		return func(b B) Fn1[C, R] {
			return func(c C) R {
				return f.Apply(a, b, c)
			}
		}
	}
}

func (f Fn3[A, B, C, R]) Partial1(a A) Fn2[B, C, R] {
	return func(b B, c C) R { return f.Apply(a, b, c) }
}

func (f Fn3[A, B, C, R]) Partial2(a A, b B) Fn1[C, R] {
	return func(c C) R { return f.Apply(a, b, c) }
}

func MakeFn3[A, B, C, R any](f func(A, B, C) R) Fn3[A, B, C, R] {
	return f
}

func Curry3[A, B, C, R any](f func(A, B, C) R) Fn1[A, Fn1[B, Fn1[C, R]]] {
	return MakeFn3(f).Curry()
}

// Uncurry3 is a function that uncurry three arguments function.
func Uncurry3[A, B, C, R any](f func(A) func(B) func(C) R) Fn3[A, B, C, R] {
	return func(a A, b B, c C) R { return f(a)(b)(c) }
}
