//go:build !go1.21

package slices

import (
	"sort"

	"github.com/go-board/std/cmp"
	constraints "github.com/go-board/std/constraints"
	"github.com/go-board/std/iter"
	"github.com/go-board/std/operator"
	"github.com/go-board/std/optional"
)

// Equal returns true if the given slices are equal.
func Equal[T comparable, S1 ~[]T, S2 ~[]T](lhs S1, rhs S2) bool {
	return EqualBy(lhs, rhs, operator.Eq[T])
}

// EqualBy returns true if the given slices are equal by the given function.
func EqualBy[T, E any, S1 ~[]T, S2 ~[]E](lhs S1, rhs S2, eq func(T, E) bool) bool {
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
func SortBy[T any, S ~[]T](slice S, cmp func(lhs, rhs T) int) S {
	sort.Sort(sortBy[T]{cmp: cmp, inner: slice})
	return slice
}

// Sort sorts the given slice in-place.
func Sort[T constraints.Ordered, S ~[]T](slice S) S {
	return SortBy(slice, cmp.Compare[T])
}

// IsSorted returns true if the given slice is sorted.
func IsSorted[T constraints.Ordered, S ~[]T](slice S) bool {
	return iter.IsSorted(Forward(slice), cmp.Compare[T])
}

// IsSortedBy returns true if the given slice is sorted by the given less function.
func IsSortedBy[T any, S ~[]T](slice S, cmp func(lhs, rhs T) int) bool {
	return iter.IsSortedFunc(Forward(slice), cmp)
}

// Index returns the index of the first element in the given slice that same with the given element.
func Index[T comparable, S ~[]T](slice S, v T) optional.Optional[int] {
	return IndexBy(slice, func(t T) bool { return t == v })
}

// IndexBy returns the index of the first element in the given slice that satisfies the given predicate.
func IndexBy[T any, S ~[]T](slice S, predicate func(T) bool) optional.Optional[int] {
	x := iter.IndexFirstFunc(Forward(slice), predicate)
	if x < 0 {
		return optional.None[int]()
	}
	return optional.Some(x)
}

// Contains returns true if the given slice contains the given element.
func Contains[T comparable, S ~[]T](slice S, v T) bool {
	return iter.Contains(Forward(slice), v)
}

// ContainsBy returns true if the given slice contains an element that satisfies the given predicate.
func ContainsBy[T any, S ~[]T](slice S, predicate func(T) bool) bool {
	return iter.ContainsFunc(Forward(slice), predicate)
}
