package core

import "github.com/go-board/std/hash"

func (self Byte) Hash(state hash.Hasher)    { state.WriteUint8(uint8(self)) }
func (self Char) Hash(state hash.Hasher)    { state.WriteUint32(uint32(self)) }
func (self Int) Hash(state hash.Hasher)     { state.WriteInt(int(self)) }
func (self Int8) Hash(state hash.Hasher)    { state.WriteInt8(int8(self)) }
func (self Int16) Hash(state hash.Hasher)   { state.WriteInt16(int16(self)) }
func (self Int32) Hash(state hash.Hasher)   { state.WriteInt32(int32(self)) }
func (self Int64) Hash(state hash.Hasher)   { state.WriteInt64(int64(self)) }
func (self Uint) Hash(state hash.Hasher)    { state.WriteUint(uint(self)) }
func (self Uint8) Hash(state hash.Hasher)   { state.WriteUint8(uint8(self)) }
func (self Uint16) Hash(state hash.Hasher)  { state.WriteUint16(uint16(self)) }
func (self Uint32) Hash(state hash.Hasher)  { state.WriteUint32(uint32(self)) }
func (self Uint64) Hash(state hash.Hasher)  { state.WriteUint64(uint64(self)) }
func (self Float32) Hash(state hash.Hasher) { state.WriteFloat32(float32(self)) }
func (self Float64) Hash(state hash.Hasher) { state.WriteFloat64(float64(self)) }
func (self Bool) Hash(state hash.Hasher)    { state.WriteBool(bool(self)) }
func (self String) Hash(state hash.Hasher)  { state.Write([]byte(self)) }

func (self Complex64) Hash(state hash.Hasher) {
	state.WriteFloat64(float64(self.Real()))
	state.WriteFloat64(float64(self.Imag()))
}

func (self Complex32) Hash(state hash.Hasher) {
	state.WriteFloat32(float32(self.Real()))
	state.WriteFloat32(float32(self.Imag()))
}
