//go:build go1.21

package slices

import (
	"cmp"
	"slices"

	"github.com/go-board/std/optional"
)

func Equal[T comparable, S1 ~[]T, S2 ~[]T](lhs S1, rhs S2) bool {
	return slices.Equal(lhs, S1(rhs))
}

func EqualBy[T, U any, S1 ~[]T, S2 ~[]U](lhs S1, rhs S2, eq func(T, U) bool) bool {
	return slices.EqualFunc(lhs, rhs, eq)
}

func Compare[T cmp.Ordered, S1 ~[]T, S2 ~[]T](lhs S1, rhs S2) int {
	return slices.Compare(lhs, S1(rhs))
}

func CompareBy[T, U any, S1 ~[]T, S2 ~[]U](lhs S1, rhs S2, cmp func(T, U) int) int {
	return slices.CompareFunc(lhs, rhs, cmp)
}

func Sort[T cmp.Ordered](slice []T) {
	slices.Sort(slice)
}

func SortBy[T any](slice []T, cmp func(T, T) int) {
	slices.SortFunc(slice, cmp)
}

func IsSorted[T cmp.Ordered](slice []T) bool {
	return slices.IsSorted(slice)
}

func IsSortedBy[T any](slice []T, cmp func(T, T) int) bool {
	return slices.IsSortedFunc(slice, cmp)
}

func Index[T comparable](slice []T, v T) optional.Optional[int] {
	return IndexBy(slice, func(t T) bool { return t == v })
}

// IndexBy returns the index of the first element in the given slice that satisfies the given predicate.
func IndexBy[T any](slice []T, predicate func(T) bool) optional.Optional[int] {
	i := slices.IndexFunc(slice, predicate)
	if i < 0 {
		return optional.None[int]()
	}
	return optional.Some(i)
}

// Contains returns true if the given slice contains the given element.
func Contains[T comparable](slice []T, v T) bool {
	return ContainsBy(slice, func(t T) bool { return t == v })
}

// ContainsBy returns true if the given slice contains an element that satisfies the given predicate.
func ContainsBy[T any](slice []T, predicate func(T) bool) bool {
	return slices.ContainsFunc(slice, predicate)
}
