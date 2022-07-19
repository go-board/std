package try_test

import (
	"errors"
	"testing"

	"github.com/frankban/quicktest"
	"github.com/go-board/std/result/try"
)

func TestTry(t *testing.T) {
	a := quicktest.New(t)
	a.Run("Error", func(c *quicktest.C) {
		defer try.CatchFunc(func(e *try.Error) {
			c.Assert(e.Unwrap().Error(), quicktest.Equals, "error")
		})
		try.Try(errors.New("error"))
	})
	a.Run("NonError", func(*quicktest.C) {
		try.Try(nil)
	})
}

func TestTryReturn(t *testing.T) {
	a := quicktest.New(t)
	a.Run("Return", func(c *quicktest.C) {
		x := try.Try1(100, nil)
		c.Assert(x, quicktest.Equals, 100)
	})
}
