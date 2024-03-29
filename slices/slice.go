package slices

import (
	"errors"
	"math/rand"

	"github.com/go-board/std/clone"
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/collections/ordered"
	"github.com/go-board/std/constraints"
	"github.com/go-board/std/iter"
	"github.com/go-board/std/iter/collector"
	"github.com/go-board/std/operator"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/result"
	"github.com/go-board/std/tuple"
)

// Forward create a [iter.Seq] in forward order.
//
// Example:
//
//	slices.Forward([]int{1,2,3}) => seq: 1,2,3
func Forward[E any, S ~[]E](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, x := range s {
			if !yield(x) {
				break
			}
		}
	}
}

// Backward create a [iter.Seq] in backward order.
//
// Example:
//
//	slices.Backward([]int{1,2,3}) => seq: 3,2,1
func Backward[E any, S ~[]E](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(s[i]) {
				break
			}
		}
	}
}

// Collect all elements in [Seq] and return a slice.
//
// Example:
//
//	slices.Collect(seq(1,2,3)) => []int{1,2,3}
func Collect[E any](s iter.Seq[E]) []E {
	rs := make([]E, 0)
	iter.CollectFunc(s, func(e E) bool {
		rs = append(rs, e)
		return true
	})
	return rs
}

// All returns true if all elements in the given slice satisfy the given predicate.
//
// Example:
//
//	slices.All([]int{1,2,3}, func(x int) bool { return x > 0 }) => true
//	slices.All([]int{1,2,3}, func(x int) bool { return x > 2 }) => false
func All[T any, S ~[]T](slice S, f func(T) bool) bool {
	return iter.All(Forward(slice), f)
}

// Any returns true if any element in the given slice satisfies the given predicate.
//
// Example:
//
//	slices.Any([]int{1,2,3}, func(x int) bool { return x > 2 }) => true
//	slices.Any([]int{1,2,3}, func(x int) bool { return x > 6 }) => false
func Any[T any, S ~[]T](slice S, f func(T) bool) bool {
	return iter.Any(Forward(slice), f)
}

// Chunk returns a new slice with the given slice split into smaller slices of the given size.
//
// Example:
//
//	slices.Chunk(seq(1,2,3,4,5,6,7,8,9), 3) => [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
//	slices.Chunk(seq(1,2,3,4,5,6,7), 3)     => [][]int{{1,2,3}, {4,5,6}, {7}}
func Chunk[T any, S ~[]T](slice S, chunk int) []S {
	x := collector.Collect(Forward(slice), collector.Chunk[T](chunk))
	return Collect(iter.Map(x, func(e iter.Seq[T]) S { return Collect(e) }))
}

// Clone returns a new slice with the same elements as the given slice.
func Clone[T any, S ~[]T](slice S) S {
	return Collect(Forward(slice))
}

func CountValue[T comparable, S ~[]T](slice S, value T) int {
	return iter.Count(Forward(slice), value)
}

func CountBy[T any, S ~[]T](slice S, f func(T) bool) int {
	return iter.CountFunc(Forward(slice), f)
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
//
// Example:
//
//	slices.Difference([]int{1,2,3}, []int{3,4,5}) => []int{1,2}
//	slices.Difference([]int{1,2,3}, []int{1,2,3,4,5}) => []int{}
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
func DifferenceBy[T any, S ~[]T](lhs S, rhs S, cmp func(T, T) int) S {
	l := ordered.NewSet(cmp)
	l.InsertIter(Forward(lhs))
	r := ordered.NewSet(cmp)
	r.InsertIter(Forward(rhs))
	return Collect(l.Difference(r).AscendIter())
}

// Distinct returns a new slice with the given slice without duplicates.
//
// Example:
//
//	slices.Distinct([]int{1,2,3,2,1}) => []int{1,2,3}
func Distinct[T comparable, S ~[]T](slice S) S {
	return DistinctBy(slice, func(x T) T { return x })
}

// DistinctBy returns a new slice with the distinct elements of the given slice by the given function.
func DistinctBy[T any, K comparable, S ~[]T](slice S, key func(T) K) S {
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
//
// Example:
//
//	slices.Filter([]int{1,2,3,4,5,6}, func(x int) bool {return x%2 == 0}) => []int{2,4,6}
func Filter[T any, S ~[]T](slice S, f func(T) bool) S {
	return Collect(iter.Filter(Forward(slice), f))
}

// FilterIndexed returns a new slice with all elements that satisfy the given predicate.
func FilterIndexed[T any, S ~[]T](slice S, f func(T, int) bool) S {
	enumerated := iter.Enumerate(Forward(slice))
	filter := iter.Filter(enumerated, func(t tuple.Pair[int, T]) bool { return f(t.Second(), t.First()) })
	mapped := iter.Map(filter, func(e tuple.Pair[int, T]) T { return e.Second() })
	return Collect(mapped)
}

func FilterMap[A, B any, S ~[]A](slice S, f func(A) (B, bool)) []B {
	return Collect(iter.FilterMap(Forward(slice), f))
}

func FilterMapIndexed[A, B any, S ~[]A](slice S, f func(A, int) (B, bool)) []B {
	return Collect(iter.FilterMap(iter.Enumerate(Forward(slice)), func(t tuple.Pair[int, A]) (B, bool) {
		return f(t.Second(), t.First())
	}))
}

// Flatten returns a new slice with all elements in the given slice and all elements in all sub-slices.
func Flatten[T any, S ~[]T, X ~[]S](slice X) S {
	return FlatMap(slice, func(t S) []T { return t })
}

// FlattenBy returns a new slice with all elements in the given slice and all elements in the given slices.
// Deprecated: use FlatMap
func FlattenBy[T, E any, S ~[]T](slice S, f func(T) []E) []E {
	return FlatMap(slice, f)
}

// FlatMap returns a new slice with all elements in the given slice and all elements in the given slices.
func FlatMap[T, E any, S ~[]T](slice S, f func(T) []E) []E {
	return Collect(iter.FlatMap(Forward(slice), func(x T) iter.Seq[E] { return Forward(f(x)) }))
}

// TryFold accumulates value starting with initial value and applying
// accumulator from left to right to current accum value and each element.
//
// Returns the final accum value or initial value if the slice is empty.
// If error occurred, return error early.
func TryFold[T, A any, S ~[]T](slice S, initial A, accumulator func(A, T) (A, error)) (res A, err error) {
	return TryFoldLeft(slice, initial, accumulator)
}

// Fold accumulates value starting with initial value and applying accumulator from left to right to current accum value and each element.
// Returns the final accum value or initial value if the slice is empty.
func Fold[T, A any, S ~[]T](slice S, initial A, accumulator func(A, T) A) A {
	return FoldLeft(slice, initial, accumulator)
}

func TryFoldLeft[T, A any, S ~[]T](slice S, initial A, accumulator func(A, T) (A, error)) (A, error) {
	return iter.TryFold(Forward(slice), initial, accumulator)
}

func FoldLeft[T, A any, S ~[]T](slice S, initial A, accumulator func(A, T) A) A {
	return iter.Fold(Forward(slice), initial, accumulator)
}

func TryFoldRight[T, A any, S ~[]T](slice S, initial A, accumulator func(A, T) (A, error)) (res A, err error) {
	return iter.TryFold(Backward(slice), initial, accumulator)
}

// FoldRight accumulates value starting with initial value and applying accumulator from right to left to current accum value and each element.
// Returns the final accum value or initial value if the slice is empty.
func FoldRight[T, A any, S ~[]T](slice S, initial A, accumulator func(A, T) A) A {
	return iter.Fold(Backward(slice), initial, accumulator)
}

// ForEach iterates over the given slice and calls the given function for each element.
func ForEach[T any, S ~[]T](slice S, f func(T)) {
	iter.ForEach(Forward(slice), f)
}

// ForEachIndexed iterates over the given slice and calls the given function for each element.
func ForEachIndexed[T any, S ~[]T](slice S, f func(T, int)) {
	iter.ForEach(iter.Enumerate(Forward(slice)), func(t tuple.Pair[int, T]) { f(t.Second(), t.First()) })
}

// GroupBy returns a new map with the given slice split into smaller slices of the given size.
func GroupBy[T any, TKey comparable, S ~[]T](slice S, f func(T) TKey) map[TKey]S {
	x := collector.Collect(Forward(slice), collector.GroupBy(f))
	res := make(map[TKey]S)
	iter.ForEach(x, func(t tuple.Pair[TKey, iter.Seq[T]]) {
		iter.ForEach(t.Second(), func(e T) {
			res[t.First()] = append(res[t.First()], e)
		})
	})
	return res
}

// IntersectionBy returns a new slice with the elements that are in both given slices by the given function.
func IntersectionBy[T any, S ~[]T](lhs S, rhs S, cmp func(T, T) int) S {
	l := ordered.NewSet(cmp)
	l.InsertIter(Forward(lhs))
	r := ordered.NewSet(cmp)
	r.InsertIter(Forward(rhs))
	return Collect(l.Intersection(r).AscendIter())
}

// Intersection returns a new slice with the elements that are in both give slices.
func Intersection[T comparable, S ~[]T](lhs S, rhs S) S {
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
	return LastIndexBy(slice, func(t T) bool { return t == v })
}

// LastIndexBy returns the index of the last element in the given slice that satisfies the given predicate.
func LastIndexBy[T any, S ~[]T](slice S, f func(T) bool) optional.Optional[int] {
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return optional.Some(i)
		}
	}
	return optional.None[int]()
}

// TryMap returns a new slice and an error.
//
// Stopping at first error occurred.
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
//
// Example:
//
//	slices.Map([]int{1,2,3}, strconv.Itoa) => []string{"1", "2", "3"}
func Map[T, U any, S ~[]T](slice S, f func(T) U) []U {
	return Collect(iter.Map(Forward(slice), f))
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
	return Collect(iter.Map(iter.Enumerate(Forward(slice)), func(x tuple.Pair[int, T]) U { return f(x.Second(), x.First()) }))
}

// Max returns the maximum element in the given slice.
//
// Example:
//
//	slices.Max([]int{1,2,1}) => 2
func Max[T cmp.Ordered, S ~[]T](slice S) optional.Optional[T] {
	return iter.MaxOption(Forward(slice))
}

// MaxBy returns the maximum element in the given slice that satisfies the given function.
func MaxBy[T any, S ~[]T](slice S, f func(T, T) int) optional.Optional[T] {
	return iter.MaxFuncOption(Forward(slice), f)
}

func MaxByKey[T any, K cmp.Ordered, S ~[]T](slice S, keyFn func(T) K) optional.Optional[T] {
	return iter.MaxByKeyOption(Forward(slice), keyFn)
}

// Min returns the minimum element in the given slice.
//
// Example:
//
//	slices.Min([]int{1,2,1}) => 1
func Min[T constraints.Ordered, S ~[]T](slice S) optional.Optional[T] {
	return iter.MinOption(Forward(slice))
}

// MinBy returns the minimum element in the given slice that satisfies the given function.
func MinBy[T any, S ~[]T](slice S, f func(T, T) int) optional.Optional[T] {
	return iter.MinFuncOption(Forward(slice), f)
}

func MinByKey[T any, K cmp.Ordered, S ~[]T](slice S, keyFn func(T) K) optional.Optional[T] {
	return iter.MinByKeyOption(Forward(slice), keyFn)
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
	return iter.NthOption(Forward(slice), n)
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
func Partition[T any, S ~[]T](slice S, f func(T) bool) (S, S) {
	x := iter.Partition(Forward(slice), f)
	return Collect(x.First()), Collect(x.Second())
}

// Reduce returns the result of applying the given function to each element in the given slice.
//
// Example:
//
//	slices.Reduce([]int{1,2,3}, func(x, y int) int {return x+y}) => 6
func Reduce[T any, S ~[]T](slice S, f func(T, T) T) optional.Optional[T] {
	return ReduceLeft(slice, f)
}

func ReduceLeft[T any, S ~[]T](slice S, f func(T, T) T) optional.Optional[T] {
	return iter.ReduceOption(Forward(slice), f)
}

// ReduceRight returns the result of applying the given function to each element in the given slice.
func ReduceRight[T any, S ~[]T](slice S, f func(T, T) T) optional.Optional[T] {
	return iter.ReduceOption(Backward(slice), f)
}

// Reverse returns a new slice with the elements in the given slice in reverse order.
//
// Example:
//
//	slices.Reverse([]int{1,2,3}) => []int{3,2,1}
func Reverse[T any, S ~[]T](slice S) S {
	return Collect(Backward(slice))
}

// Shuffle the given slice in-place.
func Shuffle[T any, S ~[]T](slice S) S {
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	return slice
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
	return iter.FirstOption(Forward(slice), func(T) bool { return true })
}

func FirstFunc[E any, S ~[]E](slice S, f func(E) bool) optional.Optional[E] {
	return iter.FirstOption(Forward(slice), f)
}

// Last finds last element
func Last[T any, S ~[]T](slice S) optional.Optional[T] {
	return iter.LastOption(Forward(slice), func(T) bool { return true })
}

func LastFunc[E any, S ~[]E](slice S, f func(E) bool) optional.Optional[E] {
	return iter.LastOption(Forward(slice), f)
}

// SpliceFirst return first element and rest if len > 0, else return (None, []T)
//
// Example:
//
//	slices.SpliceFirst([]int{1,2,3}) => Some(1), []int{2,3}
//	slices.SpliceFirst([]int{})      => None, []int{}
func SpliceFirst[T any, S ~[]T](slice S) (optional.Optional[T], S) {
	if len(slice) > 0 {
		return optional.Some(slice[0]), slice[1:]
	}
	return optional.None[T](), slice
}

// SpliceLast return last element and rest if len > 0, else return (None, []T)
//
// Example:
//
//	slices.SpliceLast([]int{1,2,3}) => Some(1), []int{1,2}
//	slices.SpliceLast([]int{})      => None, []int{}
func SpliceLast[T any, S ~[]T](slice S) (optional.Optional[T], S) {
	if len(slice) > 0 {
		return optional.Some(slice[len(slice)-1]), slice[:len(slice)-1]
	}
	return optional.None[T](), slice
}

// FirstNonZero returns first non-zero value
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
	return iter.HeadOption(iter.FilterZero(Forward(slice))).ValueOrZero()
}
