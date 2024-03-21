package fp

func Identity[T any](x T) T { return x }

func Flip[A, B, R any](f func(A, B) R) func(B, A) R {
	return func(b B, a A) R {
		return f(a, b)
	}
}
