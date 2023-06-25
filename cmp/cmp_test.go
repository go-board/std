package cmp_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/slices"
)

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
			userCmp := func(lhs, rhs user) cmp.Order {
				if lhs.Name < rhs.Name {
					return cmp.OrderLess
				} else if lhs.Name > rhs.Name {
					return cmp.OrderGreater
				} else {
					return cmp.OrderEqual
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
