package queue

import (
	"github.com/go-board/std/optional"
	"github.com/go-board/std/slices"
)

type ArrayQueue[T any] struct {
	inner []T
}

func (q *ArrayQueue[T]) Size() uint {
	return uint(len(q.inner))
}

func (q *ArrayQueue[T]) Capacity() uint {
	return uint(cap(q.inner))
}

func (q *ArrayQueue[T]) PushFront(element T) {
	q.inner = append([]T{element}, q.inner...)
}

func (q *ArrayQueue[T]) PushBack(element T) {
	q.inner = append(q.inner, element)
}

func (q *ArrayQueue[T]) PopFront() optional.Optional[T] {
	elem, slice := slices.SpliceFirst(q.inner)
	q.inner = slice
	return elem
}

func (q *ArrayQueue[T]) PopBack() optional.Optional[T] {
	elem, slice := slices.SpliceLast(q.inner)
	q.inner = slice
	return elem
}

func (q *ArrayQueue[T]) PeekFront() optional.Optional[T] {
	return slices.First(q.inner)
}

func (q *ArrayQueue[T]) PeekBack() optional.Optional[T] {
	return slices.Last(q.inner)
}

var _ Queue[any] = (*ArrayQueue[any])(nil)
