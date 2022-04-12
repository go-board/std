package slices

import (
	"sort"

	"github.com/go-board/std/optional"
)

func Sort[T any](slice []T, less func(a, b T) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
}

func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func ForEach[T any](slice []T, f func(T)) {
	for _, v := range slice {
		f(v)
	}
}

func Filter[T any](slice []T, f func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Fold[T, A any](slice []T, initial A, accumulator func(A, T) A) A {
	result := initial
	for _, v := range slice {
		result = accumulator(result, v)
	}
	return result
}

func Reduce[T any](slice []T, f func(T, T) T) optional.Optional[T] {
	if len(slice) == 0 {
		return optional.None[T]()
	}
	return optional.Some(Fold(slice[0:], slice[0], f))
}

func Any[T any](slice []T, f func(T) bool) bool {
	for _, v := range slice {
		if f(v) {
			return true
		}
	}
	return false
}
func All[T any](slice []T, f func(T) bool) bool {
	for _, v := range slice {
		if !f(v) {
			return false
		}
	}
	return true
}

func None[T any](slice []T, f func(T) bool) bool {
	return !Any(slice, f)
}

func FindIndexBy[T any](slice []T, v T, eq func(T, T) bool) int {
	for i, vv := range slice {
		if eq(v, vv) {
			return i
		}
	}
	return -1
}

func ContainsBy[T any](slice []T, v T, cmp func(T, T) bool) bool {
	return Any(slice, func(t T) bool { return cmp(t, v) })
}

func Contains[T comparable](slice []T, v T) bool {
	return ContainsBy(slice, v, func(t1, t2 T) bool { return t1 == t2 })
}

func MaxBy[T any](slice []T, cmp func(T, T) bool) optional.Optional[T] {
	if len(slice) == 0 {
		return optional.None[T]()
	}
	max := slice[0]
	for _, v := range slice {
		if cmp(v, max) {
			max = v
		}
	}
	return optional.Some(max)
}

func MinBy[T any](slice []T, cmp func(T, T) bool) optional.Optional[T] {
	if len(slice) == 0 {
		return optional.None[T]()
	}
	min := slice[0]
	for _, v := range slice {
		if cmp(min, v) {
			min = v
		}
	}
	return optional.Some(min)
}

func Nth[T any](slice []T, n int) optional.Optional[T] {
	if n < 0 {
		n = len(slice) + n
	}
	if n < 0 || n >= len(slice) {
		return optional.None[T]()
	}
	return optional.Some(slice[n])
}

func Flatten[T any](slice [][]T) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		result = append(result, v...)
	}
	return result
}

func EqualBy[T any](slice1 []T, slice2 []T, eq func(T, T) bool) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if !eq(v, slice2[i]) {
			return false
		}
	}
	return true
}

func Equal[T comparable](slice1 []T, slice2 []T) bool {
	return EqualBy(slice1, slice2, func(a, b T) bool { return a == b })
}
