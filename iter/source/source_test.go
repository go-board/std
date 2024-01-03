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

func TestGen(t *testing.T) {
	qt.Assert(t, collect(source.Gen(2, 5)), qt.DeepEquals, []int{2, 3, 4})
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
