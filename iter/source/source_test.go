package source_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-board/std/iter"
	"github.com/go-board/std/iter/source"
)

func collect[E any](s iter.Seq[E]) []E {
	rs := make([]E, 0)
	s(func(e E) bool {
		rs = append(rs, e)
		return true
	})
	return rs
}

func TestRange(t *testing.T) {
	t.Run("range 1", func(t *testing.T) {
		qt.Assert(t, collect(source.Range1(3)), qt.DeepEquals, []int{0, 1, 2})
		qt.Assert(t, collect(source.Range1(0)), qt.DeepEquals, []int{})
	})
	t.Run("range 2", func(t *testing.T) {
		qt.Assert(t, collect(source.Range2(1, 3)), qt.DeepEquals, []int{1, 2})
		qt.Assert(t, collect(source.Range2(3, 1)), qt.DeepEquals, []int{})
		qt.Assert(t, collect(source.Range2(3, 3)), qt.DeepEquals, []int{})
	})
	t.Run("range 3", func(t *testing.T) {
		qt.Assert(t, collect(source.Range3(1, 3, 1)), qt.DeepEquals, []int{1, 2})
		qt.Assert(t, collect(source.Range3(0, 10, 3)), qt.DeepEquals, []int{0, 3, 6, 9})
		qt.Assert(t, collect(source.Range3(0, 5, 0)), qt.DeepEquals, []int{})
	})
}

func TestRepeatTimes(t *testing.T) {
	qt.Assert(t, collect(source.RepeatTimes(2, 5)), qt.DeepEquals, []int{2, 2, 2, 2, 2})
}

func TestEmpty(t *testing.T) {
	qt.Assert(t, collect(source.Empty[int]()), qt.DeepEquals, []int{})
}

func TestOnce(t *testing.T) {
	qt.Assert(t, collect(source.Once(2)), qt.DeepEquals, []int{2})
}

func TestVariadic(t *testing.T) {
	qt.Assert(t, collect(source.Variadic(1, 2, 3)), qt.DeepEquals, []int{1, 2, 3})
}
