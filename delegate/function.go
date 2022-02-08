package delegate

type Function0[Out any] func() Out
type Function1[I1, Out any] func(I1) Out
type Function2[I1, I2, Out any] func(I1, I2) Out
type Function3[I1, I2, I3, Out any] func(I1, I2, I3) Out
type Function4[I1, I2, I3, I4, Out any] func(I1, I2, I3, I4) Out
type Function5[I1, I2, I3, I4, I5, Out any] func(I1, I2, I3, I4, I5) Out
type Function6[I1, I2, I3, I4, I5, I6, Out any] func(I1, I2, I3, I4, I5, I6) Out
type Function7[I1, I2, I3, I4, I5, I6, I7, Out any] func(I1, I2, I3, I4, I5, I6, I7) Out
type Function8[I1, I2, I3, I4, I5, I6, I7, I8, Out any] func(I1, I2, I3, I4, I5, I6, I7, I8) Out
type Function9[I1, I2, I3, I4, I5, I6, I7, I8, I9, Out any] func(I1, I2, I3, I4, I5, I6, I7, I8, I9) Out

func ChainFunction1[A, B, C any](f1 Function1[A, B], f2 Function1[B, C]) Function1[A, C] {
	return func(a A) C { return f2(f1(a)) }
}
