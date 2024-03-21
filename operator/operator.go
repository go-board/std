package operator

import (
	"github.com/go-board/std/cond"
	"github.com/go-board/std/constraints"
)

// Identify return self.
func Identify[E any](x E) E { return x }

// Neg return negative value of input value.
func Neg[T constraints.Numeric](v T) T { return -v }

// Eq test two value whether equal or not.
//
// if equal return true, otherwise return false.
func Eq[T comparable](lhs T, rhs T) bool { return lhs == rhs }

// Ne test two values whether equal or not.
//
// if equal return false, otherwise return true.
func Ne[T comparable](lhs T, rhs T) bool { return lhs != rhs }

// Lt test whether lhs less than rhs value.
//
// if lhs less than rhs, return true, otherwise return false.
func Lt[T constraints.Ordered](lhs T, rhs T) bool { return lhs < rhs }

// Le test whether lhs less than or equal to rhs value.
//
// if lhs less than or equal to rhs, return true, otherwise return false.
func Le[T constraints.Ordered](lhs T, rhs T) bool { return lhs <= rhs }

// Gt test whether lhs greater than rhs value.
//
// if lhs greater than rhs, return true, otherwise return false.
func Gt[T constraints.Ordered](lhs T, rhs T) bool { return lhs > rhs }

// Ge test whether lhs greater than or equal to rhs value.
//
// if lhs greater than or equal to rhs, return true, otherwise return false.
func Ge[T constraints.Ordered](lhs T, rhs T) bool { return lhs >= rhs }

// And act same as lhs && rhs.
func And[T ~bool](lhs T, rhs T) T { return lhs && rhs }

// Or act same as lhs || rhs.
func Or[T ~bool](lhs T, rhs T) T { return lhs || rhs }

// Not act same as !v.
func Not[T ~bool](v T) T { return !v }

// AndAssign act same as lhs &&= rhs in place.
func AndAssign[T ~bool](lhs *T, rhs T) { *lhs = *lhs && rhs }

// OrAssign act same as lhs ||= rhs in place.
func OrAssign[T ~bool](lhs *T, rhs T) { *lhs = *lhs || rhs }

// NotAssign invert bool value in place.
func NotAssign[T ~bool](v *T) { *v = !*v }

// Add act same as lhs + rhs.
func Add[T constraints.Numeric | ~string](lhs T, rhs T) T { return lhs + rhs }

// Sub act same as lhs - rhs.
func Sub[T constraints.Numeric](lhs T, rhs T) T { return lhs - rhs }

// Mul act same as lhs * rhs.
func Mul[T constraints.Numeric](lhs T, rhs T) T { return lhs * rhs }

// Div act same as lhs / rhs.
func Div[T constraints.Numeric](lhs T, rhs T) T { return lhs / rhs }

// Rem act same as lhs % rhs.
func Rem[T constraints.Integer](lhs T, rhs T) T { return lhs % rhs }

// AddAssign act same as lhs += rhs in place.
func AddAssign[T constraints.Numeric | ~string](lhs *T, rhs T) { *lhs += rhs }

// SubAssign act same as lhs -= rhs in place.
func SubAssign[T constraints.Numeric](lhs *T, rhs T) { *lhs -= rhs }

// MulAssign act same as lhs *= rhs in place.
func MulAssign[T constraints.Numeric](lhs *T, rhs T) { *lhs *= rhs }

// DivAssign act same as lhs /= rhs in place.
func DivAssign[T constraints.Numeric](lhs *T, rhs T) { *lhs /= rhs }

// RemAssign act same as lhs %= rhs in place.
func RemAssign[T constraints.Integer](lhs *T, rhs T) { *lhs %= rhs }

// BitAnd act same as lhs & rhs.
func BitAnd[T constraints.Integer](lhs T, rhs T) T { return lhs & rhs }

// BitOr act same as lhs | rhs.
func BitOr[T constraints.Integer](lhs T, rhs T) T { return lhs | rhs }

// BitXor act same as lhs ^ rhs.
func BitXor[T constraints.Integer](lhs T, rhs T) T { return lhs ^ rhs }

// BitAndAssign act same as lhs &= rhs in place.
func BitAndAssign[T constraints.Integer](lhs *T, rhs T) { *lhs &= rhs }

// BitOrAssign act same as lhs |= rhs in place.
func BitOrAssign[T constraints.Integer](lhs *T, rhs T) { *lhs |= rhs }

// BitXorAssign act same as lhs ^= rhs in place.
func BitXorAssign[T constraints.Integer](lhs *T, rhs T) { *lhs ^= rhs }

// Shl act same as v << bit.
func Shl[T constraints.Integer, S constraints.Integer](v T, bit S) T { return v << bit }

// Shr act same as v >> bit.
func Shr[T constraints.Integer, S constraints.Integer](v T, bit S) T { return v >> bit }

// ShlAssign act same as v <<= bit.
func ShlAssign[T constraints.Integer, S constraints.Integer](v *T, bit S) { *v = *v << bit }

// ShrAssign act same as v >>= bit.
func ShrAssign[T constraints.Integer, S constraints.Integer](v *T, bit S) { *v = *v >> bit }

// Exchange change order of the two value.
func Exchange[Lhs, Rhs any](lhs Lhs, rhs Rhs) (Rhs, Lhs) { return rhs, lhs }

// Ternary act same as ok ? lhs : rhs.
func Ternary[T any](ok bool, lhs T, rhs T) T {
	return cond.Ternary(ok, lhs, rhs)
}

func LazyTernary[T any](ok bool, lhs func() T, rhs func() T) T {
	if ok {
		return lhs()
	}
	return rhs()
}

// RangeInteger returns a slice that contains all integers in [start, end).
func RangeInteger[T constraints.Integer](start T, end T) []T {
	if end < start {
		return nil
	}
	rng := make([]T, 0, end-start)
	for i := start; i < end; i++ {
		rng = append(rng, i)
	}
	return rng
}

// Max return maximum of the given values.
func Max[T constraints.Ordered](elem T, rest ...T) T {
	max := elem
	for _, e := range rest {
		if e > max {
			max = e
		}
	}
	return max
}

// Min return minimum of the given values.
func Min[T constraints.Ordered](elem T, rest ...T) T {
	min := elem
	for _, e := range rest {
		if e < min {
			min = e
		}
	}
	return min
}
