package lazy_test

import (
	"testing"

	"github.com/frankban/quicktest"

	"github.com/go-board/std/lazy"
)

func TestLazyCell(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		lazyValue := lazy.NewCell(func() int { return 100 })
		if lazyValue.Get() != 100 {
			t.Errorf("lazy value should be 100, but got %d", lazyValue.Get())
		}
	})
}

func TestOnceFunc(t *testing.T) {
	x := 0
	f := lazy.OnceFunc(func() { x++ })
	f()
	quicktest.Assert(t, x, quicktest.Equals, 1)
	f()
	quicktest.Assert(t, x, quicktest.Equals, 1)
}

func TestOnceValue(t *testing.T) {
	i := 0
	x := lazy.OnceValue(func() int {
		i++
		return i
	})
	quicktest.Assert(t, x(), quicktest.Equals, 1)
	quicktest.Assert(t, x(), quicktest.Equals, 1)
}

func TestOnceValues(t *testing.T) {
	i := 0
	ok := false
	f := lazy.OnceValues(func() (int, bool) {
		i++
		ok = !ok
		return i, ok
	})
	a, b := f()
	quicktest.Assert(t, a, quicktest.Equals, 1)
	quicktest.Assert(t, b, quicktest.IsTrue)
	c, d := f()
	quicktest.Assert(t, c, quicktest.Equals, 1)
	quicktest.Assert(t, d, quicktest.IsTrue)
}

func TestTernary(t *testing.T) {
	x := lazy.Ternary(true, func() int {
		return 1
	}, func() int {
		return 2
	})
	quicktest.Assert(t, x, quicktest.Equals, 1)
}
