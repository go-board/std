package core

type Primitive interface{ Number | ~string | ~bool }
type ByteSequence interface{ ~[]byte | ~string }
type Float interface{ ~float32 | ~float64 }
type Complex interface{ ~complex64 | ~complex128 }
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type Integer interface{ Signed | Unsigned }
type Number interface{ Integer | Float | Complex }

func Default[T Primitive]() T {
	var p T
	return p
}
