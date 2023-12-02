package slices

import (
	"errors"
	"math/rand"

	"github.com/go-board/std/clone"
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/core"
	"github.com/go-board/std/iter"
	"github.com/go-board/std/operator"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/result"
)

func ForwardSeq[E any, S ~[]E](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, x := range s {
			if !yield(x) {
				break
			}
		}
	}
}

func BackwardSeq[E any, S ~[]E](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(s[i]) {
				break
			}
		}
	}
}

func Collect[E any](s iter.Seq[E]) []E {
	rs := make([]E, 0)
	s(func(x E) bool {
		rs = append(rs, x)
		return true
	})
	return rs
}

// All returns true if all elements in the given slice satisfy the given predicate.
func All[T any, S ~[]T](slice S, f func(T) bool) bool {
	return iter.All(ForwardSeq(slice), f)
}

// Any returns true if any element in the given slice satisfies the given predicate.
func Any[T any, S ~[]T](slice S, f func(T) bool) bool {
	return iter.Any(ForwardSeq(slice), f)
}

// Chunk returns a new slice with the given slice split into smaller slices of the given size.
func Chunk[T any, S ~[]T](slice S, chunk int) []S {
	size := len(slice)
	res := make([]S, 0, len(slice)/chunk+1)
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
func Clone[T any, S ~[]T](slice []T) []T {
	return Map(slice, func(t T) T { return t })
}

// DeepClone returns a new slice with the cloned elements.
func DeepClone[T clone.Cloneable[T], S ~[]T](slice S) S {
	return DeepCloneBy(slice, func(t T) T { return t.Clone() })
}

// DeepCloneBy returns a new slice with the cloned elements as the given slice.
func DeepCloneBy[T any, S ~[]T](slice S, clone func(T) T) S {
	return Map(slice, clone)
}

// Difference returns a new slice with the elements that are in the first slice but not in the second.
func Difference[T comparable, S ~[]T](lhs S, rhs S) S {
	s := ToHashSet[T](rhs)
	x := make(S, 0)
	for _, e := range lhs {
		if _, ok := s[e]; !ok {
			x = append(x, e)
		}
	}
	return x
}

// DifferenceBy returns a new slice with the elements that are in the first slice but not in the second by the given function.
func DifferenceBy[T any, S1 ~[]T, S2 ~[]T](lhs S1, rhs S2, eq func(T, T) bool) []T {
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
func Distinct[T comparable, S ~[]T](slice S) []T {
	return DistinctBy(slice, func(x T) T { return x })
}

// DistinctBy returns a new slice with the distinct elements of the given slice by the given function.
func DistinctBy[T any, K comparable, S ~[]T](slice S, key func(T) K) []T {
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
func Filter[T any, S ~[]T](slice S, f func(T) bool) S {
	return Collect(iter.Filter(ForwardSeq(slice), f))
}

// FilterIndexed returns a new slice with all elements that satisfy the given predicate.
func FilterIndexed[T any, S ~[]T](slice S, f func(T, int) bool) S {
	res := make([]T, 0, len(slice))
	for i, v := range slice {
		if f(v, i) {
			res = append(res, v)
		}
	}
	return res
}

// Flatten returns a new slice with all elements in the given slice and all elements in all sub-slices.
func Flatten[T any, S ~[]T](slice []S) []T {
	return FlattenBy(slice, func(t S) []T { return t })
}

// FlattenBy returns a new slice with all elements in the given slice and all elements in the given slices.
// Deprecated: use FlatMap
func FlattenBy[T, E any, S ~[]T](slice S, f func(T) []E) []E {
	return FlatMap(slice, f)
}

// FlatMap returns a new slice with all elements in the given slice and all elements in the given slices.
func FlatMap[T, E any, S ~[]T](slice S, f func(T) []E) []E {
	res := make([]E, 0, len(slice))
	for _, v := range slice {
		res = append(res, f(v)...)
	}
	return res
}

// TryFold accumulates value starting with initial value and applying accumulator from left to right to current accum value and each element.
// Returns the final accum value or initial value if the slice is empty.
// If error occurred, return error early.
func TryFold[T, A any, S ~[]T](slice S, initial A, accumulator func(A, T) (A, error)) (res A, err error) {
	return iter.TryFold(ForwardSeq(slice), initial, accumulator)
}

// Fold accumulates value starting with initial value and applying accumulator from left to right to current accum value and each element.
// Returns the final accum value or initial value if the slice is empty.
func Fold[T, A any, S ~[]T](slice S, initial A, accumulator func(A, T) A) A {
	return iter.Fold(ForwardSeq(slice), initial, accumulator)
}

func TryFoldRight[T, A any, S ~[]T](slice S, initial A, accumulator func(A, T) (A, error)) (res A, err error) {
	return iter.TryFold(BackwardSeq(slice), initial, accumulator)
}

// FoldRight accumulates value starting with initial value and applying accumulator from right to left to current accum value and each element.
// Returns the final accum value or initial value if the slice is empty.
func FoldRight[T, A any, S ~[]T](slice S, initial A, accumulator func(A, T) A) A {
	return iter.Fold(BackwardSeq(slice), initial, accumulator)
}

// ForEach iterates over the given slice and calls the given function for each element.
func ForEach[T any, S ~[]T](slice S, f func(T)) {
	iter.ForEach(ForwardSeq(slice), f)
}

// ForEachIndexed iterates over the given slice and calls the given function for each element.
func ForEachIndexed[T any, S ~[]T](slice S, f func(T, int)) {
	for i, v := range slice {
		f(v, i)
	}
}

// GroupBy returns a new map with the given slice split into smaller slices of the given size.
func GroupBy[T any, TKey comparable, TValue any, S ~[]T](slice S, group func(T) (TKey, TValue)) map[TKey][]TValue {
	res := make(map[TKey][]TValue)
	for _, v := range slice {
		key, value := group(v)
		res[key] = append(res[key], value)
	}
	return res
}

// IntersectionBy returns a new slice with the elements that are in both given slices by the given function.
func IntersectionBy[T any, S1 ~[]T, S2 ~[]T](lhs S1, rhs S2, eq func(T, T) bool) []T {
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
	s := ToHashSet(lhs)
	res := make([]T, 0)
	for _, x := range rhs {
		if _, ok := s[x]; ok {
			res = append(res, x)
		}
	}
	return res
}

// LastIndex returns the index of the last element in the given slice that same with the given element.
func LastIndex[T comparable, S ~[]T](slice S, v T) optional.Optional[int] {
	return LastIndexBy(slice, v, operator.Eq[T])
}

// LastIndexBy returns the index of the last element in the given slice that satisfies the given predicate.
func LastIndexBy[T any, S ~[]T](slice S, v T, eq func(T, T) bool) optional.Optional[int] {
	for i := len(slice) - 1; i >= 0; i-- {
		if eq(v, slice[i]) {
			return optional.Some(i)
		}
	}
	return optional.None[int]()
}

func TryMap[T, U any, S ~[]T](slice S, f func(T) (U, error)) ([]U, error) {
	res := make([]U, len(slice))
	for i, v := range slice {
		x, err := f(v)
		if err != nil {
			return res, err
		}
		res[i] = x
	}
	return res, nil
}

// Map returns a new slice with the results of applying the given function to each element in the given slice.
func Map[T, U any, S ~[]T](slice S, f func(T) U) []U {
	return Collect(iter.Map(ForwardSeq(slice), f))
}

func TryMapIndexed[T, U any, S ~[]T](slice S, f func(T, int) (U, error)) ([]U, error) {
	res := make([]U, len(slice))
	for i, v := range slice {
		x, err := f(v, i)
		if err != nil {
			return res, err
		}
		res[i] = x
	}
	return res, nil
}

// MapIndexed returns a new slice with the results of applying the given function to each element in the given slice.
func MapIndexed[T, U any, S ~[]T](slice S, f func(T, int) U) []U {
	res := make([]U, len(slice))
	for i, v := range slice {
		res[i] = f(v, i)
	}
	return res
}

// Max returns the maximum element in the given slice.
func Max[T cmp.Ordered, S ~[]T](slice S) optional.Optional[T] {
	return optional.FromPair(iter.Max(ForwardSeq(slice)))
}

// MaxBy returns the maximum element in the given slice that satisfies the given function.
func MaxBy[T any, S ~[]T](slice S, f func(T, T) int) optional.Optional[T] {
	return optional.FromPair(iter.MaxFunc(ForwardSeq(slice), f))
}

func MaxByKey[T any, K cmp.Ordered, S ~[]T](slice S, keyFn func(T) K) optional.Optional[T] {
	return optional.FromPair(iter.MaxFunc(ForwardSeq(slice), func(x, y T) int {
		return cmp.Compare(keyFn(x), keyFn(y))
	}))
}

// Min returns the minimum element in the given slice.
func Min[T core.Ordered, S ~[]T](slice S) optional.Optional[T] {
	return optional.FromPair(iter.Min(ForwardSeq(slice)))
}

// MinBy returns the minimum element in the given slice that satisfies the given function.
func MinBy[T any, S ~[]T](slice S, f func(T, T) int) optional.Optional[T] {
	return optional.FromPair(iter.MinFunc(ForwardSeq(slice), f))
}

func MinByKey[T any, K cmp.Ordered, S ~[]T](slice S, keyFn func(T) K) optional.Optional[T] {
	return optional.FromPair(iter.MinFunc(ForwardSeq(slice), func(x, y T) int {
		return cmp.Compare(keyFn(x), keyFn(y))
	}))
}

// None returns true if no element in the given slice satisfies the given predicate.
func None[T any, S ~[]T](slice S, f func(T) bool) bool {
	return !Any(slice, f)
}

// Nth returns the nth element in the given slice.
//
// If n is negative, it returns the last element plus one.
// If n is greater than the length of the slice, it returns [optional.None].
func Nth[T any, S ~[]T](slice S, n int) optional.Optional[T] {
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
func Partition[T any, S ~[]T](slice S, f func(T) bool) ([]T, []T) {
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
func Reduce[T any, S ~[]T](slice S, f func(T, T) T) optional.Optional[T] {
	return optional.FromPair(iter.Reduce(ForwardSeq(slice), f))
}

// ReduceRight returns the result of applying the given function to each element in the given slice.
func ReduceRight[T any, S ~[]T](slice S, f func(T, T) T) optional.Optional[T] {
	return optional.FromPair(iter.Reduce(BackwardSeq(slice), f))
}

// Reverse returns a new slice with the elements in the given slice in reverse order.
func Reverse[T any, S ~[]T](slice S) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shuffle the given slice in-place.
func Shuffle[T any, S ~[]T](slice S) {
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}

// Single returns the single element,
// or return an error if the collection is empty or has more than one element.
func Single[T any, S ~[]T](slice S) result.Result[T] {
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
	S ~[]T,
](
	slice S,
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
func ToIndexedMap[T any, S ~[]T](slice S) map[int]T {
	return ToHashMap(slice, operator.Exchange[T, int])
}

// ToHashSet returns a new set with the given slice.
func ToHashSet[T comparable, S ~[]T](slice S) map[T]struct{} {
	m := make(map[T]struct{}, len(slice))
	for _, elem := range slice {
		m[elem] = struct{}{}
	}
	return m
}

// First finds first element
func First[T any, S ~[]T](slice S) optional.Optional[T] {
	if len(slice) > 0 {
		return optional.Some(slice[0])
	}
	return optional.None[T]()
}

// Last finds last element
func Last[T any, S ~[]T](slice S) optional.Optional[T] {
	if len(slice) > 0 {
		return optional.Some(slice[len(slice)-1])
	}
	return optional.None[T]()
}

// SpliceFirst return first element and rest if len > 0, else return (None, []T)
func SpliceFirst[T any, S ~[]T](slice S) (optional.Optional[T], []T) {
	if len(slice) > 0 {
		return optional.Some(slice[0]), slice[1:]
	}
	return optional.None[T](), slice
}

// SpliceLast return last element and rest if len > 0, else return (None, []T)
func SpliceLast[T any, S ~[]T](slice S) (optional.Optional[T], []T) {
	if len(slice) > 0 {
		return optional.Some(slice[len(slice)-1]), slice[:len(slice)-1]
	}
	return optional.None[T](), slice
}

// FirstNonZero returns first non zero value
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
func FirstNonZero[T comparable, S ~[]T](slice S) T {
	var zero T
	for _, v := range slice {
		if zero != v {
			return v
		}
	}
	return zero
}
