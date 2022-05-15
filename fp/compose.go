package fp

// Compose1 is a function that compose one function.
func Compose1[A, R any](f func(A) R) func(A) R {
	return func(a A) R { return f(a) }
}

// Compose2 is a function that compose two functions.
func Compose2[A, B, R any](f func(A) B, g func(B) R) func(A) R {
	return func(x A) R { return g(f(x)) }
}

// Compose3 is a function that compose three functions.
func Compose3[A, B, C, R any](f func(A) B, g func(B) C, h func(C) R) func(A) R {
	return func(x A) R { return h(g(f(x))) }
}

// Compose4 is a function that compose four functions.
func Compose4[A, B, C, D, R any](f func(A) B, g func(B) C, h func(C) D, i func(D) R) func(A) R {
	return func(x A) R { return i(h(g(f(x)))) }
}

// Compose5 is a function that compose five functions.
func Compose5[A, B, C, D, E, R any](f func(A) B, g func(B) C, h func(C) D, i func(D) E, j func(E) R) func(A) R {
	return func(x A) R { return j(i(h(g(f(x))))) }
}
