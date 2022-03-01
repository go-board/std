package core

type Unit struct{}

type (
	Byte byte
	Char rune
)

type (
	Int   int
	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64
)

type (
	Uint   uint
	Uint8  uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64
)

type (
	Float32 float32
	Float64 float64
)

type Bool bool

type String string

type (
	Complex64 complex128
	Complex32 complex64
)

type (
	Slice[T any]             []T
	Map[K comparable, V any] map[K]V
	Chan[T any]              chan T
)

func (self Complex32) Real() Float32 { return Float32(real(self)) }
func (self Complex32) Imag() Float32 { return Float32(imag(self)) }
func (self Complex64) Real() Float64 { return Float64(real(self)) }
func (self Complex64) Imag() Float64 { return Float64(imag(self)) }

func IsPrimitive[T any]() bool {
	var t0 T
	switch any(t0).(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool, string, complex64, complex128:
		return true
	case Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Float32, Float64, Bool, String, Complex64, Complex32:
		return true
	}
	return false
}
