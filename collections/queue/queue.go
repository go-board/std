package queue

import (
	"time"

	"github.com/go-board/std/optional"
)

type Queue[TElement any] interface {
	Size() uint
	Capacity() uint

	PushFront(element TElement)
	PushBack(element TElement)

	PopFront() optional.Optional[TElement]
	PopBack() optional.Optional[TElement]

	PollFront(timeout time.Duration) optional.Optional[TElement]
	PollBack(timeout time.Duration) optional.Optional[TElement]

	PeekFront() optional.Optional[TElement]
	PeekBack() optional.Optional[TElement]
}
