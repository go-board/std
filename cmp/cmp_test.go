package cmp_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/slices"
)

func TestOrdering(t *testing.T) {
	c := qt.New(t)
	c.Run("cmp", func(c *qt.C) {
		c.Assert(cmp.Order(cmp.Less, cmp.Less), qt.Equals, cmp.Equal)
		c.Assert(cmp.Order(cmp.Less, cmp.Equal), qt.Equals, cmp.Less)
		c.Assert(cmp.Order(cmp.Less, cmp.Greater), qt.Equals, cmp.Less)
		c.Assert(cmp.Order(cmp.Equal, cmp.Less), qt.Equals, cmp.Greater)
		c.Assert(cmp.Order(cmp.Equal, cmp.Equal), qt.Equals, cmp.Equal)
		c.Assert(cmp.Order(cmp.Equal, cmp.Greater), qt.Equals, cmp.Less)
		c.Assert(cmp.Order(cmp.Greater, cmp.Less), qt.Equals, cmp.Greater)
		c.Assert(cmp.Order(cmp.Greater, cmp.Equal), qt.Equals, cmp.Greater)
		c.Assert(cmp.Order(cmp.Greater, cmp.Greater), qt.Equals, cmp.Equal)
	})

	c.Run("string", func(c *qt.C) {
		c.Assert(cmp.Less.String(), qt.Equals, "Less")
		c.Assert(cmp.Equal.String(), qt.Equals, "Equal")
		c.Assert(cmp.Greater.String(), qt.Equals, "Greater")
	})
}

type user struct {
	Name string
	Age  int
	Tags []string
}

func (u user) Eq(rhs user) bool {
	return u.Name == rhs.Name && u.Age == rhs.Age && slices.Equal(u.Tags, rhs.Tags)
}

func (u user) Ne(rhs user) bool {
	return !u.Eq(rhs)
}

func TestCmp(t *testing.T) {
	c := qt.New(t)
	c.Run("max", func(c *qt.C) {
		c.Run("ordered", func(c *qt.C) {
			c.Assert(cmp.MaxOrdered(1, 2), qt.Equals, 2)
			c.Assert(cmp.MaxOrdered(2, 1), qt.Equals, 2)
			c.Assert(cmp.MaxOrdered(2, 2), qt.Equals, 2)
		})
		c.Run("by", func(c *qt.C) {
			userCmp := func(lhs, rhs user) cmp.Ordering {
				if lhs.Name < rhs.Name {
					return cmp.Less
				} else if lhs.Name > rhs.Name {
					return cmp.Greater
				} else {
					return cmp.Equal
				}
			}
			lhs := user{Name: "a", Age: 1, Tags: []string{"a"}}
			rhs := user{Name: "b", Age: 2, Tags: []string{"b"}}
			c.Assert(cmp.MaxBy(userCmp, lhs, rhs), qt.DeepEquals, rhs)
			c.Assert(cmp.MaxBy(userCmp, rhs, lhs), qt.DeepEquals, rhs)
		})
	})
}

func TestEq(t *testing.T) {
	c := qt.New(t)
	c.Run("user", func(c *qt.C) {
		lhs := user{Name: "a", Age: 1, Tags: []string{"a"}}
		rhs := user{Name: "a", Age: 1, Tags: []string{"a"}}
		c.Assert(lhs.Eq(rhs), qt.IsTrue)
		c.Assert(lhs.Ne(rhs), qt.IsFalse)
	})
}
