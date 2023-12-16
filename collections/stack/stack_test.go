package stack_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-board/std/collections/stack"
	"github.com/go-board/std/slices"
)

func TestFromIter(t *testing.T) {
	s := stack.FromIter(slices.Forward([]int{1, 2, 3}))
	qt.Assert(t, s.Size(), qt.Equals, 3)
	qt.Assert(t, slices.Collect(s.Iter()), qt.DeepEquals, []int{3, 2, 1})
}

func TestStack_Iter(t *testing.T) {
	s := stack.FromSlice(1, 2, 3, 4).Iter()
	qt.Assert(t, slices.Collect(s), qt.DeepEquals, []int{4, 3, 2, 1})
}

func TestStack_Push(t *testing.T) {
	s := stack.FromSlice[int]()
	s.Push(1)
	qt.Assert(t, slices.Collect(s.Iter()), qt.DeepEquals, []int{1})
	s.Push(2)
	qt.Assert(t, slices.Collect(s.Iter()), qt.DeepEquals, []int{2, 1})
}

func TestStack_Peek(t *testing.T) {
	s := stack.FromSlice[int]()
	_, ok := s.Peek()
	qt.Assert(t, ok, qt.IsFalse)
	s.Push(1)
	e1, ok := s.Peek()
	qt.Assert(t, ok, qt.IsTrue)
	qt.Assert(t, e1, qt.Equals, 1)
	s.Push(2)
	e2, ok := s.Peek()
	qt.Assert(t, ok, qt.IsTrue)
	qt.Assert(t, e2, qt.Equals, 2)
}

func TestStack_Pop(t *testing.T) {
	s := stack.FromSlice[int]()
	_, ok := s.Pop()
	qt.Assert(t, ok, qt.IsFalse)
	s.Push(1)
	e1, ok := s.Pop()
	qt.Assert(t, ok, qt.IsTrue)
	qt.Assert(t, e1, qt.Equals, 1)
	qt.Assert(t, s.IsEmpty(), qt.IsTrue)
	s.Push(2)
	s.Push(3)
	e2, ok := s.Pop()
	qt.Assert(t, ok, qt.IsTrue)
	qt.Assert(t, e2, qt.Equals, 3)
}

func TestStack_IsEmpty(t *testing.T) {
	s := stack.FromSlice[int]()
	qt.Assert(t, s.IsEmpty(), qt.IsTrue)
	s.Push(1)
	qt.Assert(t, s.IsEmpty(), qt.IsFalse)
	s.Pop()
	qt.Assert(t, s.IsEmpty(), qt.IsTrue)
}
