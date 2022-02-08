package delegate

type (
	Consumer1[I1 any]                                 func(I1)
	Consumer2[I1, I2 any]                             func(I1, I2)
	Consumer3[I1, I2, I3 any]                         func(I1, I2, I3)
	Consumer4[I1, I2, I3, I4 any]                     func(I1, I2, I3, I4)
	Consumer5[I1, I2, I3, I4, I5 any]                 func(I1, I2, I3, I4, I5)
	Consumer6[I1, I2, I3, I4, I5, I6 any]             func(I1, I2, I3, I4, I5, I6)
	Consumer7[I1, I2, I3, I4, I5, I6, I7 any]         func(I1, I2, I3, I4, I5, I6, I7)
	Consumer8[I1, I2, I3, I4, I5, I6, I7, I8 any]     func(I1, I2, I3, I4, I5, I6, I7, I8)
	Consumer9[I1, I2, I3, I4, I5, I6, I7, I8, I9 any] func(I1, I2, I3, I4, I5, I6, I7, I8, I9)
)
