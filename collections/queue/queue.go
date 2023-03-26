package queue

import (
	"github.com/go-board/std/optional"
)

type Queue[TElement any] interface {
	Size() uint
	Capacity() uint

	PushFront(element TElement)
	PushBack(element TElement)

	PopFront() optional.Optional[TElement]
	PopBack() optional.Optional[TElement]

	PeekFront() optional.Optional[TElement]
	PeekBack() optional.Optional[TElement]
}
