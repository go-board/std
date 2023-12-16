package stack

import (
	"github.com/go-board/std/iter"
)

type Stack[E any] struct{ elements []E }

// FromIter create a Stack from [iter.Seq].
func FromIter[E any](it iter.Seq[E]) *Stack[E] {
	s := new(Stack[E])
	iter.CollectFunc(it, func(e E) bool { s.Push(e); return true })
	return s
}

// FromSlice create a Stack from variadic elements.
func FromSlice[E any](elems ...E) *Stack[E] {
	return &Stack[E]{elements: elems}
}

// Push an element to Stack top.
func (s *Stack[E]) Push(elem E) { s.elements = append(s.elements, elem) }

// Pop an element from Stack top.
//
// Pop returns two results, the first is element and the second is a bool
// that indicate the value whether is valid.
func (s *Stack[E]) Pop() (e E, ok bool) {
	if len(s.elements) == 0 {
		return
	}
	e = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	ok = true
	return
}

// Peek returns the top element, without pop it.
func (s *Stack[E]) Peek() (e E, ok bool) {
	if len(s.elements) == 0 {
		return
	}
	e = s.elements[len(s.elements)-1]
	ok = true
	return
}

// IsEmpty check whether Stack is empty.
func (s *Stack[E]) IsEmpty() bool { return s.Size() == 0 }

// Size returns size of stack.
func (s *Stack[E]) Size() int { return len(s.elements) }

// Iter make Stack to [iter.Seq] in pop order.
func (s *Stack[E]) Iter() iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := len(s.elements) - 1; i >= 0; i-- {
			if !yield(s.elements[i]) {
				break
			}
		}
	}
}
