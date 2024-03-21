package fp

// Fn2 is a function that takes two arguments.
type Fn2[A, B, R any] func(A, B) R

func (f Fn2[A, B, R]) Apply(a A, b B) R { return f(a, b) }

func (f Fn2[A, B, R]) Curry() Fn1[A, Fn1[B, R]] {
	return func(a A) Fn1[B, R] {
		return func(b B) R {
			return f.Apply(a, b)
		}
	}
}

func (f Fn2[A, B, R]) Partial1(a A) Fn1[B, R] {
	return func(b B) R { return f.Apply(a, b) }
}

func MakeFn2[A, B, R any](f func(A, B) R) Fn2[A, B, R] { return f }

func Curry2[A, B, R any](f func(A, B) R) Fn1[A, Fn1[B, R]] { return MakeFn2(f).Curry() }

// Uncurry2 is a function that uncurry two arguments function.
func Uncurry2[A, B, R any](f func(A) func(B) R) Fn2[A, B, R] {
	return func(a A, b B) R { return f(a)(b) }
}
