package tuple

import (
	"fmt"
)

// Pair is a tuple of two elements.
type Pair[A, B any] struct {
	First  A
	Second B
}

func (self Pair[A, B]) String() string {
	return fmt.Sprintf("Pair(%+v, %+v)", self.First, self.Second)
}

func NewPair[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{First: a, Second: b}
}

func PairOf[A, B any](a A, b B) Pair[A, B] { return Pair[A, B]{First: a, Second: b} }

// Triple is a tuple of three elements.
type Triple[A, B, C any] struct {
	First  A
	Second B
	Third  C
}

func (self Triple[A, B, C]) String() string {
	return fmt.Sprintf("Triple(%+v, %+v, %+v)", self.First, self.Second, self.Third)
}

func MakeTriple[A, B, C any](a A, b B, c C) Triple[A, B, C] {
	return Triple[A, B, C]{First: a, Second: b, Third: c}
}

func TripleOf[A, B, C any](a A, b B, c C) Triple[A, B, C] { return MakeTriple(a, b, c) }

// Tuple4 is a tuple of four elements.
type Tuple4[A, B, C, D any] struct {
	First  A
	Second B
	Third  C
	Fourth D
}

func (self Tuple4[A, B, C, D]) String() string {
	return fmt.Sprintf("Tuple4(%+v, %+v, %+v, %+v)", self.First, self.Second, self.Third, self.Fourth)
}

func MakeTuple4[A, B, C, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	return Tuple4[A, B, C, D]{First: a, Second: b, Third: c, Fourth: d}
}

// Tuple5 is a tuple of five elements.
type Tuple5[A, B, C, D, E any] struct {
	First  A
	Second B
	Third  C
	Fourth D
	Fifth  E
}

func (self Tuple5[A, B, C, D, E]) String() string {
	return fmt.Sprintf("Tuple5(%+v, %+v, %+v, %+v, %+v)", self.First, self.Second, self.Third, self.Fourth, self.Fifth)
}

func MakeTuple5[A, B, C, D, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	return Tuple5[A, B, C, D, E]{First: a, Second: b, Third: c, Fourth: d, Fifth: e}
}

// Tuple6 is a tuple of six elements.
type Tuple6[A, B, C, D, E, F any] struct {
	First  A
	Second B
	Third  C
	Fourth D
	Fifth  E
	Sixth  F
}

func (self Tuple6[A, B, C, D, E, F]) String() string {
	return fmt.Sprintf("Tuple6(%+v, %+v, %+v, %+v, %+v, %+v)", self.First, self.Second, self.Third, self.Fourth, self.Fifth, self.Sixth)
}

func MakeTuple6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F] {
	return Tuple6[A, B, C, D, E, F]{First: a, Second: b, Third: c, Fourth: d, Fifth: e, Sixth: f}
}

// Tuple7 is a tuple of seven elements.
type Tuple7[A, B, C, D, E, F, G any] struct {
	First   A
	Second  B
	Third   C
	Fourth  D
	Fifth   E
	Sixth   F
	Seventh G
}

func (self Tuple7[A, B, C, D, E, F, G]) String() string {
	return fmt.Sprintf("Tuple7(%+v, %+v, %+v, %+v, %+v, %+v, %+v)", self.First, self.Second, self.Third, self.Fourth, self.Fifth, self.Sixth, self.Seventh)
}

func MakeTuple7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) Tuple7[A, B, C, D, E, F, G] {
	return Tuple7[A, B, C, D, E, F, G]{First: a, Second: b, Third: c, Fourth: d, Fifth: e, Sixth: f, Seventh: g}
}

// Tuple8 is a tuple of eight elements.
type Tuple8[A, B, C, D, E, F, G, H any] struct {
	First   A
	Second  B
	Third   C
	Fourth  D
	Fifth   E
	Sixth   F
	Seventh G
	Eighth  H
}

func (self Tuple8[A, B, C, D, E, F, G, H]) String() string {
	return fmt.Sprintf("Tuple8(%+v, %+v, %+v, %+v, %+v, %+v, %+v, %+v)", self.First, self.Second, self.Third, self.Fourth, self.Fifth, self.Sixth, self.Seventh, self.Eighth)
}

func MakeTuple8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tuple8[A, B, C, D, E, F, G, H] {
	return Tuple8[A, B, C, D, E, F, G, H]{First: a, Second: b, Third: c, Fourth: d, Fifth: e, Sixth: f, Seventh: g, Eighth: h}
}

// Tuple9 is a tuple of nine elements.
type Tuple9[A, B, C, D, E, F, G, H, I any] struct {
	First   A
	Second  B
	Third   C
	Fourth  D
	Fifth   E
	Sixth   F
	Seventh G
	Eighth  H
	Ninth   I
}

func (self Tuple9[A, B, C, D, E, F, G, H, I]) String() string {
	return fmt.Sprintf("Tuple9(%+v, %+v, %+v, %+v, %+v, %+v, %+v, %+v, %+v)", self.First, self.Second, self.Third, self.Fourth, self.Fifth, self.Sixth, self.Seventh, self.Eighth, self.Ninth)
}

func MakeTuple9[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Tuple9[A, B, C, D, E, F, G, H, I] {
	return Tuple9[A, B, C, D, E, F, G, H, I]{First: a, Second: b, Third: c, Fourth: d, Fifth: e, Sixth: f, Seventh: g, Eighth: h, Ninth: i}
}
