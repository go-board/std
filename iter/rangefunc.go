//go:build goexperiment.rangefunc

package iter

import (
	"iter"

	"github.com/go-board/std/cmp"
)

// Cmp compares two [Seq] and return first non-equal result.
//
// Example:
//
//	iter.Cmp(seq(1,2,3), seq(1,2,3)) => 0
//	iter.Cmp(seq(1,2,3), seq(1,2,4)) => -1
//	iter.Cmp(seq(1,2,3), seq(1,2,2)) => 1
//	iter.Cmp(seq(1,2), seq(1)) => 1
//	iter.Cmp(seq(1), seq(1,2)) => -1
func Cmp[E cmp.Ordered](x Seq[E], y Seq[E]) int {
	return CmpFunc(x, y, cmp.Compare[E])
}

// CmpFunc compares two [Seq] using the give compare function.
//
// And returning at first non-equal result.
func CmpFunc[E, F any](x Seq[E], y Seq[F], f func(E, F) int) int {
	itx, sx := iter.Pull(iter.Seq[E](x))
	ity, sy := iter.Pull(iter.Seq[F](y))
	defer sx()
	defer sy()
	for {
		xe, oke := itx()
		xf, okf := ity()
		if oke == okf {
			if c := f(xe, xf); c != 0 {
				return c
			}
			continue
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
func EqFunc[E, F any](x Seq[E], y Seq[F], f func(E, F) bool) bool {
	itx, sx := iter.Pull(iter.Seq[E](x))
	ity, sy := iter.Pull(iter.Seq[F](y))
	defer sx()
	defer sy()
	for {
		xe, oke := itx()
		xf, okf := ity()
		return oke == okf && f(xe, xf)
	}
}

// Ne test whether two [Seq] are not equality.
func Ne[E comparable](x Seq[E], y Seq[E]) bool { return !Eq(x, y) }

// NeFunc test whether two [Seq] are not equality using the give compare function.
func NeFunc[E, F any](x Seq[E], y Seq[F], f func(E, F) bool) bool { return !EqFunc(x, y, f) }
