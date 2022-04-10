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

func Reduce[T any](slice []T, f func(T, T) T) T {
	result := slice[0]
	for i := 1; i < len(slice); i++ {
		result = f(result, slice[i])
	}
	return result
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

func Find[T any](slice []T, f func(T) bool) (int, bool) {
	for i, v := range slice {
		if f(v) {
			return i, true
		}
	}
	return 0, false
}

func FindIndex[T any](slice []T, f func(T) bool) int {
	for i, v := range slice {
		if f(v) {
			return i
		}
	}
	return -1
}

func Contains[T any](slice []T, v T, cmp func(T, T) bool) bool {
	return FindIndex(slice, func(x T) bool {
		return cmp(x, v)
	}) != -1
}

func ContainsAny[T any](slice []T, v []T, cmp func(T, T) bool) bool {
	for _, x := range v {
		if Contains(slice, x, cmp) {
			return true
		}
	}
	return false
}

func Max[T any](slice []T, cmp func(T, T) bool) optional.Optional[T] {
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

func Min[T any](slice []T, cmp func(T, T) bool) optional.Optional[T] {
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
	return optional.Some[T](slice[n])
}

func Flatten[T any](slice [][]T) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		result = append(result, v...)
	}
	return result
}

func Equal[T any](slice1 []T, slice2 []T, eq func(T, T) bool) bool {
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
