package tuple

import (
	"fmt"

	"github.com/go-board/std/cmp"
)

// Pair is a tuple of two elements.
type Pair[A, B any] struct {
	first  A
	second B
}

func (self Pair[A, B]) String() string {
	return fmt.Sprintf("Pair(%+v, %+v)", self.first, self.second)
}

func (self Pair[A, B]) Unpack() (A, B)    { return self.first, self.second }
func (self Pair[A, B]) Clone() Pair[A, B] { return MakePair(self.first, self.second) }
func (self Pair[A, B]) First() A          { return self.first }
func (self Pair[A, B]) Second() B         { return self.second }

// MakePair create a pair from a and b.
func MakePair[A, B any](a A, b B) Pair[A, B] { return Pair[A, B]{a, b} }

func PairComparator[A, B any](cmpA func(A, A) int, cmpB func(B, B) int) cmp.Comparator[Pair[A, B]] {
	return cmp.MakeComparatorFunc(func(lhs Pair[A, B], rhs Pair[A, B]) int {
		if c := cmpA(lhs.first, rhs.first); c != 0 {
			return c
		}
		return cmpB(lhs.second, rhs.second)
	})
}

// Triple is a tuple of three elements.
type Triple[A, B, C any] struct {
	first  A
	second B
	third  C
}

func (self Triple[A, B, C]) String() string {
	return fmt.Sprintf("Triple(%+v, %+v, %+v)", self.first, self.second, self.third)
}

func (self Triple[A, B, C]) Unpack() (A, B, C) {
	return self.first, self.second, self.third
}

func (self Triple[A, B, C]) Clone() Triple[A, B, C] {
	return MakeTriple(self.first, self.second, self.third)
}

func (self Triple[A, B, C]) First() A  { return self.first }
func (self Triple[A, B, C]) Second() B { return self.second }
func (self Triple[A, B, C]) Third() C  { return self.third }

func MakeTriple[A, B, C any](a A, b B, c C) Triple[A, B, C] {
	return Triple[A, B, C]{a, b, c}
}

// Tuple4 is a tuple of four elements.
type Tuple4[A, B, C, D any] struct {
	first  A
	second B
	third  C
	fourth D
}

func (self Tuple4[A, B, C, D]) String() string {
	return fmt.Sprintf("Tuple4(%+v, %+v, %+v, %+v)", self.first, self.second, self.third, self.fourth)
}

func (self Tuple4[A, B, C, D]) Unpack() (A, B, C, D) {
	return self.first, self.second, self.third, self.fourth
}

func (self Tuple4[A, B, C, D]) Clone() Tuple4[A, B, C, D] {
	return MakeTuple4(self.first, self.second, self.third, self.fourth)
}

func (self Tuple4[A, B, C, D]) First() A  { return self.first }
func (self Tuple4[A, B, C, D]) Second() B { return self.second }
func (self Tuple4[A, B, C, D]) Third() C  { return self.third }
func (self Tuple4[A, B, C, D]) Fourth() D { return self.fourth }

func MakeTuple4[A, B, C, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	return Tuple4[A, B, C, D]{a, b, c, d}
}

// Tuple5 is a tuple of five elements.
type Tuple5[A, B, C, D, E any] struct {
	first  A
	second B
	third  C
	fourth D
	fifth  E
}

func (self Tuple5[A, B, C, D, E]) String() string {
	return fmt.Sprintf("Tuple5(%+v, %+v, %+v, %+v, %+v)", self.first, self.second, self.third, self.fourth, self.fifth)
}

func (self Tuple5[A, B, C, D, E]) Unpack() (A, B, C, D, E) {
	return self.first, self.second, self.third, self.fourth, self.fifth
}

func (self Tuple5[A, B, C, D, E]) Clone() Tuple5[A, B, C, D, E] {
	return MakeTuple5(self.first, self.second, self.third, self.fourth, self.fifth)
}

func (self Tuple5[A, B, C, D, E]) First() A  { return self.first }
func (self Tuple5[A, B, C, D, E]) Second() B { return self.second }
func (self Tuple5[A, B, C, D, E]) Third() C  { return self.third }
func (self Tuple5[A, B, C, D, E]) Fourth() D { return self.fourth }
func (self Tuple5[A, B, C, D, E]) Fifth() E  { return self.fifth }

func MakeTuple5[A, B, C, D, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	return Tuple5[A, B, C, D, E]{a, b, c, d, e}
}

// Tuple6 is a tuple of six elements.
type Tuple6[A, B, C, D, E, F any] struct {
	first  A
	second B
	third  C
	fourth D
	fifth  E
	sixth  F
}

func (self Tuple6[A, B, C, D, E, F]) String() string {
	return fmt.Sprintf("Tuple6(%+v, %+v, %+v, %+v, %+v, %+v)", self.first, self.second, self.third, self.fourth, self.fifth, self.sixth)
}

func (self Tuple6[A, B, C, D, E, F]) Unpack() (A, B, C, D, E, F) {
	return self.first, self.second, self.third, self.fourth, self.fifth, self.sixth
}

func (self Tuple6[A, B, C, D, E, F]) Clone() Tuple6[A, B, C, D, E, F] {
	return MakeTuple6(self.first, self.second, self.third, self.fourth, self.fifth, self.sixth)
}

func (self Tuple6[A, B, C, D, E, F]) First() A  { return self.first }
func (self Tuple6[A, B, C, D, E, F]) Second() B { return self.second }
func (self Tuple6[A, B, C, D, E, F]) Third() C  { return self.third }
func (self Tuple6[A, B, C, D, E, F]) Fourth() D { return self.fourth }
func (self Tuple6[A, B, C, D, E, F]) Fifth() E  { return self.fifth }
func (self Tuple6[A, B, C, D, E, F]) Sixth() F  { return self.sixth }

func MakeTuple6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F] {
	return Tuple6[A, B, C, D, E, F]{a, b, c, d, e, f}
}

// Tuple7 is a tuple of seven elements.
type Tuple7[A, B, C, D, E, F, G any] struct {
	first   A
	second  B
	third   C
	fourth  D
	fifth   E
	sixth   F
	seventh G
}

func (self Tuple7[A, B, C, D, E, F, G]) String() string {
	return fmt.Sprintf("Tuple7(%+v, %+v, %+v, %+v, %+v, %+v, %+v)", self.first, self.second, self.third, self.fourth, self.fifth, self.sixth, self.seventh)
}

func (self Tuple7[A, B, C, D, E, F, G]) Unpack() (A, B, C, D, E, F, G) {
	return self.first, self.second, self.third, self.fourth, self.fifth, self.sixth, self.seventh
}

func (self Tuple7[A, B, C, D, E, F, G]) Clone() Tuple7[A, B, C, D, E, F, G] {
	return MakeTuple7(self.first, self.second, self.third, self.fourth, self.fifth, self.sixth, self.seventh)
}

func (self Tuple7[A, B, C, D, E, F, G]) First() A   { return self.first }
func (self Tuple7[A, B, C, D, E, F, G]) Second() B  { return self.second }
func (self Tuple7[A, B, C, D, E, F, G]) Third() C   { return self.third }
func (self Tuple7[A, B, C, D, E, F, G]) Fourth() D  { return self.fourth }
func (self Tuple7[A, B, C, D, E, F, G]) Fifth() E   { return self.fifth }
func (self Tuple7[A, B, C, D, E, F, G]) Sixth() F   { return self.sixth }
func (self Tuple7[A, B, C, D, E, F, G]) Seventh() G { return self.seventh }

func MakeTuple7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) Tuple7[A, B, C, D, E, F, G] {
	return Tuple7[A, B, C, D, E, F, G]{a, b, c, d, e, f, g}
}

// Tuple8 is a tuple of eight elements.
type Tuple8[A, B, C, D, E, F, G, H any] struct {
	first   A
	second  B
	third   C
	fourth  D
	fifth   E
	sixth   F
	seventh G
	eighth  H
}

func (self Tuple8[A, B, C, D, E, F, G, H]) String() string {
	return fmt.Sprintf("Tuple8(%+v, %+v, %+v, %+v, %+v, %+v, %+v, %+v)", self.first, self.second, self.third, self.fourth, self.fifth, self.sixth, self.seventh, self.eighth)
}

func (self Tuple8[A, B, C, D, E, F, G, H]) Unpack() (A, B, C, D, E, F, G, H) {
	return self.first, self.second, self.third, self.fourth, self.fifth, self.sixth, self.seventh, self.eighth
}

func (self Tuple8[A, B, C, D, E, F, G, H]) Clone() Tuple8[A, B, C, D, E, F, G, H] {
	return MakeTuple8(self.first, self.second, self.third, self.fourth, self.fifth, self.sixth, self.seventh, self.eighth)
}

func (self Tuple8[A, B, C, D, E, F, G, H]) First() A   { return self.first }
func (self Tuple8[A, B, C, D, E, F, G, H]) Second() B  { return self.second }
func (self Tuple8[A, B, C, D, E, F, G, H]) Third() C   { return self.third }
func (self Tuple8[A, B, C, D, E, F, G, H]) Fourth() D  { return self.fourth }
func (self Tuple8[A, B, C, D, E, F, G, H]) Fifth() E   { return self.fifth }
func (self Tuple8[A, B, C, D, E, F, G, H]) Sixth() F   { return self.sixth }
func (self Tuple8[A, B, C, D, E, F, G, H]) Seventh() G { return self.seventh }
func (self Tuple8[A, B, C, D, E, F, G, H]) Eighth() H  { return self.eighth }

func MakeTuple8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tuple8[A, B, C, D, E, F, G, H] {
	return Tuple8[A, B, C, D, E, F, G, H]{a, b, c, d, e, f, g, h}
}

// Tuple9 is a tuple of nine elements.
type Tuple9[A, B, C, D, E, F, G, H, I any] struct {
	first   A
	second  B
	third   C
	fourth  D
	fifth   E
	sixth   F
	seventh G
	eighth  H
	ninth   I
}

func (self Tuple9[A, B, C, D, E, F, G, H, I]) String() string {
	return fmt.Sprintf("Tuple9(%+v, %+v, %+v, %+v, %+v, %+v, %+v, %+v, %+v)", self.first, self.second, self.third, self.fourth, self.fifth, self.sixth, self.seventh, self.eighth, self.ninth)
}

func (self Tuple9[A, B, C, D, E, F, G, H, I]) Unpack() (A, B, C, D, E, F, G, H, I) {
	return self.first, self.second, self.third, self.fourth, self.fifth, self.sixth, self.seventh, self.eighth, self.ninth
}

func (self Tuple9[A, B, C, D, E, F, G, H, I]) Clone() Tuple9[A, B, C, D, E, F, G, H, I] {
	return MakeTuple9(self.first, self.second, self.third, self.fourth, self.fifth, self.sixth, self.seventh, self.eighth, self.ninth)
}

func (self Tuple9[A, B, C, D, E, F, G, H, I]) First() A   { return self.first }
func (self Tuple9[A, B, C, D, E, F, G, H, I]) Second() B  { return self.second }
func (self Tuple9[A, B, C, D, E, F, G, H, I]) Third() C   { return self.third }
func (self Tuple9[A, B, C, D, E, F, G, H, I]) Fourth() D  { return self.fourth }
func (self Tuple9[A, B, C, D, E, F, G, H, I]) Fifth() E   { return self.fifth }
func (self Tuple9[A, B, C, D, E, F, G, H, I]) Sixth() F   { return self.sixth }
func (self Tuple9[A, B, C, D, E, F, G, H, I]) Seventh() G { return self.seventh }
func (self Tuple9[A, B, C, D, E, F, G, H, I]) Eighth() H  { return self.eighth }
func (self Tuple9[A, B, C, D, E, F, G, H, I]) Ninth() I   { return self.ninth }

func MakeTuple9[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Tuple9[A, B, C, D, E, F, G, H, I] {
	return Tuple9[A, B, C, D, E, F, G, H, I]{a, b, c, d, e, f, g, h, i}
}
