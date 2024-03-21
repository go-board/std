package fp

import (
	"strconv"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestCurry(t *testing.T) {
	a := qt.New(t)
	a.Run("curry1", func(c *qt.C) {
		f := MakeFn1(func(i int) string { return strconv.Itoa(i) })
		c.Assert(f.Curry()(3), qt.Equals, "3")
	})
	a.Run("curry2", func(c *qt.C) {
		f := MakeFn2(func(a int, b int) string { return strconv.Itoa(a + b) })
		c.Assert(f.Curry()(1)(2), qt.Equals, "3")
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
		x := Compose(m, n)
		c.Assert(x([]int{1, 2, 3, 4, 5}), qt.DeepEquals, []string{"2", "4", "6", "8", "10"})
	})
}

func map_[T, U any](f func(T) U, slice []T) []U {
	res := make([]U, 0, len(slice))
	for _, s := range slice {
		res = append(res, f(s))
	}
	return res
}
