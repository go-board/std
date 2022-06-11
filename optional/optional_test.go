package optional_test

import (
	"github.com/frankban/quicktest"
	"github.com/go-board/std/optional"
	"testing"
)

func TestCtor(t *testing.T) {
	a := quicktest.New(t)

	a.Run("some", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.IsSome(), quicktest.IsTrue)
	})
	a.Run("none", func(c *quicktest.C) {
		x := optional.None[int]()
		c.Assert(x.IsNone(), quicktest.IsTrue)
	})

	a.Run("some_and", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.IsSomeAnd(func(i int) bool { return i == 100 }), quicktest.IsTrue)
		y := optional.None[int]()
		c.Assert(y.IsSomeAnd(func(i int) bool { return i == 100 }), quicktest.IsFalse)
		z := optional.Some(200)
		c.Assert(z.IsSomeAnd(func(i int) bool { return i == 100 }), quicktest.IsFalse)
	})
}

func TestValue(t *testing.T) {
	a := quicktest.New(t)

	a.Run("value", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.Value(), quicktest.Equals, 100)
		y := optional.None[int]()
		c.Assert(func() { y.Value() }, quicktest.PanicMatches, "Unwrap empty value")
	})

	a.Run("value_or", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.ValueOr(200), quicktest.Equals, 100)
		y := optional.None[int]()
		c.Assert(y.ValueOr(200), quicktest.Equals, 200)
	})

	a.Run("value_or_else", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.ValueOrElse(func() int { return 200 }), quicktest.Equals, 100)
		y := optional.None[int]()
		c.Assert(y.ValueOrElse(func() int { return 200 }), quicktest.Equals, 200)
	})

	a.Run("value_or_zero", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.ValueOrZero(), quicktest.Equals, 100)
		y := optional.None[int]()
		c.Assert(y.ValueOrZero(), quicktest.Equals, 0)
	})
}

func TestBitOp(t *testing.T) {
	a := quicktest.New(t)

	a.Run("and", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.And(optional.Some(200)).Value(), quicktest.Equals, 200)
		y := optional.None[int]()
		c.Assert(y.And(optional.Some(200)).IsNone(), quicktest.IsTrue)
		z := optional.Some(200)
		c.Assert(x.And(z).Value(), quicktest.Equals, 200)
	})
	a.Run("and_then", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.AndThen(func(i int) optional.Optional[int] { return optional.Some(i * 2) }).Value(), quicktest.Equals, 200)
		y := optional.None[int]()
		c.Assert(y.AndThen(func(i int) optional.Optional[int] { return optional.Some(i * 2) }).IsNone(), quicktest.IsTrue)
	})
	a.Run("or", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.Or(optional.Some(200)).Value(), quicktest.Equals, 100)
		y := optional.None[int]()
		c.Assert(y.Or(optional.Some(200)).Value(), quicktest.Equals, 200)
		z := optional.Some(200)
		c.Assert(x.Or(z).Value(), quicktest.Equals, 100)
	})
	a.Run("or_else", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.OrElse(func() optional.Optional[int] { return optional.Some(200) }).Value(), quicktest.Equals, 100)
		y := optional.None[int]()
		c.Assert(y.OrElse(func() optional.Optional[int] { return optional.Some(200) }).Value(), quicktest.Equals, 200)
	})

	a.Run("xor", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.Xor(optional.Some(200)).IsNone(), quicktest.IsTrue)
		y := optional.None[int]()
		c.Assert(y.Xor(optional.Some(200)).IsSome(), quicktest.IsTrue)
		z := optional.Some(200)
		c.Assert(x.Xor(z).IsSome(), quicktest.IsFalse)
	})
}

func TestMap(t *testing.T) {
	a := quicktest.New(t)
	a.Run("map", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.Map(func(i int) int { return i * 2 }).Value(), quicktest.Equals, 200)
		y := optional.None[int]()
		c.Assert(y.Map(func(i int) int { return i * 2 }).IsNone(), quicktest.IsTrue)
	})
	a.Run("map_or", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.MapOr(300, func(i int) int { return i * 2 }), quicktest.Equals, 200)
		y := optional.None[int]()
		c.Assert(y.MapOr(300, func(i int) int { return i * 2 }), quicktest.Equals, 300)
	})
	a.Run("map_or_else", func(c *quicktest.C) {
		x := optional.Some(100)
		c.Assert(x.MapOrElse(func() int { return 300 }, func(i int) int { return i * 2 }), quicktest.Equals, 200)
		y := optional.None[int]()
		c.Assert(y.MapOrElse(func() int { return 300 }, func(i int) int { return i * 2 }), quicktest.Equals, 300)
	})
}
