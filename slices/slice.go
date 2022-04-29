package slices

import (
	"sort"

	"github.com/go-board/std/optional"
)

type sortBy[T any] struct {
	less  func(a, b T) bool
	inner []T
}

func (s sortBy[T]) Len() int           { return len(s.inner) }
func (s sortBy[T]) Less(i, j int) bool { return s.less(s.inner[i], s.inner[j]) }
func (s sortBy[T]) Swap(i, j int)      { s.inner[i], s.inner[j] = s.inner[j], s.inner[i] }

// SortBy sort the given slice in-place by the given less function.
func SortBy[T any](slice []T, less func(a, b T) bool) {
	sort.Sort(sortBy[T]{less: less, inner: slice})
}

// Map returns a new slice with the results of applying the given function to each element in the given slice.
func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// ForEach iterates over the given slice and calls the given function for each element.
func ForEach[T any](slice []T, f func(T)) {
	for _, v := range slice {
		f(v)
	}
}

// Filter returns a new slice with all elements that satisfy the given predicate.
func Filter[T any](slice []T, f func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Fold returns the result of applying the given function to each element in the given slice.
func Fold[T, A any](slice []T, initial A, accumulator func(A, T) A) A {
	result := initial
	for _, v := range slice {
		result = accumulator(result, v)
	}
	return result
}

// Reduce returns the result of applying the given function to each element in the given slice.
func Reduce[T any](slice []T, f func(T, T) T) optional.Optional[T] {
	if len(slice) == 0 {
		return optional.None[T]()
	}
	return optional.Some(Fold(slice[0:], slice[0], f))
}

// Any returns true if any element in the given slice satisfies the given predicate.
func Any[T any](slice []T, f func(T) bool) bool {
	for _, v := range slice {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the given slice satisfy the given predicate.
func All[T any](slice []T, f func(T) bool) bool {
	for _, v := range slice {
		if !f(v) {
			return false
		}
	}
	return true
}

// None returns true if no element in the given slice satisfies the given predicate.
func None[T any](slice []T, f func(T) bool) bool {
	return !Any(slice, f)
}

// FindIndexBy returns the index of the first element in the given slice that satisfies the given predicate.
func FindIndexBy[T any](slice []T, v T, eq func(T, T) bool) int {
	for i, vv := range slice {
		if eq(v, vv) {
			return i
		}
	}
	return -1
}

// ContainsBy returns true if the given slice contains an element that satisfies the given predicate.
func ContainsBy[T any](slice []T, v T, cmp func(T, T) bool) bool {
	return Any(slice, func(t T) bool { return cmp(t, v) })
}

// Contains returns true if the given slice contains the given element.
func Contains[T comparable](slice []T, v T) bool {
	return ContainsBy(slice, v, func(t1, t2 T) bool { return t1 == t2 })
}

// MaxBy returns the maximium element in the given slice that satisfies the given function.
func MaxBy[T any](slice []T, less func(T, T) bool) optional.Optional[T] {
	return Reduce(slice, func(a, b T) T {
		if less(a, b) {
			return b
		} else {
			return a
		}
	})
}

// MinBy returns the minimium element in the given slice that satisfies the given function.
func MinBy[T any](slice []T, less func(T, T) bool) optional.Optional[T] {
	return Reduce(slice, func(a, b T) T {
		if less(a, b) {
			return a
		} else {
			return b
		}
	})
}

// Nth returns the nth element in the given slice.
func Nth[T any](slice []T, n int) optional.Optional[T] {
	if n < 0 {
		n = len(slice) + n
	}
	if n < 0 || n >= len(slice) {
		return optional.None[T]()
	}
	return optional.Some(slice[n])
}

// Flatten returns a new slice with all elements in the given slice and all elements in all sub-slices.
func Flatten[T any](slice [][]T) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		result = append(result, v...)
	}
	return result
}

// Chunk returns a new slice with the given slice split into smaller slices of the given size.
func Chunk[T any](slice []T, chunk int) [][]T {
	result := make([][]T, 0, len(slice)/chunk+1)
	for i := 0; i < len(slice); i += chunk {
		result = append(result, slice[i:i+chunk])
	}
	return result
}

// GroupBy returns a new map with the given slice split into smaller slices of the given size.
func GroupBy[T any, TKey comparable](slice []T, group func(T) TKey) map[TKey][]T {
	result := make(map[TKey][]T)
	for _, v := range slice {
		key := group(v)
		result[key] = append(result[key], v)
	}
	return result
}

// EqualBy returns true if the given slices are equal by the given function.
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

// Equal returns true if the given slices are equal.
func Equal[T comparable](slice1 []T, slice2 []T) bool {
	return EqualBy(slice1, slice2, func(a, b T) bool { return a == b })
}

// DeepClone returns a new slice with the same elements as the given slice.
func DeepClone[T any](slice []T, clone func(T) T) []T {
	return Map(slice, clone)
}
