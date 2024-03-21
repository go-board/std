//go:build goexperiment.rangefunc

package iter

import (
	"iter"

	"github.com/go-board/std/cmp"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/tuple"
)

// Cmp compares two [Seq] and return first non-equal result.
//
// Example:
//
//	iter.Cmp(seq(1,2,3), seq(1,2,3)) => 0
//	iter.Cmp(seq(1,2,3), seq(1,2,4)) => -1
//	iter.Cmp(seq(1,2,3), seq(1,2,2)) => 1
//	iter.Cmp(seq(1,2), seq(1))       => 1
//	iter.Cmp(seq(1), seq(1,2))       => -1
func Cmp[E cmp.Ordered](x Seq[E], y Seq[E]) int {
	return CmpFunc(x, y, cmp.Compare[E])
}

// CmpFunc compares two [Seq] using the give compare function.
//
// And returning at first non-equal result.
func CmpFunc[A, B any](x Seq[A], y Seq[B], f func(A, B) int) int {
	itx, sx := iter.Pull(iter.Seq[A](x))
	ity, sy := iter.Pull(iter.Seq[B](y))
	defer sx()
	defer sy()
	for {
		xe, oke := itx()
		xf, okf := ity()
		if oke && okf {
			if c := f(xe, xf); c != 0 {
				return c
			}
		}
		if !oke && !okf {
			return 0
		}
		if oke {
			return +1
		} else {
			return -1
		}
	}
}

// Gt test if left seq great-than the right seq.
func Gt[E cmp.Ordered](x Seq[E], y Seq[E]) bool { return Cmp(x, y) > 0 }

// Ge test if left seq great-equal the right seq.
func Ge[E cmp.Ordered](x Seq[E], y Seq[E]) bool { return Cmp(x, y) >= 0 }

// Lt test if left seq less-than the right seq.
func Lt[E cmp.Ordered](x Seq[E], y Seq[E]) bool { return Cmp(x, y) < 0 }

// Le test if left seq less-equal the right seq.
func Le[E cmp.Ordered](x Seq[E], y Seq[E]) bool { return Cmp(x, y) <= 0 }

// Eq test whether two [Seq] are equality.
func Eq[E comparable](x Seq[E], y Seq[E]) bool {
	return EqFunc(x, y, func(l, r E) bool { return l == r })
}

// EqFunc test whether two [Seq] are equality using the give compare function.
func EqFunc[A, B any](x Seq[A], y Seq[B], f func(A, B) bool) bool {
	itx, sx := iter.Pull(iter.Seq[A](x))
	ity, sy := iter.Pull(iter.Seq[B](y))
	defer sx()
	defer sy()
	for {
		xe, oke := itx()
		xf, okf := ity()
		if oke && okf && f(xe, xf) {
			continue
		}
		if !oke && !okf {
			return true
		}
		return false
	}
}

// Ne test whether two [Seq] are not equality.
func Ne[E comparable](x Seq[E], y Seq[E]) bool {
	return NeFunc(x, y, func(l E, r E) bool { return l == r })
}

// NeFunc test whether two [Seq] are not equality using the give equality function.
func NeFunc[A, B any](x Seq[A], y Seq[B], f func(A, B) bool) bool {
	itx, sx := iter.Pull(iter.Seq[A](x))
	ity, sy := iter.Pull(iter.Seq[B](y))
	defer sx()
	defer sy()
	for {
		xe, oke := itx()
		xf, okf := ity()
		if oke && okf && f(xe, xf) {
			continue
		}
		if !oke && !okf {
			return false
		}
		return true
	}
}

// Zip zips two [Seq] into a single [Seq] of pairs.
//
// If either [Seq] finished, the zipped [Seq] also finish.
//
// Example:
//
//	iter.Zip(seq(1,3,5), seq(2,4,6)) => seq: pair(1,2), pair(3,4), pair(5,6)
func Zip[A, B any](x Seq[A], y Seq[B]) Seq[tuple.Pair[A, B]] {
	return ZipFunc(x, y, tuple.MakePair[A, B])
}

// ZipFunc zips two [Seq] into a single [Seq] of a new type zipped by the given function.
//
// If either [Seq] finished, the zipped [Seq] also finish.
//
// Example:
//
//	iter.ZipFunc(seq(1,3,5), seq(2,4,6), func(x, y int) int { return x+y }) => seq: 3,7,11
func ZipFunc[A, B, C any](x Seq[A], y Seq[B], f func(A, B) C) Seq[C] {
	itx, sx := iter.Pull(iter.Seq[A](x))
	ity, sy := iter.Pull(iter.Seq[B](y))
	return func(yield func(C) bool) {
		defer sx()
		defer sy()
		for {
			xa, oka := itx()
			xb, okb := ity()
			if !(oka && okb && yield(f(xa, xb))) {
				break
			}
		}
	}
}

// PullOption turn a [Seq] into a pull style iterator and a stop function,
// iterator function returns an optional value on each call.
func PullOption[E any](s Seq[E]) (func() optional.Optional[E], func()) {
	it, stop := iter.Pull(iter.Seq[E](s))
	return func() optional.Optional[E] { return optional.FromPair(it()) }, stop
}
