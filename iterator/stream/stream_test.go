package stream_test

import (
	"testing"

	"github.com/go-board/std/iterator/adapters"
	"github.com/go-board/std/iterator/stream"
)

func TestStreamBool(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		x := stream.FromIterator(adapters.OfSlice(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
		all := x.All(func(x int) bool { return true })
		if !all {
			t.Fail()
		}
	})
	t.Run("any", func(t *testing.T) {
		x := stream.FromIterator(adapters.OfSlice(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
		all := x.Any(func(x int) bool { return x == 3 })
		if !all {
			t.Fail()
		}
	})
}
