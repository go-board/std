package iter

import (
	"github.com/go-board/std/optional"
)

// OfChan create a new Iter[T] from chan T
func OfChan[T any](ch <-chan T) Iter[T] {
	return IterFunc[T](func() optional.Optional[T] {
		v, ok := <-ch
		return optional.FromPair[T](v, ok)
	})
}

// OfSliceVariadic create a new Iter[T] from variadic sequence of T
func OfSliceVariadic[T any](elems ...T) PrevIter[T] {
	return &sliceIter[T]{slice: elems, nextIndex: -1, prevIndex: len(elems)}
}

// OfSlice create a new Iter[T] from sequence of T.
func OfSlice[T any](slice []T) PrevIter[T] {
	return &sliceIter[T]{slice: slice, nextIndex: -1, prevIndex: len(slice)}
}

type sliceIter[T any] struct {
	slice     []T
	nextIndex int
	prevIndex int
}

func (i *sliceIter[T]) Next() optional.Optional[T] {
	i.nextIndex++
	if i.nextIndex < len(i.slice) {
		return optional.Some(i.slice[i.nextIndex])
	}
	return optional.None[T]()
}

func (i *sliceIter[T]) Prev() optional.Optional[T] {
	i.prevIndex--
	if i.prevIndex >= 0 {
		return optional.Some(i.slice[i.prevIndex])
	}
	return optional.None[T]()
}

// Empty create an Iter always return None.
func Empty[T any]() Iter[T] {
	return IterFunc[T](func() optional.Optional[T] { return optional.None[T]() })
}

// Repeat create an Iter always return the same value.
func Repeat[T any](v T) Iter[T] {
	return IterFunc[T](func() optional.Optional[T] { return optional.Some(v) })
}

func Successor[T any](init T, successor func(T) optional.Optional[T]) Iter[T] {
	value := optional.Some[T](init)
	return IterFunc[T](func() optional.Optional[T] {
		if value.IsNone() {
			return optional.None[T]()
		}
		v := value
		value = successor(value.Value())
		return v
	})
}

// Counter return an endless counter [Iter] start from 0
func Counter() Iter[int] {
	current := -1
	return IterFunc[int](func() optional.Optional[int] {
		current++
		return optional.Some(current)
	})
}

// Range returns an Iter that iterate over [start, end)
// when end <= start, it panics
func Range(start int, end int) Iter[int] {
	if end <= start {
		panic("iter.SubSet: end not great than start")
	}
	current := start - 1
	return IterFunc[int](func() optional.Optional[int] {
		current++
		if current < end {
			return optional.Some(current)
		}
		return optional.None[int]()
	})
}
