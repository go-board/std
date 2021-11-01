package sets

import (
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	s := NewHashSet[int]()
	s.Add(1, 2)
	if !s.Contains(1) {
		t.Fatalf("element 1 must in set")
	}
	if !s.Contains(2) {
		t.Fatalf("element 2 must in set")
	}
}
