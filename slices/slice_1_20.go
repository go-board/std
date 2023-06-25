//go:build !go1.21

package slices

import (
	"sort"

	"github.com/go-board/std/cmp"
	"github.com/go-board/std/core"
	"github.com/go-board/std/operator"
	"github.com/go-board/std/optional"
)

// Equal returns true if the given slices are equal.
func Equal[T comparable](lhs []T, rhs []T) bool {
	return EqualBy(lhs, rhs, operator.Eq[T])
}

// EqualBy returns true if the given slices are equal by the given function.
func EqualBy[T any](lhs []T, rhs []T, eq func(T, T) bool) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i, v := range lhs {
		if !eq(v, rhs[i]) {
			return false
		}
	}
	return true
}

type sortBy[T any] struct {
	cmp   func(lhs, rhs T) int
	inner []T
}

func (s sortBy[T]) Len() int           { return len(s.inner) }
func (s sortBy[T]) Less(i, j int) bool { return s.cmp(s.inner[i], s.inner[j]) < 0 }
func (s sortBy[T]) Swap(i, j int)      { s.inner[i], s.inner[j] = s.inner[j], s.inner[i] }

// SortBy sorts the given slice in-place by the given less function.
func SortBy[T any](slice []T, cmp func(lhs, rhs T) int) []T {
	sort.Sort(sortBy[T]{cmp: cmp, inner: slice})
	return slice
}

// Sort sorts the given slice in-place.
func Sort[T core.Ordered](slice []T) []T {
	return SortBy(slice, cmp.Compare[T])
}

// IsSorted returns true if the given slice is sorted.
func IsSorted[T core.Ordered](slice []T) bool {
	return IsSortedBy(slice, cmp.Compare[T])
}

// IsSortedBy returns true if the given slice is sorted by the given less function.
func IsSortedBy[T any](slice []T, cmp func(lhs, rhs T) int) bool {
	return sort.IsSorted(sortBy[T]{cmp: cmp, inner: slice})
}

// Index returns the index of the first element in the given slice that same with the given element.
func Index[T comparable](slice []T, v T) optional.Optional[int] {
	return IndexBy(slice, func(t T) bool { return t == v })
}

// IndexBy returns the index of the first element in the given slice that satisfies the given predicate.
func IndexBy[T any](slice []T, predicate func(T) bool) optional.Optional[int] {
	for i, v := range slice {
		if predicate(v) {
			return optional.Some(i)
		}
	}
	return optional.None[int]()
}

// Contains returns true if the given slice contains the given element.
func Contains[T comparable](slice []T, v T) bool {
	return Index(slice, v).IsSome()
}

// ContainsBy returns true if the given slice contains an element that satisfies the given predicate.
func ContainsBy[T any](slice []T, predicate func(T) bool) bool {
	return IndexBy(slice, predicate).IsSome()
}
