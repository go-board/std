package iter

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/go-board/std/cmp"
	"github.com/go-board/std/cond"
	"github.com/go-board/std/operator"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/tuple"
)

type AllOfIter[T any] interface {
	Iter[T]
	AllOf(predicate func(T) bool) bool
}

// AllOf tests if every element of the iterator matches a predicate.
//
// allOf() takes a closure that returns true or false. It applies this
// closure to each element of the iterator, and if they all return true,
// then so does allOf(). If any of them return false, it returns false.
//
// allOf() is short-circuiting; in other words, it will stop processing
// as soon as it finds a false, given that no matter what else happens,
// the result will also be false.
//
// An empty iterator returns true.
//
// Examples:
//
//	s := []int{1, 2, 3}
//	assertTrue(iter.AllOf(slices.Iter(s)), func(x int) bool { return x < 10 })
//	assertFalse(iter.AllOf(slices.Iter(s)), func(x int) bool { return x > 100 })
func AllOf[T any](iter Iter[T], predicate func(T) bool) bool {
	if allOfIter, ok := iter.(AllOfIter[T]); ok {
		return allOfIter.AllOf(predicate)
	}
	return FoldLeft(iter, true, func(b bool, t T) bool { return b && predicate(t) })
}

type AnyOfIter[T any] interface {
	Iter[T]
	AnyOf(predicate func(T) bool) bool
}

// AnyOf tests if any element of the iterator matches a predicate.
//
// anyOf() takes a closure that returns true or false.
// It applies this closure to each element of the iterator,
// and if any of them return true, then so does any().
// If they all return false, it returns false.
//
// anyOf() is short-circuiting; in other words,
// it will stop processing as soon as it finds a true,
// given that no matter what else happens, the result will also be true.
//
// An empty iterator returns false.
//
// Examples:
//
//	s := []int{1, 2, 3}
//	assertTrue(iter.Any(slices.Iter(s)), func(x int) bool { return x%2 == 0 })
//	assertFalse(iter.Any(slices.Iter(s)), func(x int) bool { return x > 100 })
func AnyOf[T any](iter Iter[T], predicate func(T) bool) bool {
	if anyOfIter, ok := iter.(AnyOfIter[T]); ok {
		return anyOfIter.AnyOf(predicate)
	}
	return FoldLeft(iter, false, func(b bool, t T) bool { return b || predicate(t) })
}

type NoneOfIter[T any] interface {
	Iter[T]
	NoneOf(predicate func(T) bool) bool
}

// NoneOf tests if none element of the iterator matches a predicate.
//
// noneOf() takes a closure that returns true or false.
// It applies this closure to each element of the iterator,
// and if none of them return true, then so does none().
// If they all return false, it returns true.
//
// An empty iterator returns true.
//
// Examples:
//
//	s := []int{1, 2, 3}
//	assertTrue(iter.None(slices.Iter(s)), func(x int) bool { return x%2 == 0 })
//	assertFalse(iter.None(slices.Iter(s)), func(x int) bool { return x > 100 })
func NoneOf[T any](iter Iter[T], predicate func(T) bool) bool {
	if noneOfIter, ok := iter.(NoneOfIter[T]); ok {
		return noneOfIter.NoneOf(predicate)
	}
	if anyOfIter, ok := iter.(AnyOfIter[T]); ok {
		return !anyOfIter.AnyOf(predicate)
	}
	return FoldLeft(iter, true, func(b bool, t T) bool {
		return b && !predicate(t)
	})
}

type MapIter[T, U any] interface {
	Iter[T]
	Map(f func(T) U) Iter[U]
}

func Map[T, U any](iter Iter[T], f func(T) U) Iter[U] {
	if mapIter, ok := iter.(MapIter[T, U]); ok {
		return mapIter.Map(f)
	}
	return IterFunc[U](func() optional.Optional[U] {
		return optional.Map(iter.Next(), f)
	})
}

type MapWhileIter[T, U any] interface {
	Iter[T]
	MapWhile(f func(T) optional.Optional[U]) Iter[U]
}

func MapWhile[T, U any](iter Iter[T], f func(T) optional.Optional[U]) Iter[U] {
	if mapWhileIter, ok := iter.(MapWhileIter[T, U]); ok {
		return mapWhileIter.MapWhile(f)
	}
	return IterFunc[U](func() optional.Optional[U] {
		item := iter.Next()
		if item.IsNone() {
			return optional.None[U]()
		}
		return f(item.Value())
	})
}

type FilterIter[T any] interface {
	Iter[T]
	Filter(f func(T) bool) Iter[T]
}

func Filter[T any](iter Iter[T], f func(T) bool) Iter[T] {
	if filterIter, ok := iter.(FilterIter[T]); ok {
		return filterIter.Filter(f)
	}
	return IterFunc[T](func() optional.Optional[T] {
		for {
			item := iter.Next()
			if item.IsNone() {
				return optional.None[T]()
			}
			if f(item.Value()) {
				return optional.Some(item.Value())
			}
		}
	})
}

type FilterMapIter[T, U any] interface {
	Iter[T]
	FilterMap(f func(T) optional.Optional[U]) Iter[U]
}

func FilterMap[T, U any](iter Iter[T], f func(T) optional.Optional[U]) Iter[U] {
	if filterMapIter, ok := iter.(FilterMapIter[T, U]); ok {
		return filterMapIter.FilterMap(f)
	}
	return IterFunc[U](func() optional.Optional[U] {
		for {
			item := iter.Next()
			if item.IsNone() {
				return optional.None[U]()
			}
			mapped := f(item.Value())
			if mapped.IsSome() {
				return mapped
			}
		}
	})
}

type FindLeftIter[T any] interface {
	Iter[T]
	FindLeft(f func(T) bool) optional.Optional[T]
}

func FindLeft[T any](iter Iter[T], f func(T) bool) optional.Optional[T] {
	if findIter, ok := iter.(FindLeftIter[T]); ok {
		return findIter.FindLeft(f)
	}
	for item := iter.Next(); item.IsSome(); item = iter.Next() {
		if f(item.Value()) {
			return item
		}
	}
	return optional.None[T]()
}

type FindRightIter[T any] interface {
	Iter[T]
	FindRight(f func(T) bool) optional.Optional[T]
}

func FindRight[T any](iter PrevIter[T], f func(T) bool) optional.Optional[T] {
	if findIter, ok := iter.(FindRightIter[T]); ok {
		return findIter.FindRight(f)
	}
	for item := iter.Prev(); item.IsSome(); item = iter.Prev() {
		if f(item.Value()) {
			return item
		}
	}
	return optional.None[T]()
}

type FindMapIter[T, U any] interface {
	Iter[T]
	FindMap(f func(T) optional.Optional[U]) optional.Optional[U]
}

func FindMap[T, U any](iter Iter[T], f func(T) optional.Optional[U]) optional.Optional[U] {
	if findMapIter, ok := iter.(FindMapIter[T, U]); ok {
		return findMapIter.FindMap(f)
	}
	for item := iter.Next(); item.IsSome(); item = iter.Next() {
		u := f(item.Value())
		if u.IsSome() {
			return u
		}
	}
	return optional.None[U]()
}

type TakeIter[T any] interface {
	Iter[T]
	Take(n int) Iter[T]
}

func Take[T any](iter Iter[T], n int) Iter[T] {
	if takeIter, ok := iter.(TakeIter[T]); ok {
		return takeIter.Take(n)
	}
	current := -1
	return IterFunc[T](func() optional.Optional[T] {
		current++
		if current < n {
			return iter.Next()
		}
		return optional.None[T]()
	})
}

type SkipIter[T any] interface {
	Iter[T]
	Skip(n int) Iter[T]
}

func Skip[T any](iter Iter[T], n int) Iter[T] {
	if skipIter, ok := iter.(SkipIter[T]); ok {
		return skipIter.Skip(n)
	}
	i := 0
	for i < n {
		iter.Next()
		i++
	}
	return iter
}

type TakeWhileIter[T any] interface {
	Iter[T]
	TakeWhile(take func(T) bool) Iter[T]
}

type SkipWhileIter[T any] interface {
	Iter[T]
	SkipWhile(skip func(T) bool) Iter[T]
}

type StepByIter[T any] interface {
	Iter[T]
	StepBy(step int) Iter[T]
}

func StepBy[T any](iter Iter[T], step int) Iter[T] {
	if stepByIter, ok := iter.(StepByIter[T]); ok {
		return stepByIter.StepBy(step)
	}
	n := 0
	finished := false
	return IterFunc[T](func() optional.Optional[T] {
		var none optional.Optional[T]
		if finished {
			return none
		}
		for {
			item := iter.Next()
			if item.IsNone() {
				finished = true
				return none
			}
			if n%step == 0 {
				return item
			}
			n++
		}
	})
}

type TailIter[T any] interface {
	Iter[T]
	Tail() Iter[T]
}

type PartitionIter[T any] interface {
	Iter[T]
	Partition(partition func(T) bool) (Iter[T], Iter[T])
}

func Partition[T any](iter Iter[T], partition func(T) bool) (Iter[T], Iter[T]) {
	if partitionIter, ok := iter.(PartitionIter[T]); ok {
		return partitionIter.Partition(partition)
	}
	trueSlice := make([]T, 0)
	falseSlice := make([]T, 0)
	ForEach(iter, func(t T) {
		if partition(t) {
			trueSlice = append(trueSlice, t)
		} else {
			falseSlice = append(falseSlice, t)
		}
	})
	return OfSlice(trueSlice), OfSlice(falseSlice)
}

type MaxIter[T cmp.Ordered] interface {
	Iter[T]
	Max() optional.Optional[T]
}

func Max[T cmp.Ordered](iter Iter[T]) optional.Optional[T] {
	if maxIter, ok := iter.(MaxIter[T]); ok {
		return maxIter.Max()
	}
	return ReduceLeft(iter, cmp.MaxOrdered[T])
}

type MaxByIter[T any] interface {
	Iter[T]
	MaxBy(cmp func(T, T) int) optional.Optional[T]
}

func MaxBy[T any](iter Iter[T], cmp func(T, T) int) optional.Optional[T] {
	if maxByIter, ok := iter.(MaxByIter[T]); ok {
		return maxByIter.MaxBy(cmp)
	}
	return ReduceLeft(iter, func(lhs T, rhs T) T {
		if cmp(lhs, rhs) > 0 {
			return lhs
		}
		return rhs
	})
}

type MaxByKeyIter[T any, K cmp.Ordered] interface {
	Iter[T]
	MaxByKey(keyFn func(T) K) optional.Optional[T]
}

func MaxByKey[T any, K cmp.Ordered](iter Iter[T], keyFn func(T) K) optional.Optional[T] {
	if maxByKeyIter, ok := iter.(MaxByKeyIter[T, K]); ok {
		return maxByKeyIter.MaxByKey(keyFn)
	}
	return ReduceLeft(iter, func(lhs T, rhs T) T {
		if cmp.Compare(keyFn(lhs), keyFn(rhs)) > 0 {
			return lhs
		}
		return rhs
	})
}

type MinIter[T any] interface {
	Iter[T]
	Min() optional.Optional[T]
}

func Min[T cmp.Ordered](iter Iter[T]) optional.Optional[T] {
	if minIter, ok := iter.(MinIter[T]); ok {
		return minIter.Min()
	}
	return ReduceLeft(iter, cmp.MinOrdered[T])
}

type MinByIter[T any] interface {
	Iter[T]
	MinBy(cmp func(T, T) int) optional.Optional[T]
}

func MinBy[T any](iter Iter[T], cmp func(T, T) int) optional.Optional[T] {
	if minByIter, ok := iter.(MinByIter[T]); ok {
		return minByIter.MinBy(cmp)
	}
	return ReduceLeft(iter, func(lhs T, rhs T) T {
		if cmp(lhs, rhs) < 0 {
			return lhs
		}
		return rhs
	})
}

type MinByKeyIter[T any, K cmp.Ordered] interface {
	Iter[T]
	MinByKey(keyFn func(T) K) optional.Optional[T]
}

func MinByKey[T any, K cmp.Ordered](iter Iter[T], keyFn func(T) K) optional.Optional[T] {
	if minByKeyIter, ok := iter.(MinByKeyIter[T, K]); ok {
		return minByKeyIter.MinByKey(keyFn)
	}
	return ReduceLeft(iter, func(lhs T, rhs T) T {
		if cmp.Compare(keyFn(lhs), keyFn(rhs)) < 0 {
			return lhs
		}
		return rhs
	})
}

func ternary[T any](ok bool, lhs T, rhs T) T {
	if ok {
		return lhs
	}
	return rhs
}

func minmax[T any](iter Iter[T], cmp func(T, T) int) optional.Optional[tuple.Pair[T, T]] {
	first := iter.Next()
	if first.IsNone() {
		return optional.None[tuple.Pair[T, T]]()
	}
	return optional.Some(FoldLeft[T, tuple.Pair[T, T]](iter, tuple.PairOf(first.Value(), first.Value()), func(a tuple.Pair[T, T], t T) tuple.Pair[T, T] {
		min := ternary(cmp(a.First, t) > 0, t, a.First)
		max := ternary(cmp(a.Second, t) > 0, a.Second, t)
		return tuple.PairOf(min, max)
	}))
}

type MinMaxIter[T cmp.Ordered] interface {
	Iter[T]
	MinMax() optional.Optional[tuple.Pair[T, T]]
}

func MinMax[T cmp.Ordered](iter Iter[T]) optional.Optional[tuple.Pair[T, T]] {
	if minMaxIter, ok := iter.(MinMaxIter[T]); ok {
		return minMaxIter.MinMax()
	}
	return minmax(iter, cmp.Compare[T])
}

type MinMaxByIter[T any] interface {
	Iter[T]
	MinMaxBy(cmp func(T, T) int) optional.Optional[tuple.Pair[T, T]]
}

func MinMaxBy[T any](iter Iter[T], cmp func(T, T) int) optional.Optional[tuple.Pair[T, T]] {
	if minMaxByIter, ok := iter.(MinMaxByIter[T]); ok {
		return minMaxByIter.MinMaxBy(cmp)
	}
	return minmax(iter, cmp)
}

type MinMaxByKeyIter[T any, K cmp.Ordered] interface {
	Iter[T]
	MinMaxByKey(keyFn func(T) K) optional.Optional[tuple.Pair[T, T]]
}

func MinMaxByKey[T any, K cmp.Ordered](iter Iter[T], keyFn func(T) K) optional.Optional[tuple.Pair[T, T]] {
	if minMaxByKeyIter, ok := iter.(MinMaxByKeyIter[T, K]); ok {
		return minMaxByKeyIter.MinMaxByKey(keyFn)
	}
	return minmax(iter, func(lhs T, rhs T) int { return cmp.Compare(keyFn(lhs), keyFn(rhs)) })
}

func compareIter[T any, U any](iter Iter[T], other Iter[U], cmp func(T, U) int) int {
	for {
		lvalue := iter.Next()
		rvalue := other.Next()
		if lvalue.IsSome() && rvalue.IsSome() {
			if c := cmp(lvalue.Value(), rvalue.Value()); c != 0 {
				return c
			}
		} else if lvalue.IsNone() && rvalue.IsNone() {
			return 0
		} else if lvalue.IsSome() {
			return +1
		} else {
			return -1
		}
	}
}

type CmpIter[T cmp.Ordered] interface {
	Iter[T]
	Cmp(other Iter[T]) int
}

func Cmp[T cmp.Ordered](iter Iter[T], other Iter[T]) int {
	if cmpIter, ok := iter.(CmpIter[T]); ok {
		return cmpIter.Cmp(other)
	}
	return compareIter(iter, other, cmp.Compare[T])
}

type CmpByIter[T any, U any] interface {
	Iter[T]
	CmpBy(other Iter[U], cmp func(T, U) int) int
}

func CmpBy[T any, U any](iter Iter[T], other Iter[U], cmp func(T, U) int) int {
	if cmpByIter, ok := iter.(CmpByIter[T, U]); ok {
		return cmpByIter.CmpBy(other, cmp)
	}
	return compareIter(iter, other, cmp)
}

type CmpByKeyIter[T any, K cmp.Ordered] interface {
	Iter[T]
	CmpByKey(other Iter[T], keyFn func(T) K) int
}

func CmpByKey[T any, K cmp.Ordered](iter Iter[T], other Iter[T], keyFn func(T) K) int {
	if cmpByKeyIter, ok := iter.(CmpByKeyIter[T, K]); ok {
		return cmpByKeyIter.CmpByKey(other, keyFn)
	}
	return compareIter(iter, other, func(lhs T, rhs T) int {
		return cmp.Compare(keyFn(lhs), keyFn(rhs))
	})
}

func isSorted[T any](iter Iter[T], cmp func(T, T) int) bool {
	prev := iter.Next()
	if prev.IsNone() {
		return true
	}
	for item := iter.Next(); item.IsSome(); item = iter.Next() {
		if cmp(prev.Value(), item.Value()) > 0 {
			return false
		}
		prev = item
	}
	return true
}

type IsSortedIter[T cmp.Ordered] interface {
	Iter[T]
	IsSorted() bool
}

func IsSorted[T cmp.Ordered](iter Iter[T]) bool {
	if isSortedIter, ok := iter.(IsSortedIter[T]); ok {
		return isSortedIter.IsSorted()
	}
	return isSorted(iter, cmp.Compare[T])
}

type IsSortedByIter[T any] interface {
	Iter[T]
	IsSortedBy(cmp func(T, T) int) bool
}

func IsSortedBy[T any](iter Iter[T], cmp func(T, T) int) bool {
	if isSortedByIter, ok := iter.(IsSortedByIter[T]); ok {
		return isSortedByIter.IsSortedBy(cmp)
	}
	return isSorted(iter, cmp)
}

type IsSortedByKeyIter[T any, K cmp.Ordered] interface {
	Iter[T]
	IsSortedByKey(keyFn func(T) K) bool
}

func IsSortedByKey[T any, K cmp.Ordered](iter Iter[T], keyFn func(T) K) bool {
	if isSortedByIter, ok := iter.(IsSortedByKeyIter[T, K]); ok {
		return isSortedByIter.IsSortedByKey(keyFn)
	}
	return isSorted(iter, func(t1, t2 T) int { return cmp.Compare(keyFn(t1), keyFn(t2)) })
}

type FoldLeftIter[T any, A any] interface {
	Iter[T]
	FoldLeft(init A, f func(A, T) A) A
}

func FoldLeft[T any, A any](iter Iter[T], init A, f func(A, T) A) A {
	if foldIter, ok := iter.(FoldLeftIter[T, A]); ok {
		return foldIter.FoldLeft(init, f)
	}
	for item := iter.Next(); item.IsSome(); item = iter.Next() {
		init = f(init, item.Value())
	}
	return init
}

func ReduceLeft[T any](iter Iter[T], f func(T, T) T) optional.Optional[T] {
	first := iter.Next()
	if first.IsNone() {
		return optional.None[T]()
	}
	return optional.Some(FoldLeft(iter, first.Value(), f))
}

type FoldRightIter[T any, A any] interface {
	PrevIter[T]
	FoldRight(init A, f func(A, T) A) A
}

func FoldRight[T any, A any](iter PrevIter[T], init A, f func(A, T) A) A {
	if foldRightIter, ok := iter.(FoldRightIter[T, A]); ok {
		return foldRightIter.FoldRight(init, f)
	}
	for item := iter.Prev(); item.IsSome(); item = iter.Prev() {
		init = f(init, item.Value())
	}
	return init
}

func ReduceRight[T any](iter PrevIter[T], f func(T, T) T) optional.Optional[T] {
	first := iter.Prev()
	if first.IsNone() {
		return optional.None[T]()
	}
	return optional.Some(FoldRight(iter, first.Value(), f))
}

func equalIter[T any, U any](iter Iter[T], other Iter[U], eq func(T, U) bool) bool {
	for {
		lvalue := iter.Next()
		rvalue := other.Next()
		if lvalue.IsSome() && rvalue.IsSome() {
			if !eq(lvalue.Value(), rvalue.Value()) {
				return false
			}
		} else if lvalue.IsNone() && rvalue.IsNone() {
			return true
		} else if lvalue.IsSome() {
			return false
		} else {
			return false
		}
	}
}

type EqualIter[T comparable] interface {
	Iter[T]
	Equal(other Iter[T]) bool
}

func Equal[T comparable](iter Iter[T], other Iter[T]) bool {
	if equalIter, ok := iter.(EqualIter[T]); ok {
		return equalIter.Equal(other)
	}
	return equalIter(iter, other, operator.Eq[T])
}

type EqualByIter[T, U any] interface {
	Iter[T]
	EqualBy(other Iter[U], eq func(T, U) bool) bool
}

func EqualBy[T, U any](iter Iter[T], other Iter[U], eq func(T, U) bool) bool {
	if equalByIter, ok := iter.(EqualByIter[T, U]); ok {
		return equalByIter.EqualBy(other, eq)
	}
	return equalIter(iter, other, eq)
}

type EqualByKeyIter[T any, K comparable] interface {
	Iter[T]
	EqualByKey(other Iter[T], keyFn func(T) K) bool
}

func EqualByKey[T any, K comparable](iter Iter[T], other Iter[T], keyFn func(T) K) bool {
	if equalByKeyIter, ok := iter.(EqualByKeyIter[T, K]); ok {
		return equalByKeyIter.EqualByKey(other, keyFn)
	}
	return equalIter(iter, other, func(lhs T, rhs T) bool { return keyFn(lhs) == keyFn(rhs) })
}

type EnumerateIter[T any] interface {
	Iter[T]
	Enumerate() Iter[tuple.Pair[int, T]]
}

func Enumerate[T any](iter Iter[T]) Iter[tuple.Pair[int, T]] {
	if enumerateIter, ok := iter.(EnumerateIter[T]); ok {
		return enumerateIter.Enumerate()
	}
	current := -1
	return IterFunc[tuple.Pair[int, T]](func() optional.Optional[tuple.Pair[int, T]] {
		current++
		return optional.Map(iter.Next(), func(t T) tuple.Pair[int, T] {
			return tuple.PairOf(current, t)
		})
	})
}

type collectSlice[T any] struct {
	slice []T
}

func (s *collectSlice[T]) Extend(iter Iter[T]) {
	for item := iter.Next(); item.IsSome(); item = iter.Next() {
		s.slice = append(s.slice, item.Value())
	}
}

func CollectToSlice[T any](iter Iter[T]) []T {
	slice := &collectSlice[T]{}
	CollectInto(iter, slice)
	return slice.slice
}

type Extend[T any] interface{ Extend(iter Iter[T]) }

func CollectInto[T any, E Extend[T]](iter Iter[T], collection E) E {
	collection.Extend(iter)
	return collection
}

func Chain[T any](iter Iter[T], other Iter[T]) Iter[T] {
	var finish atomic.Bool
	return IterFunc[T](func() optional.Optional[T] {
		if !finish.Load() {
			item := iter.Next()
			if item.IsNone() {
				finish.Store(true)
				return other.Next()
			}
			return item
		} else {
			return other.Next()
		}
	})
}

type ForEachIter[T any] interface {
	Iter[T]
	ForEach(f func(T))
}

func ForEach[T any](iter Iter[T], f func(T)) {
	if forEachIter, ok := iter.(ForEachIter[T]); ok {
		forEachIter.ForEach(f)
	} else {
		for item := iter.Next(); item.IsSome(); item = iter.Next() {
			f(item.Value())
		}
	}
}

type ZipByIter[A, B, C any] interface {
	Iter[A]
	ZipBy(other Iter[B], f func(A, B) C) Iter[C]
}

func ZipBy[A, B, C any](iter Iter[A], other Iter[B], f func(A, B) C) Iter[C] {
	if zipByIter, ok := iter.(ZipByIter[A, B, C]); ok {
		return zipByIter.ZipBy(other, f)
	}
	return IterFunc[C](func() optional.Optional[C] {
		lvalue := iter.Next()
		rvalue := other.Next()
		if lvalue.IsSome() && rvalue.IsSome() {
			return optional.Some(f(lvalue.Value(), rvalue.Value()))
		}
		return optional.None[C]()
	})
}

func Zip[A, B any](iter Iter[A], other Iter[B]) Iter[tuple.Pair[A, B]] {
	return ZipBy(iter, other, tuple.PairOf[A, B])
}

type UnzipByIter[A, B, C any] interface {
	Iter[A]
	UnzipBy(unzip func(A) (B, C)) (Iter[B], Iter[C])
}

func UnzipBy[A, B, C any](iter Iter[A], unzip func(A) (B, C)) (Iter[B], Iter[C]) {
	if unzipByIter, ok := iter.(UnzipByIter[A, B, C]); ok {
		return unzipByIter.UnzipBy(unzip)
	}
	lhsSlice := make([]B, 0)
	rhsSlice := make([]C, 0)
	ForEach(iter, func(t A) {
		lhs, rhs := unzip(t)
		lhsSlice = append(lhsSlice, lhs)
		rhsSlice = append(rhsSlice, rhs)
	})
	return OfSlice(lhsSlice), OfSlice(rhsSlice)
}

type PositionLeftIter[T any] interface {
	Iter[T]
	PositionLeft(f func(T) bool) optional.Optional[int]
}

// PositionLeft searches for an element in an iterator, returning its index.
//
// PositionLeft() takes a closure that returns true or false. It applies
// this closure to each element of the iterator, and if one of them returns true,
// then PositionLeft() returns Some(index). If all of them return false, it returns None.
//
// PositionLeft() is short-circuiting; in other words, it will stop processing as soon as it finds a true.
func PositionLeft[T any](iter Iter[T], f func(T) bool) optional.Optional[int] {
	if positionIter, ok := iter.(PositionLeftIter[T]); ok {
		return positionIter.PositionLeft(f)
	}
	return FindMap(Enumerate(iter), func(t tuple.Pair[int, T]) optional.Optional[int] {
		if f(t.Second) {
			return optional.Some(t.First)
		}
		return optional.None[int]()
	})
}

type CountByIter[T any] interface {
	Iter[T]
	CountBy(predicate func(T) bool) int
}

func CountBy[T any](iter Iter[T], predicate func(T) bool) int {
	if countByIter, ok := iter.(CountByIter[T]); ok {
		return countByIter.CountBy(predicate)
	}
	return FoldLeft[T, int](iter, 0, func(i int, t T) int {
		return cond.Ternary[int](predicate(t), i+1, i)
	})
}

type CountIter[T any] interface {
	Iter[T]
	Count() int
}

func Count[T any](iter Iter[T]) int {
	if countIter, ok := iter.(CountIter[T]); ok {
		return countIter.Count()
	}
	return CountBy(iter, func(T) bool { return true })
}

func Rev[T any](iter PrevIter[T]) Iter[T] {
	return IterFunc[T](func() optional.Optional[T] { return iter.Prev() })
}

func Nth[T any](iter Iter[T], n int) optional.Optional[T] {
	if n < 0 {
		return optional.None[T]()
	}
	idx := -1
	for {
		idx++
		item := iter.Next()
		if item.IsNone() || idx == n {
			return item
		}
	}
}

type ToStringIter[T fmt.Stringer] interface {
	Iter[T]
	ToString(sep string) string
}

func ToString[T fmt.Stringer](iter Iter[T], sep string) string {
	if toStringIter, ok := iter.(ToStringIter[T]); ok {
		return toStringIter.ToString(sep)
	}
	return ToStringBy(iter, sep, T.String)
}

type ToStringByIter[T any] interface {
	Iter[T]
	ToStringBy(sep string, f func(T) string) string
}

func ToStringBy[T any](iter Iter[T], sep string, f func(T) string) string {
	if toStringByIter, ok := iter.(ToStringByIter[T]); ok {
		return toStringByIter.ToStringBy(sep, f)
	}
	first := iter.Next()
	if first.IsNone() {
		return ""
	}
	var b strings.Builder
	b.WriteString(f(first.Value()))
	b = FoldLeft(iter, b, func(a strings.Builder, t T) strings.Builder {
		a.WriteString(sep)
		a.WriteString(f(t))
		return a
	})
	return b.String()
}
