package lazy_test

import (
	"testing"

	"github.com/frankban/quicktest"

	"github.com/go-board/std/lazy"
)

func TestLazyCell(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		lazyValue := lazy.NewLazyCell(func() int { return 100 })
		if lazyValue.Get() != 100 {
			t.Errorf("lazy value should be 100, but got %d", lazyValue.Get())
		}
	})
}

func TestOnceCell(t *testing.T) {
	cell := lazy.OnceCell[int]{}
	a := quicktest.New(t)
	a.Run("Get", func(c *quicktest.C) {
		val := cell.Get()
		c.Assert(val.IsNone(), quicktest.IsTrue)
	})
	a.Run("GetOrInit", func(c *quicktest.C) {
		val := cell.GetOrInit(func() int { return 1 })
		c.Assert(val, quicktest.Equals, 1)
	})
	a.Run("Get after init", func(c *quicktest.C) {
		val := cell.Get()
		c.Assert(val.IsSome(), quicktest.IsTrue)
		c.Assert(val.Value(), quicktest.Equals, 1)
	})
}
