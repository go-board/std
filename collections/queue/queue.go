package queue

import (
	"github.com/go-board/std/iter"
	"github.com/go-board/std/optional"
)

// ArrayQueue is a double ended queue backed by array.
type ArrayQueue[E any] struct{ inner []E }

// New create an empty Queue.
func New[E any]() *ArrayQueue[E] {
	return &ArrayQueue[E]{}
}

// FromSlice creates an [ArrayQueue] from slice.
func FromSlice[E any](elems ...E) *ArrayQueue[E] {
	return &ArrayQueue[E]{inner: elems}
}

// FromIter creates an [ArrayQueue] from [iter.Seq].
func FromIter[E any](it iter.Seq[E]) *ArrayQueue[E] {
	q := new(ArrayQueue[E])
	q.PushBackIter(it)
	return q
}

// Forward creates an [iter.Seq] in forward order.
func (q *ArrayQueue[E]) Forward() iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, x := range q.inner {
			if !yield(x) {
				break
			}
		}
	}
}

// Backward creates an [iter.Seq] in backward order.
func (q *ArrayQueue[E]) Backward() iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := len(q.inner) - 1; i >= 0; i-- {
			if !yield(q.inner[i]) {
				break
			}
		}
	}
}

// Size returns size of [ArrayQueue].
func (q *ArrayQueue[E]) Size() int { return len(q.inner) }

func (q *ArrayQueue[E]) PushFront(element E) {
	q.inner = append([]E{element}, q.inner...)
}

func (q *ArrayQueue[E]) PushFrontIter(it iter.Seq[E]) {
	iter.ForEach(it, q.PushFront)
}

func (q *ArrayQueue[E]) PushBack(element E) {
	q.inner = append(q.inner, element)
}

func (q *ArrayQueue[E]) PushBackIter(it iter.Seq[E]) {
	iter.ForEach(it, q.PushBack)
}

func (q *ArrayQueue[E]) PopFront() optional.Optional[E] {
	if len(q.inner) == 0 {
		return optional.None[E]()
	}
	front := q.inner[0]
	q.inner = q.inner[1:]
	return optional.Some(front)
}

func (q *ArrayQueue[E]) PopBack() optional.Optional[E] {
	if len(q.inner) == 0 {
		return optional.None[E]()
	}
	back := q.inner[len(q.inner)-1]
	q.inner = q.inner[:len(q.inner)-1]
	return optional.Some(back)
}

func (q *ArrayQueue[E]) PeekFront() optional.Optional[E] {
	if len(q.inner) == 0 {
		return optional.None[E]()
	}
	return optional.Some(q.inner[0])
}

func (q *ArrayQueue[E]) PeekBack() optional.Optional[E] {
	if len(q.inner) == 0 {
		return optional.None[E]()
	}
	return optional.Some(q.inner[len(q.inner)-1])
}
