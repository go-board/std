package queue

import (
	"github.com/go-board/std/iter"
	"github.com/go-board/std/optional"
)

type ArrayQueue[E any] struct{ inner []E }

func (q *ArrayQueue[E]) Iter() iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, x := range q.inner {
			if !yield(x) {
				break
			}
		}
	}
}

func (q *ArrayQueue[E]) Size() int {
	return len(q.inner)
}

func (q *ArrayQueue[E]) PushFront(element E) {
	q.inner = append([]E{element}, q.inner...)
}

func (q *ArrayQueue[E]) PushBack(element E) {
	q.inner = append(q.inner, element)
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
