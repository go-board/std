//go:build goexpriments.rangefunc

package iter_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-board/std/iter"
	"github.com/go-board/std/tuple"
)

func TestCmp(t *testing.T) {
	t.Run("gt", func(t *testing.T) {
		x := seq(1, 2)
		y := seq(1)
		ok := iter.Gt(x, y)
		qt.Assert(t, ok, qt.IsTrue)
	})
	t.Run("ge", func(t *testing.T) {
		x := seq(1, 2)
		y := seq(1)
		ok := iter.Ge(x, y)
		qt.Assert(t, ok, qt.IsTrue)
	})
	t.Run("lt", func(t *testing.T) {
		x := seq(1, 2)
		y := seq(1, 2, 3)
		ok := iter.Lt(x, y)
		qt.Assert(t, ok, qt.IsTrue)
	})
	t.Run("le", func(t *testing.T) {
		x := seq(1, 2)
		y := seq(1, 2, 3)
		ok := iter.Le(x, y)
		qt.Assert(t, ok, qt.IsTrue)
	})
}

func TestEqual(t *testing.T) {
	t.Run("eq", func(t *testing.T) {
		ok := iter.Eq(seq(1), seq(1))
		qt.Assert(t, ok, qt.IsTrue)
	})
	t.Run("ne", func(t *testing.T) {
		ok := iter.Ne(seq(1), seq(2))
		qt.Assert(t, ok, qt.IsTrue)
	})
}

func TestZip(t *testing.T) {
	t.Run("zip", func(t *testing.T) {
		x := seq(1, 2, 3)
		y := seq(1, 2, 3)
		z := iter.Zip(x, y)
		qt.Assert(t, collect(z), qt.DeepEquals, []tuple.Pair[int, int]{tuple.MakePair(1, 1), tuple.MakePair(2, 2), tuple.MakePair(3, 3)})
	})
	t.Run("zip func", func(t *testing.T) {
		x := seq(1, 2, 3)
		y := seq(1, 2, 3)
		z := iter.ZipFunc(x, y, func(a, b int) int { return a + b })
		qt.Assert(t, collect(z), qt.DeepEquals, []int{2, 4, 6})
	})
}

func TestPullOption(t *testing.T) {
	it, stop := iter.PullOption(seq(1, 2))
	defer stop()
	first := it()
	qt.Assert(t, first.IsSomeAnd(func(i int) bool { return i == 1 }), qt.IsTrue)
	second := it()
	qt.Assert(t, second.IsSomeAnd(func(i int) bool { return i == 2 }), qt.IsTrue)
	third := it()
	qt.Assert(t, third.IsNone(), qt.IsTrue)
}

func TestRangeSeq(t *testing.T) {
	for x := range seq(1, 2, 3) {
		t.Logf("%+v\n", x)
	}
}
