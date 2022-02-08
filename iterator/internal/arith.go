package internal

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

func MaxBy[T any](iter iterator.Iterator[T], cmp delegate.Comparison[T]) optional.Optional[T] {
	return iter.Next()
}

func MinBy[T any](iter iterator.Iterator[T], cmp delegate.Comparison[T]) optional.Optional[T] {
	return iter.Next()
}

func SumBy[T any](iter iterator.Iterator[T], sum delegate.Add[T, T, T]) optional.Optional[T] {
	return iter.Next()
}

type Numbric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

func Max[T Numbric](iter iterator.Iterator[T]) optional.Optional[T] {
	return MaxBy(iter, func(t1, t2 T) int {
		if t1 > t2 {
			return 1
		} else if t1 < t2 {
			return -1
		}
		return 0
	})
}

func Min[T Numbric](iter iterator.Iterator[T]) optional.Optional[T] {
	return MinBy(iter, func(t1, t2 T) int {
		if t1 > t2 {
			return 1
		} else if t1 < t2 {
			return -1
		}
		return 0
	})
}

func Sum[T Numbric](iter iterator.Iterator[T]) optional.Optional[T] {
	return SumBy(iter, func(t1 T, t2 T) T { return t1 + t2 })
}

func Nth[T any](iter iterator.Iterator[T], n uint) optional.Optional[T] {
	return iter.Next()
}
