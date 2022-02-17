package hash_test

import (
	"testing"

	"github.com/go-board/std/hash"
)

func TestBytesLike(t *testing.T) {
	x := hash.BytesLike("Hello,world")
	y := hash.BytesLike([]byte("Hello,world"))
	if x != y {
		t.Fail()
	}
	t.Logf("%v", x)
}

type User struct {
	ID   int64
	Name string
	Age  uint8
}

func (self *User) Hash(state hash.Hasher) {
	state.WriteInt64(self.ID)
	state.Write([]byte(self.Name))
	state.WriteUint8(self.Age)
}

func TestHash(t *testing.T) {
	u1 := &User{1, "Alice", 12}
	u2 := &User{1, "Alice", 12}
	u3 := &User{2, "Alice", 13}
	if hash.Hash(u1) != hash.Hash(u2) {
		t.Fail()
	}
	if hash.Hash(u1) == hash.Hash(u3) {
		t.Fail()
	}
}
