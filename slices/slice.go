package slices

import (
	"errors"
	"math/rand"

	"github.com/go-board/std/clone"
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/cond"
	"github.com/go-board/std/core"
	"github.com/go-board/std/operator"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/result"
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
	res := make([][]T, 0, len(slice)/chunk+1)
	for i := 0; i < size; i += chunk {
		if i+chunk > size {
			tmp := make([]T, len(slice[i:]))
			copy(tmp, slice[i:])
			res = append(res, tmp)
		} else {
			tmp := make([]T, chunk)
			copy(tmp, slice[i:i+chunk])
			res = append(res, tmp)
		}
	}
	return res
}

// Clone returns a new slice with the same elements as the given slice.
func Clone[T any](slice []T) []T {
	return Map(slice, func(t T) T { return t })
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
	res := make([]T, 0, len(lhs))
	for _, v := range lhs {
		// TODO: optimize use O(1) lookup
		if !ContainsBy(rhs, func(t T) bool { return eq(v, t) }) {
			res = append(res, v)
		}
	}
	return res
}

// Distinct returns a new slice with the given slice without duplicates.
func Distinct[T comparable](slice []T) []T {
	res := make([]T, 0)
	m := make(map[T]struct{})
	for _, elem := range slice {
		if _, ok := m[elem]; !ok {
			res = append(res, elem)
			m[elem] = struct{}{}
		}
	}
	return res
}

// DistinctBy returns a new slice with the distinct elements of the given slice by the given function.
func DistinctBy[T any, K comparable](slice []T, key func(T) K) []T {
	m := make(map[K]T)
	for _, v := range slice {
		m[key(v)] = v
	}
	used := make(map[K]bool)
	res := make([]T, 0, len(m))
	for _, v := range slice {
		if _, ok := used[key(v)]; !ok {
			used[key(v)] = true
			res = append(res, m[key(v)])
		}
	}
	return res
}

// Filter returns a new slice with all elements that satisfy the given predicate.
func Filter[T any](slice []T, f func(T) bool) []T {
	res := make([]T, 0, len(slice))
	for _, v := range slice {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// FilterIndexed returns a new slice with all elements that satisfy the given predicate.
func FilterIndexed[T any](slice []T, f func(T, int) bool) []T {
	res := make([]T, 0, len(slice))
	for i, v := range slice {
		if f(v, i) {
			res = append(res, v)
		}
	}
	return res
}

// Flatten returns a new slice with all elements in the given slice and all elements in all sub-slices.
func Flatten[T any](slice [][]T) []T {
	return FlattenBy(slice, func(t []T) []T { return t })
}

// FlattenBy returns a new slice with all elements in the given slice and all elements in the given slices.
func FlattenBy[T, S any](slice []T, f func(T) []S) []S {
	res := make([]S, 0, len(slice))
	for _, v := range slice {
		res = append(res, f(v)...)
	}
	return res
}

// Fold accumulates value starting with initial value and applying accumulator from left to right to current accum value and each element.
// Returns the final accum value or initial value if the slice is empty.
func Fold[T, A any](slice []T, initial A, accumulator func(A, T) A) A {
	res := initial
	for _, v := range slice {
		res = accumulator(res, v)
	}
	return res
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
	res := make(map[TKey][]TValue)
	for _, v := range slice {
		key, value := group(v)
		res[key] = append(res[key], value)
	}
	return res
}

// IntersectionBy returns a new slice with the elements that are in both given slices by the given function.
func IntersectionBy[T any](lhs []T, rhs []T, eq func(T, T) bool) []T {
	res := make([]T, 0, len(lhs))
	for _, v := range lhs {
		// TODO: optimize use O(1) lookup
		if ContainsBy(rhs, func(t T) bool { return eq(t, v) }) {
			res = append(res, v)
		}
	}
	return res
}

func Intersection[T comparable](lhs []T, rhs []T) []T {
	return IntersectionBy(lhs, rhs, operator.Eq[T])
}

// LastIndex returns the index of the last element in the given slice that same with the given element.
func LastIndex[T comparable](slice []T, v T) optional.Optional[int] {
	return LastIndexBy(slice, v, operator.Eq[T])
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
	res := make([]U, len(slice))
	for i, v := range slice {
		res[i] = f(v)
	}
	return res
}

// MapIndexed returns a new slice with the results of applying the given function to each element in the given slice.
func MapIndexed[T, U any](slice []T, f func(T, int) U) []U {
	res := make([]U, len(slice))
	for i, v := range slice {
		res[i] = f(v, i)
	}
	return res
}

// Max returns the maximum element in the given slice.
func Max[T core.Ordered](slice []T) optional.Optional[T] {
	return MaxBy(slice, operator.Lt[T])
}

// MaxBy returns the maximum element in the given slice that satisfies the given function.
func MaxBy[T any](slice []T, less func(T, T) bool) optional.Optional[T] {
	return Reduce(slice, func(lhs, rhs T) T {
		return cond.Ternary(less(lhs, rhs), rhs, lhs)
	})
}

func MaxByKey[T any, K cmp.Ordered](slice []T, keyFn func(T) K) optional.Optional[T] {
	return Reduce(slice, func(lhs, rhs T) T {
		return cond.Ternary(keyFn(lhs) < keyFn(rhs), rhs, lhs)
	})
}

// Min returns the minimum element in the given slice.
func Min[T core.Ordered](slice []T) optional.Optional[T] {
	return MinBy(slice, operator.Lt[T])
}

// MinBy returns the minimum element in the given slice that satisfies the given function.
func MinBy[T any](slice []T, less func(T, T) bool) optional.Optional[T] {
	return Reduce(slice, func(lhs, rhs T) T {
		return cond.Ternary(less(lhs, rhs), lhs, rhs)
	})
}

func MinByKey[T any, K cmp.Ordered](slice []T, keyFn func(T) K) optional.Optional[T] {
	return Reduce(slice, func(lhs, rhs T) T {
		return cond.Ternary(keyFn(lhs) < keyFn(rhs), lhs, rhs)
	})
}

// None returns true if no element in the given slice satisfies the given predicate.
func None[T any](slice []T, f func(T) bool) bool {
	return !Any(slice, f)
}

// Nth returns the nth element in the given slice.
//
// If n is negative, it returns the last element plus one.
// If n is greater than the length of the slice, it returns [optional.None].
func Nth[T any](slice []T, n int) optional.Optional[T] {
	if n < 0 {
		n = len(slice) + n
	}
	if n < 0 || n >= len(slice) {
		return optional.None[T]()
	}
	return optional.Some(slice[n])
}

// Partition split slice into two slices according to a predicate.
//
// The first slice will contain items for which the predicate returned true,
// and the second slice will contain items for which the predicate returned false.
//
// For Example:
//
//	Partition([]int{1, 2, 3}, func(s int) bool { return s % 2 == 0 })
//	returns: ([2], [1, 3])
func Partition[T any](slice []T, f func(T) bool) ([]T, []T) {
	lhs := make([]T, 0)
	rhs := make([]T, 0)
	for _, e := range slice {
		if f(e) {
			lhs = append(lhs, e)
		} else {
			rhs = append(rhs, e)
		}
	}
	return lhs, rhs
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

// Shuffle the given slice in-place.
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

// ToHashMap converts the given slice to a map by the given key function.
func ToHashMap[
	T any,
	TKey comparable,
	TValue any,
	F ~func(T, int) (TKey, TValue),
](
	slice []T,
	f F,
) map[TKey]TValue {
	res := make(map[TKey]TValue)
	for i, v := range slice {
		key, value := f(v, i)
		res[key] = value
	}
	return res
}

// ToIndexedMap converts the given slice to a map from index to element.
func ToIndexedMap[T any](slice []T) map[int]T {
	return ToHashMap(slice, operator.Exchange[T, int])
}

// ToHashSet returns a new set with the given slice.
func ToHashSet[T comparable](slice []T) map[T]struct{} {
	m := make(map[T]struct{}, len(slice))
	for _, elem := range slice {
		m[elem] = struct{}{}
	}
	return m
}

// First finds first element
func First[T any](slice []T) optional.Optional[T] {
	if len(slice) > 0 {
		return optional.Some(slice[0])
	}
	return optional.None[T]()
}

// Last finds last element
func Last[T any](slice []T) optional.Optional[T] {
	if len(slice) > 0 {
		return optional.Some(slice[len(slice)-1])
	}
	return optional.None[T]()
}

// SpliceFirst return first element and rest if len > 0, else return (None, []T)
func SpliceFirst[T any](slice []T) (optional.Optional[T], []T) {
	if len(slice) > 0 {
		return optional.Some(slice[0]), slice[1:]
	}
	return optional.None[T](), slice
}

// SpliceLast return last element and rest if len > 0, else return (None, []T)
func SpliceLast[T any](slice []T) (optional.Optional[T], []T) {
	if len(slice) > 0 {
		return optional.Some(slice[len(slice)-1]), slice[:len(slice)-1]
	}
	return optional.None[T](), slice
}

// FilterNonZero removes zero value from slice
//
// zero value are:
//
//	integer and float: 0
//	bool: false
//	string: empty string, aka: ""
//	pointer: nil pointer
//	struct with all field is zero value
//	interface: nil
//	chan/map/slice: nil
func FilterNonZero[T comparable](slice []T) []T {
	var empty T
	return Filter(slice, func(t T) bool { return t == empty })
}

func FirstNonZero[T comparable](slice []T) T {
	var zero T
	for _, v := range slice {
		if zero != v {
			return v
		}
	}
	return zero
}
