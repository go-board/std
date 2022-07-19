package ops

import (
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/iterator"
)

func All[T any](iter iterator.Iterator[T], predicate func(T) bool) bool {
	for s := iter.Next(); s.IsSome(); s = iter.Next() {
		if !predicate(s.Value()) {
			return false
		}
	}
	return true
}

func Any[T any](iter iterator.Iterator[T], predicate func(T) bool) bool {
	for s := iter.Next(); s.IsSome(); s = iter.Next() {
		if predicate(s.Value()) {
			return true
		}
	}
	return false
}

func Once[T any](iter iterator.Iterator[T], predicate func(T) bool) bool {
	return matchCount(iter, predicate, 1)
}

func matchCount[T any](iter iterator.Iterator[T], predicate func(T) bool, n uint) bool {
	hitCnt := uint(0)
	for s := iter.Next(); s.IsSome(); s = iter.Next() {
		if predicate(s.Value()) {
			hitCnt++
		}
	}
	return hitCnt == n
}

func None[T any](iter iterator.Iterator[T], predicate func(T) bool) bool {
	return All(iter, func(t T) bool { return !predicate(t) })
}

func ContainsBy[T any](iter iterator.Iterator[T], target T, eq cmp.EqFunc[T]) bool {
	return Any(iter, func(t T) bool { return eq(t, target) })
}

func Contains[T comparable](iter iterator.Iterator[T], target T) bool {
	return ContainsBy(iter, target, func(t1 T, t2 T) bool { return t1 == t2 })
}
