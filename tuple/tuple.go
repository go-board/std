package tuple

import (
	"fmt"
)

// Pair is a tuple of two elements.
type Pair[A, B any] struct {
	a A
	b B
}

func (self Pair[A, B]) String() string {
	return fmt.Sprintf("Pair(%+v, %+v)", self.a, self.b)
}

func NewPair[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{a: a, b: b}
}

func PairFromA[A, B any](a A) Pair[A, B] {
	return Pair[A, B]{a: a}
}

func PairFromB[A, B any](b B) Pair[A, B] {
	return Pair[A, B]{b: b}
}

func (self Pair[A, B]) First() A  { return self.a }
func (self Pair[A, B]) Second() B { return self.b }

// Triple is a tuple of three elements.
type Triple[A, B, C any] struct {
	a A
	b B
	c C
}

func (self Triple[A, B, C]) String() string {
	return fmt.Sprintf("Triple(%+v, %+v, %+v)", self.a, self.b, self.c)
}

func NewTriple[A, B, C any](a A, b B, c C) Triple[A, B, C] {
	return Triple[A, B, C]{a: a, b: b, c: c}
}

func (self Triple[A, B, C]) First() A  { return self.a }
func (self Triple[A, B, C]) Second() B { return self.b }
func (self Triple[A, B, C]) Third() C  { return self.c }

// Tuple4 is a tuple of four elements.
type Tuple4[A, B, C, D any] struct {
	a A
	b B
	c C
	d D
}

func (self Tuple4[A, B, C, D]) String() string {
	return fmt.Sprintf("Tuple4(%+v, %+v, %+v, %+v)", self.a, self.b, self.c, self.d)
}

func NewTuple4[A, B, C, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	return Tuple4[A, B, C, D]{a: a, b: b, c: c, d: d}
}

func (self Tuple4[A, B, C, D]) First() A  { return self.a }
func (self Tuple4[A, B, C, D]) Second() B { return self.b }
func (self Tuple4[A, B, C, D]) Third() C  { return self.c }
func (self Tuple4[A, B, C, D]) Fourth() D { return self.d }

// Tuple5 is a tuple of five elements.
type Tuple5[A, B, C, D, E any] struct {
	a A
	b B
	c C
	d D
	e E
}

func (self Tuple5[A, B, C, D, E]) String() string {
	return fmt.Sprintf("Tuple5(%+v, %+v, %+v, %+v, %+v)", self.a, self.b, self.c, self.d, self.e)
}

func NewTuple5[A, B, C, D, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	return Tuple5[A, B, C, D, E]{a: a, b: b, c: c, d: d, e: e}
}

func (self Tuple5[A, B, C, D, E]) First() A  { return self.a }
func (self Tuple5[A, B, C, D, E]) Second() B { return self.b }
func (self Tuple5[A, B, C, D, E]) Third() C  { return self.c }
func (self Tuple5[A, B, C, D, E]) Fourth() D { return self.d }
func (self Tuple5[A, B, C, D, E]) Fifth() E  { return self.e }
