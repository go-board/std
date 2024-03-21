package fp_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-board/std/fp"
)

func TestEither_IsLeft(t *testing.T) {
	l := fp.Left[int, string](1)
	qt.Assert(t, l.IsLeft(), qt.IsTrue)
	r := fp.Right[int, string]("")
	qt.Assert(t, r.IsLeft(), qt.IsFalse)
}

func TestEither_IsRight(t *testing.T) {
	l := fp.Left[int, string](1)
	qt.Assert(t, l.IsRight(), qt.IsFalse)
	r := fp.Right[int, string]("")
	qt.Assert(t, r.IsRight(), qt.IsTrue)
}

func TestEither_Left(t *testing.T) {
	a := qt.New(t)
	a.Run("left", func(c *qt.C) {
		l := fp.Left[int, string](1)
		c.Assert(l.Left(), qt.Equals, 1)
	})
	a.Run("right", func(c *qt.C) {
		defer func() {
			x := recover()
			a.Assert(x, qt.IsNotNil)
		}()
		r := fp.Right[int, string]("")
		c.Assert(r.Left(), qt.Equals, 0)
	})
}

func TestEither_Right(t *testing.T) {
	a := qt.New(t)
	a.Run("left", func(c *qt.C) {
		defer func() {
			x := recover()
			a.Assert(x, qt.IsNotNil)
		}()
		l := fp.Left[int, string](1)
		c.Assert(l.Right(), qt.Equals, "")
	})
	a.Run("right", func(c *qt.C) {
		r := fp.Right[int, string]("")
		c.Assert(r.Right(), qt.Equals, "")
	})
}
