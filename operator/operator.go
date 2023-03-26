package operator

import (
	"github.com/go-board/std/cond"
	"github.com/go-board/std/core"
)

func Neg[T core.Number](v T) T { return -v }

func Eq[T comparable](lhs T, rhs T) bool   { return lhs == rhs }
func Ne[T comparable](lhs T, rhs T) bool   { return lhs != rhs }
func Lt[T core.Ordered](lhs T, rhs T) bool { return lhs < rhs }
func Le[T core.Ordered](lhs T, rhs T) bool { return lhs <= rhs }
func Gt[T core.Ordered](lhs T, rhs T) bool { return lhs > rhs }
func Ge[T core.Ordered](lhs T, rhs T) bool { return lhs >= rhs }

func And[T ~bool](lhs T, rhs T) T { return lhs && rhs }
func Or[T ~bool](lhs T, rhs T) T  { return lhs || rhs }
func Not[T ~bool](v T) T          { return !v }

func AndAssign[T ~bool](lhs *T, rhs T) { *lhs = *lhs && rhs }
func OrAssign[T ~bool](lhs *T, rhs T)  { *lhs = *lhs || rhs }
func NotAssign[T ~bool](v *T)          { *v = !*v }

func Add[T core.Number | ~string](lhs T, rhs T) T { return lhs + rhs }
func Sub[T core.Number](lhs T, rhs T) T           { return lhs - rhs }
func Mul[T core.Number](lhs T, rhs T) T           { return lhs * rhs }
func Div[T core.Number](lhs T, rhs T) T           { return lhs / rhs }
func Rem[T core.Integer](lhs T, rhs T) T          { return lhs % rhs }

func AddAssign[T core.Number | ~string](lhs *T, rhs T) { *lhs += rhs }
func SubAssign[T core.Number](lhs *T, rhs T)           { *lhs -= rhs }
func MulAssign[T core.Number](lhs *T, rhs T)           { *lhs *= rhs }
func DivAssign[T core.Number](lhs *T, rhs T)           { *lhs /= rhs }
func RemAssign[T core.Integer](lhs *T, rhs T)          { *lhs %= rhs }

func BitAnd[T core.Integer](lhs T, rhs T) T { return lhs & rhs }
func BitOr[T core.Integer](lhs T, rhs T) T  { return lhs | rhs }
func BitXor[T core.Integer](lhs T, rhs T) T { return lhs ^ rhs }

func BitAndAssign[T core.Integer](lhs *T, rhs T) { *lhs &= rhs }
func BitOrAssign[T core.Integer](lhs *T, rhs T)  { *lhs |= rhs }
func BitXorAssign[T core.Integer](lhs *T, rhs T) { *lhs ^= rhs }

func Shl[T core.Integer, S core.Integer](v T, bit S) T { return v << bit }
func Shr[T core.Integer, S core.Integer](v T, bit S) T { return v >> bit }

func ShlAssign[T core.Integer, S core.Integer](v *T, bit S) { *v = *v << bit }
func ShrAssign[T core.Integer, S core.Integer](v *T, bit S) { *v = *v >> bit }

func Exchange[Lhs, Rhs any](lhs Lhs, rhs Rhs) (Rhs, Lhs) { return rhs, lhs }

func Ternary[T any](ok bool, lhs T, rhs T) T {
	return cond.Ternary(ok, lhs, rhs)
}

func RangeInteger[T core.Integer](start T, end T) []T {
	if end < start {
		return nil
	}
	rng := make([]T, 0, end-start)
	for i := start; i < end; i++ {
		rng = append(rng, i)
	}
	return rng
}

func Max[T core.Ordered](elem T, rest ...T) T {
	max := elem
	for _, e := range rest {
		if e > max {
			max = e
		}
	}
	return max
}

func Min[T core.Ordered](elem T, rest ...T) T {
	min := elem
	for _, e := range rest {
		if e < min {
			min = e
		}
	}
	return min
}
