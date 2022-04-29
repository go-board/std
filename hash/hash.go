package hash

import (
	"encoding/binary"
	"hash"
	"hash/fnv"
)

type Hashable interface {
	Hash(state Hasher)
}

type Hasher interface {
	Finish() uint64

	Write(data []byte)
	WriteInt(i int)
	WriteInt8(i int8)
	WriteInt16(i int16)
	WriteInt32(i int32)
	WriteInt64(i int64)
	WriteUint(i uint)
	WriteUint8(i uint8)
	WriteUint16(i uint16)
	WriteUint32(i uint32)
	WriteUint64(i uint64)
	WriteFloat32(f float32)
	WriteFloat64(f float64)
	WriteBool(v bool)
}

type baseHasher struct {
	x         hash.Hash64
	byteOrder binary.ByteOrder
}

func newBaseHasher() *baseHasher {
	return &baseHasher{x: fnv.New64(), byteOrder: binary.LittleEndian}
}

var _ Hasher = (*baseHasher)(nil)

func (self *baseHasher) Finish() uint64         { return self.x.Sum64() }
func (self *baseHasher) Write(data []byte)      { binary.Write(self.x, self.byteOrder, data) }
func (self *baseHasher) WriteInt(i int)         { binary.Write(self.x, self.byteOrder, int64(i)) }
func (self *baseHasher) WriteInt8(i int8)       { binary.Write(self.x, self.byteOrder, i) }
func (self *baseHasher) WriteInt16(i int16)     { binary.Write(self.x, self.byteOrder, i) }
func (self *baseHasher) WriteInt32(i int32)     { binary.Write(self.x, self.byteOrder, i) }
func (self *baseHasher) WriteInt64(i int64)     { binary.Write(self.x, self.byteOrder, i) }
func (self *baseHasher) WriteUint(i uint)       { binary.Write(self.x, self.byteOrder, uint64(i)) }
func (self *baseHasher) WriteUint8(i uint8)     { binary.Write(self.x, self.byteOrder, i) }
func (self *baseHasher) WriteUint16(i uint16)   { binary.Write(self.x, self.byteOrder, i) }
func (self *baseHasher) WriteUint32(i uint32)   { binary.Write(self.x, self.byteOrder, i) }
func (self *baseHasher) WriteUint64(i uint64)   { binary.Write(self.x, self.byteOrder, i) }
func (self *baseHasher) WriteFloat32(f float32) { binary.Write(self.x, self.byteOrder, f) }
func (self *baseHasher) WriteFloat64(f float64) { binary.Write(self.x, self.byteOrder, f) }
func (self *baseHasher) WriteBool(v bool)       { binary.Write(self.x, self.byteOrder, v) }

func Hash[H Hashable](h H) uint64 {
	state := newBaseHasher()
	h.Hash(state)
	return state.Finish()
}

func HashSlice[H Hashable, HS ~[]H](hs HS) uint64 {
	state := newBaseHasher()
	for _, h := range hs {
		h.Hash(state)
	}
	return state.Finish()
}

func Int64(x int64) uint64 {
	state := newBaseHasher()
	state.WriteInt64(x)
	return state.Finish()
}

func Uint64(x uint64) uint64 {
	state := newBaseHasher()
	state.WriteUint64(x)
	return state.Finish()
}

func Float64(x float64) uint64 {
	state := newBaseHasher()
	state.WriteFloat64(x)
	return state.Finish()
}

func Bool(x bool) uint64 {
	state := newBaseHasher()
	state.WriteBool(x)
	return state.Finish()
}

func Complex64(c complex128) uint64 {
	state := newBaseHasher()
	state.WriteFloat64(real(c))
	state.WriteFloat64(imag(c))
	return state.Finish()
}

func BytesLike[H ~string | ~[]byte](x H) uint64 {
	state := newBaseHasher()
	state.Write([]byte(x))
	return state.Finish()
}
