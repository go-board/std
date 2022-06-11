package fp

import (
	"strconv"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestCurry(t *testing.T) {
	a := qt.New(t)
	a.Run("curry1", func(c *qt.C) {
		f := func(a int) int { return a + 2 }
		curryAdd1 := Curry1(f)
		c.Assert(curryAdd1(1), qt.Equals, 3)
	})
	a.Run("curry2", func(t *qt.C) {
		add2 := func(a, b int) int { return a + b }
		curryAdd2 := Curry2(add2)
		t.Assert(curryAdd2(1)(2), qt.Equals, 3)
	})
	a.Run("curry3", func(c *qt.C) {
		add3 := func(a, b, c int) int { return a + b + c }
		curryAdd3 := Curry3(add3)
		c.Assert(curryAdd3(1)(2)(3), qt.Equals, 6)
	})
	a.Run("curry4", func(c *qt.C) {
		add4 := func(a, b, c, d int) int { return a + b + c + d }
		curryAdd4 := Curry4(add4)
		c.Assert(curryAdd4(1)(2)(3)(4), qt.Equals, 10)
	})
	a.Run("curry5", func(c *qt.C) {
		add5 := func(a, b, c, d, e int) int { return a + b + c + d + e }
		curryAdd5 := Curry5(add5)
		c.Assert(curryAdd5(1)(2)(3)(4)(5), qt.Equals, 15)
	})
}

func TestUncurry(t *testing.T) {
	a := qt.New(t)
	a.Run("uncurry1", func(c *qt.C) {
		f := func(a int) int { return a + 2 }
		c.Assert(Uncurry1(f)(1), qt.Equals, 3)
	})
	a.Run("uncurry2", func(c *qt.C) {
		f := func(a int) func(int) int {
			return func(b int) int { return a + b }
		}
		c.Assert(Uncurry2(f)(1, 2), qt.Equals, 3)
	})
	a.Run("uncurry3", func(c *qt.C) {
		f := func(a int) func(int) func(int) int {
			return func(b int) func(int) int {
				return func(c int) int { return a + b + c }
			}
		}
		c.Assert(Uncurry3(f)(1, 2, 3), qt.Equals, 6)
	})
	a.Run("uncurry4", func(c *qt.C) {
		f := func(a int) func(int) func(int) func(int) int {
			return func(b int) func(int) func(int) int {
				return func(c int) func(int) int {
					return func(d int) int { return a + b + c + d }
				}
			}
		}
		c.Assert(Uncurry4(f)(1, 2, 3, 4), qt.Equals, 10)
	})
	a.Run("uncurry5", func(c *qt.C) {
		f := func(a int) func(int) func(int) func(int) func(int) int {
			return func(b int) func(int) func(int) func(int) int {
				return func(c int) func(int) func(int) int {
					return func(d int) func(int) int {
						return func(e int) int { return a + b + c + d + e }
					}
				}
			}
		}
		c.Assert(Uncurry5(f)(1, 2, 3, 4, 5), qt.Equals, 15)
	})
}

func TestApply(t *testing.T) {
	a := qt.New(t)
	a.Run("apply1", func(c *qt.C) {
		f := func(a int) int { return a + 2 }
		c.Assert(Apply1(f, 1), qt.Equals, 3)
	})
	a.Run("apply2", func(c *qt.C) {
		f := func(a, b int) int { return a + b }
		c.Assert(Apply2(f, 1, 2), qt.Equals, 3)
	})
	a.Run("apply3", func(c *qt.C) {
		f := func(a, b, c int) int { return a + b + c }
		c.Assert(Apply3(f, 1, 2, 3), qt.Equals, 6)
	})
	a.Run("apply4", func(c *qt.C) {
		f := func(a, b, c, d int) int { return a + b + c + d }
		c.Assert(Apply4(f, 1, 2, 3, 4), qt.Equals, 10)
	})
	a.Run("apply5", func(c *qt.C) {
		f := func(a, b, c, d, e int) int { return a + b + c + d + e }
		c.Assert(Apply5(f, 1, 2, 3, 4, 5), qt.Equals, 15)
	})
}

func TestCompose(t *testing.T) {
	a := qt.New(t)
	a.Run("compose", func(c *qt.C) {
		f := func(a int) int { return a + 1 }
		c.Assert(Compose1(f)(1), qt.Equals, 2)
	})
	a.Run("compose2", func(c *qt.C) {
		f := func(a int) int { return a + 1 }
		g := func(a int) int { return a + 2 }
		c.Assert(Compose2(f, g)(1), qt.Equals, 4)
	})

	a.Run("compose3", func(c *qt.C) {
		f := func(a int) int { return a + 1 }
		g := func(a int) int { return a + 2 }
		h := func(a int) int { return a + 3 }
		c.Assert(Compose3(f, g, h)(1), qt.Equals, 7)
	})

	a.Run("compose4", func(c *qt.C) {
		f := func(a int) int { return a + 1 }
		g := func(a int) int { return a + 2 }
		h := func(a int) int { return a + 3 }
		i := func(a int) int { return a + 4 }
		c.Assert(Compose4(f, g, h, i)(1), qt.Equals, 11)
	})

	a.Run("compose5", func(c *qt.C) {
		f := func(a int) int { return a + 1 }
		g := func(a int) int { return a + 2 }
		h := func(a int) int { return a + 3 }
		i := func(a int) int { return a + 4 }
		j := func(a int) int { return a + 5 }
		c.Assert(Compose5(f, g, h, i, j)(1), qt.Equals, 16)
	})

	a.Run("multiple type compose", func(c *qt.C) {
		f := func(a int) int32 { return int32(a) + 2 }
		g := func(a int32) int64 { return int64(a) * 3 }
		h := func(a int64) int { return int(a) + 4 }
		i := func(a int) int { return a * 100 }
		j := func(a int) string { return strconv.FormatInt(int64(a), 10) }
		c.Assert(Compose5(f, g, h, i, j)(1), qt.Equals, "1300")
	})
}

func TestState(t *testing.T) {
	a := qt.New(t)
	a.Run("state", func(c *qt.C) {
		getter, setter := UseState(1)
		c.Assert(getter(), qt.Equals, 1)
		setter(2)
		c.Assert(getter(), qt.Equals, 2)
	})
	a.Run("funcState", func(c *qt.C) {
		getter, setter := UseFuncState(1)
		c.Assert(getter(), qt.Equals, 1)
		setter(func(prevState int) int { return prevState + 1 })
		c.Assert(getter(), qt.Equals, 2)
	})
}

func TestPartial(t *testing.T) {
	a := qt.New(t)
	a.Run("partial1", func(c *qt.C) {
		m := Fn2[func(int) int, []int, []int](map_[int, int]).Partial1(func(i int) int { return i * 2 })
		n := Fn2[func(int) string, []int, []string](map_[int, string]).Partial1(strconv.Itoa)
		x := Compose2(m.(Fn1[[]int, []int]), n.(Fn1[[]int, []string]))
		y := ComposeFunction2(m, n)
		c.Assert(x([]int{1, 2, 3, 4, 5}), qt.DeepEquals, []string{"2", "4", "6", "8", "10"})
		c.Assert(y.Apply([]int{1, 2, 3, 4, 5}), qt.DeepEquals, []string{"2", "4", "6", "8", "10"})
	})
}

func map_[T, U any](f func(T) U, slice []T) []U {
	res := make([]U, 0, len(slice))
	for _, s := range slice {
		res = append(res, f(s))
	}
	return res
}
