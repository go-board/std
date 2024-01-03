package queue_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-board/std/collections/queue"
)

func TestArrayQueue_Size(t *testing.T) {
	q := queue.FromSlice(1, 2)
	qt.Assert(t, q.Size(), qt.Equals, 2)
}
