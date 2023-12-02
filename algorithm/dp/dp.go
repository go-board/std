package dp

func LcsBy[T any, S ~[]T](left S, right S, eq func(T, T) bool) S {
	if len(left) == 0 || len(right) == 0 {
		return nil
	}
	if eq(left[0], right[0]) {
		return append([]T{left[0]}, LcsBy(left[1:], right[1:], eq)...)
	}
	lcs1, lcs2 := LcsBy(left[1:], right, eq), LcsBy(left, right[1:], eq)
	if len(lcs1) > len(lcs2) {
		return lcs1
	}
	return lcs2
}

func Lcs[T comparable, S ~[]T](left S, right S) S {
	return LcsBy(left, right, func(a, b T) bool { return a == b })
}
