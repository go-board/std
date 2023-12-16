//go:build goexperiment.rangefunc

package iter

import (
	"iter"

	"github.com/go-board/std/cmp"
	"github.com/go-board/std/operator"
)

func Cmp[E cmp.Ordered](x Seq[E], y Seq[E]) int {
	return CmpFunc(x, y, cmp.Compare[E])
}

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

func Gt[E cmp.Ordered](x Seq[E], y Seq[E]) bool { return Cmp(x, y) > 0 }
func Ge[E cmp.Ordered](x Seq[E], y Seq[E]) bool { return Cmp(x, y) >= 0 }
func Lt[E cmp.Ordered](x Seq[E], y Seq[E]) bool { return Cmp(x, y) < 0 }
func Le[E cmp.Ordered](x Seq[E], y Seq[E]) bool { return Cmp(x, y) <= 0 }

func Eq[E comparable](x Seq[E], y Seq[E]) bool {
	return EqFunc(x, y, operator.Eq[E])
}

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

func Ne[E comparable](x Seq[E], y Seq[E]) bool { return !Eq(x, y) }

func NeFunc[E, F any](x Seq[E], y Seq[F], f func(E, F) bool) bool { return !EqFunc(x, y, f) }
