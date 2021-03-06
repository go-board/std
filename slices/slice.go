package slices

import (
	"errors"
	"math/rand"
	"sort"

	"github.com/go-board/std/clone"
	"github.com/go-board/std/cond"
	"github.com/go-board/std/core"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/result"
	"github.com/go-board/std/sets"
)

// All returns true if all elements in the given slice satisfy the given predicate.
func All[T any](slice []T, f func(T) bool) bool {
	for _, v := range slice {
		if !f(v) {
			return false
		}
	}
	return true
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

// Chunk returns a new slice with the given slice split into smaller slices of the given size.
func Chunk[T any](slice []T, chunk int) [][]T {
	size := len(slice)
	result := make([][]T, 0, len(slice)/chunk+1)
	for i := 0; i < size; i += chunk {
		if i+chunk > size {
			result = append(result, slice[i:])
		} else {
			result = append(result, slice[i:i+chunk])
		}
	}
	return result
}

// Clone returns a new slice with the same elements as the given slice.
func Clone[T any](slice []T) []T {
	return DeepCloneBy(slice, func(t T) T { return t })
}

// Contains returns true if the given slice contains the given element.
func Contains[T comparable](slice []T, v T) bool {
	return ContainsBy(slice, v, func(lhs, rhs T) bool { return lhs == rhs })
}

// ContainsBy returns true if the given slice contains an element that satisfies the given predicate.
func ContainsBy[T any](slice []T, v T, cmp func(T, T) bool) bool {
	return Any(slice, func(t T) bool { return cmp(t, v) })
}

// DeepClone returns a new slice with the cloned elements.
func DeepClone[T clone.Cloneable[T]](slice []T) []T {
	return DeepCloneBy(slice, func(t T) T { return t.Clone() })
}

// DeepCloneBy returns a new slice with the cloned elements as the given slice.
func DeepCloneBy[T any](slice []T, clone func(T) T) []T {
	return Map(slice, clone)
}

// DifferenceBy returns a new slice with the elements that are in the first slice but not in the second by the given function.
func DifferenceBy[T any](lhs []T, rhs []T, eq func(T, T) bool) []T {
	result := make([]T, 0, len(lhs))
	for _, v := range lhs {
		// TODO: optimize use O(1) lookup
		if !ContainsBy(rhs, v, eq) {
			result = append(result, v)
		}
	}
	return result
}

// Distinct returns a new slice with the given slice without duplicates.
func Distinct[T comparable](slice []T) []T {
	return sets.NewHashSet(slice...).ToSlice()
}

// DistinctBy returns a new slice with the distinct elements of the given slice by the given function.
func DistinctBy[T any, K comparable](slice []T, key func(T) K) []T {
	m := make(map[K]T, len(slice))
	for _, v := range slice {
		m[key(v)] = v
	}
	result := make([]T, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// Equal returns true if the given slices are equal.
func Equal[T comparable](lhs []T, rhs []T) bool {
	return EqualBy(lhs, rhs, func(lhs, rhs T) bool { return lhs == rhs })
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

// FilterIndexed returns a new slice with all elements that satisfy the given predicate.
func FilterIndexed[T any](slice []T, f func(T, int) bool) []T {
	result := make([]T, 0, len(slice))
	for i, v := range slice {
		if f(v, i) {
			result = append(result, v)
		}
	}
	return result
}

// FindIndex returns the index of the first element in the given slice that satisfies the given predicate.
// Deprecated: use Index instead.
func FindIndex[T comparable](slice []T, v T) int {
	return Index(slice, v).ValueOr(-1)
}

// FindIndexBy returns the index of the first element in the given slice that satisfies the given predicate.
// Deprecated: use IndexBy instead.
func FindIndexBy[T any](slice []T, v T, eq func(T, T) bool) int {
	return IndexBy(slice, v, eq).ValueOr(-1)
}

// Flatten returns a new slice with all elements in the given slice and all elements in all sub-slices.
func Flatten[T any](slice [][]T) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		result = append(result, v...)
	}
	return result
}

// FlattenBy returns a new slice with all elements in the given slice and all elements in the given slices.
func FlattenBy[T, S any](slice []T, f func(T) []S) []S {
	result := make([]S, 0, len(slice))
	for _, v := range slice {
		result = append(result, f(v)...)
	}
	return result
}

// Fold accumulates value starting with initial value and applying accumulator from left to right to current accum value and each element.
// Returns the final accum value or initial value if the slice is empty.
func Fold[T, A any](slice []T, initial A, accumulator func(A, T) A) A {
	result := initial
	for _, v := range slice {
		result = accumulator(result, v)
	}
	return result
}

// FoldRight accumulates value starting with initial value and applying accumulator from right to left to current accum value and each element.
// Returns the final accum value or initial value if the slice is empty.
func FoldRight[T, A any](slice []T, initial A, accumulator func(A, T) A) A {
	accum := initial
	for i := len(slice) - 1; i >= 0; i-- {
		accum = accumulator(accum, slice[i])
	}
	return accum
}

// ForEach iterates over the given slice and calls the given function for each element.
func ForEach[T any](slice []T, f func(T)) {
	for _, v := range slice {
		f(v)
	}
}

// ForEachIndexed iterates over the given slice and calls the given function for each element.
func ForEachIndexed[T any](slice []T, f func(T, int)) {
	for i, v := range slice {
		f(v, i)
	}
}

// GroupBy returns a new map with the given slice split into smaller slices of the given size.
func GroupBy[T any, TKey comparable, TValue any](slice []T, group func(T) (TKey, TValue)) map[TKey][]TValue {
	result := make(map[TKey][]TValue)
	for _, v := range slice {
		key, value := group(v)
		result[key] = append(result[key], value)
	}
	return result
}

// Index returns the index of the first element in the given slice that same with the given element.
func Index[T comparable](slice []T, v T) optional.Optional[int] {
	return IndexBy(slice, v, func(lhs, rhs T) bool { return lhs == rhs })
}

// IndexBy returns the index of the first element in the given slice that satisfies the given predicate.
func IndexBy[T any](slice []T, v T, eq func(T, T) bool) optional.Optional[int] {
	for i, vv := range slice {
		if eq(v, vv) {
			return optional.Some(i)
		}
	}
	return optional.None[int]()
}

// IntersectionBy returns a new slice with the elements that are in both given slices by the given function.
func IntersectionBy[T any](lhs []T, rhs []T, eq func(T, T) bool) []T {
	result := make([]T, 0, len(lhs))
	for _, v := range lhs {
		// TODO: optimize use O(1) lookup
		if ContainsBy(rhs, v, eq) {
			result = append(result, v)
		}
	}
	return result
}

// IsSorted returns true if the given slice is sorted.
func IsSorted[T core.Ordered](slice []T) bool {
	return IsSortedBy(slice, func(lhs, rhs T) bool { return lhs < rhs })
}

// IsSortedBy returns true if the given slice is sorted by the given less function.
func IsSortedBy[T any](slice []T, less func(lhs, rhs T) bool) bool {
	return sort.IsSorted(sortBy[T]{less: less, inner: slice})
}

// LastIndex returns the index of the last element in the given slice that same with the given element.
func LastIndex[T comparable](slice []T, v T) optional.Optional[int] {
	return LastIndexBy(slice, v, func(lhs, rhs T) bool { return lhs == rhs })
}

// LastIndexBy returns the index of the last element in the given slice that satisfies the given predicate.
func LastIndexBy[T any](slice []T, v T, eq func(T, T) bool) optional.Optional[int] {
	for i := len(slice) - 1; i >= 0; i-- {
		if eq(v, slice[i]) {
			return optional.Some(i)
		}
	}
	return optional.None[int]()
}

// Map returns a new slice with the results of applying the given function to each element in the given slice.
func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// MapIndexed returns a new slice with the results of applying the given function to each element in the given slice.
func MapIndexed[T, U any](slice []T, f func(T, int) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v, i)
	}
	return result
}

// Max returns the maximum element in the given slice.
func Max[T core.Ordered](slice []T) optional.Optional[T] {
	return MaxBy(slice, func(lhs, rhs T) bool { return lhs < rhs })
}

// MaxBy returns the maximum element in the given slice that satisfies the given function.
func MaxBy[T any](slice []T, less func(T, T) bool) optional.Optional[T] {
	return Reduce(slice, func(lhs, rhs T) T {
		return cond.Ternary(less(lhs, rhs), rhs, lhs)
	})
}

// Min returns the minimum element in the given slice.
func Min[T core.Ordered](slice []T) optional.Optional[T] {
	return MinBy(slice, func(lhs, rhs T) bool { return lhs < rhs })
}

// MinBy returns the minimum element in the given slice that satisfies the given function.
func MinBy[T any](slice []T, less func(T, T) bool) optional.Optional[T] {
	return Reduce(slice, func(lhs, rhs T) T {
		return cond.Ternary(less(lhs, rhs), lhs, rhs)
	})
}

// None returns true if no element in the given slice satisfies the given predicate.
func None[T any](slice []T, f func(T) bool) bool {
	return !Any(slice, f)
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

// Reduce returns the result of applying the given function to each element in the given slice.
func Reduce[T any](slice []T, f func(T, T) T) optional.Optional[T] {
	if len(slice) == 0 {
		return optional.None[T]()
	}
	return optional.Some(Fold(slice[1:], slice[0], f))
}

// ReduceRight returns the result of applying the given function to each element in the given slice.
func ReduceRight[T any](slice []T, f func(T, T) T) optional.Optional[T] {
	if len(slice) == 0 {
		return optional.None[T]()
	}
	return optional.Some(FoldRight(slice[:len(slice)-1], slice[len(slice)-1], f))
}

// Reverse returns a new slice with the elements in the given slice in reverse order.
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shuffle returns a new slice with the given slice shuffled.
func Shuffle[T any](slice []T) {
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}

// Single returns the single element,
// or return an error if the collection is empty or has more than one element.
func Single[T any](slice []T) result.Result[T] {
	if len(slice) == 1 {
		return result.Ok(slice[0])
	}
	return result.Err[T](errors.New("slice is not scalar"))
}

type sortBy[T any] struct {
	less  func(lhs, rhs T) bool
	inner []T
}

func (s sortBy[T]) Len() int           { return len(s.inner) }
func (s sortBy[T]) Less(i, j int) bool { return s.less(s.inner[i], s.inner[j]) }
func (s sortBy[T]) Swap(i, j int)      { s.inner[i], s.inner[j] = s.inner[j], s.inner[i] }

// SortBy sorts the given slice in-place by the given less function.
func SortBy[T any](slice []T, less func(lhs, rhs T) bool) []T {
	sort.Sort(sortBy[T]{less: less, inner: slice})
	return slice
}

// Sort sorts the given slice in-place.
func Sort[T core.Ordered](slice []T) []T {
	return SortBy(slice, func(lhs, rhs T) bool { return lhs < rhs })
}

// ToHashMap converts the given slice to a map by the given key function.
func ToHashMap[
	T any,
	TKey comparable,
	TValue any,
	F func(T, int) (TKey, TValue),
](
	slice []T,
	f F,
) map[TKey]TValue {
	result := make(map[TKey]TValue)
	for i, v := range slice {
		key, value := f(v, i)
		result[key] = value
	}
	return result
}

// ToIndexedMap converts the given slice to a map from index to element.
func ToIndexedMap[T any](slice []T) map[int]T {
	return ToHashMap(slice, func(v T, i int) (int, T) { return i, v })
}

// ToHashSet returns a new set with the given slice.
func ToHashSet[T comparable](slice []T) map[T]struct{} {
	return sets.NewHashSet(slice...).ToMap()
}
