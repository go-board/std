package core

type Primitive interface{ Numeric | ~string | ~bool }

type ByteSeq interface{ ~[]byte | ~string }

type Float interface{ ~float32 | ~float64 }

type Complex interface{ ~complex64 | ~complex128 }

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface{ Signed | Unsigned }

type Numeric interface{ Integer | Float | Complex }

type Ordered interface{ Integer | Float | ~string }

// Deprecated: use [Numeric] instead
type Number = Numeric
