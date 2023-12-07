package stack

import (
	"github.com/go-board/std/iter"
	"github.com/go-board/std/slices"
)

type Stack[E any] struct {
	elements []E
}

func FromIter[E any](it iter.Seq[E]) *Stack[E] {
	return &Stack[E]{elements: slices.Collect(it)}
}

func FromSlice[E any](elems ...E) *Stack[E] {
	return &Stack[E]{elements: elems}
}

func (s *Stack[E]) Push(elem E) { s.elements = append(s.elements, elem) }

func (s *Stack[E]) Pop() (e E, ok bool) {
	if len(s.elements) == 0 {
		return
	}
	e = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	ok = true
	return
}

func (s *Stack[E]) Peek() (e E, ok bool) {
	if len(s.elements) == 0 {
		return
	}
	e = s.elements[len(s.elements)-1]
	ok = true
	return
}

func (s *Stack[E]) IsEmpty() bool     { return len(s.elements) == 0 }
func (s *Stack[E]) Size() int         { return len(s.elements) }
func (s *Stack[E]) Iter() iter.Seq[E] { return slices.Backward(s.elements) }
