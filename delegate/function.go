package delegate

// Deprecated: Use func() Out instead.
type Function0[Out any] func() Out

// Deprecated: Use fp.Fn1 instead.
type Function1[I1, Out any] func(I1) Out

// Deprecated: Use fp.Fn2 instead.
type Function2[I1, I2, Out any] func(I1, I2) Out

// Deprecated: Use fp.Fn3 instead.
type Function3[I1, I2, I3, Out any] func(I1, I2, I3) Out

// Deprecated: Use fp.Fn4 instead.
type Function4[I1, I2, I3, I4, Out any] func(I1, I2, I3, I4) Out

// Deprecated: Use fp.Fn5 instead.
type Function5[I1, I2, I3, I4, I5, Out any] func(I1, I2, I3, I4, I5) Out

// Deprecated: Use raw func instead.
type Function6[I1, I2, I3, I4, I5, I6, Out any] func(I1, I2, I3, I4, I5, I6) Out

// Deprecated: Use raw func instead.
type Function7[I1, I2, I3, I4, I5, I6, I7, Out any] func(I1, I2, I3, I4, I5, I6, I7) Out

// Deprecated: Use raw func instead.
type Function8[I1, I2, I3, I4, I5, I6, I7, I8, Out any] func(I1, I2, I3, I4, I5, I6, I7, I8) Out

// Deprecated: Use raw func instead.
type Function9[I1, I2, I3, I4, I5, I6, I7, I8, I9, Out any] func(I1, I2, I3, I4, I5, I6, I7, I8, I9) Out

// Deprecated: Use fp.Compose2 instead.
func ChainFunction1[A, B, C any](f1 Function1[A, B], f2 Function1[B, C]) Function1[A, C] {
	return func(a A) C { return f2(f1(a)) }
}
