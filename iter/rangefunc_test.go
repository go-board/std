//go:build goexperiment.rangefunc

package iter_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-board/std/iter"
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
