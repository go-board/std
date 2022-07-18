package ioc_test

import (
	"testing"

	"github.com/frankban/quicktest"
	"github.com/go-board/std/ioc"
)

type Service struct{ data int }

type ScopedService struct{ data int }

type InvalidalService struct{ data int }

func TestIoc(t *testing.T) {
	f := ioc.Configure(func(c *ioc.Factory) {
		ioc.AddSingleton(c, &Service{data: 100})
		ioc.AddPrototype(c, func(c *ioc.Factory) (*ScopedService, error) {
			s := ioc.Get[*Service](c)
			if s.IsErr() {
				return nil, s.Error()
			}
			return &ScopedService{data: s.Value().data * 10}, nil
		})
	})
	a := quicktest.New(t)
	a.Run("Singleton", func(c *quicktest.C) {
		service := ioc.Get[*Service](f)
		c.Assert(service.IsErr(), quicktest.IsFalse)
		c.Assert(service.Value().data, quicktest.Equals, 100)
		anotherService := ioc.Get[*Service](f)
		c.Assert(service.Value(), quicktest.Equals, anotherService.Value())
	})
	a.Run("Prototype", func(c *quicktest.C) {
		service := ioc.Get[*ScopedService](f)
		c.Assert(service.IsErr(), quicktest.IsFalse)
		c.Assert(service.Value().data, quicktest.Equals, 1000)
		anotherService := ioc.Get[*ScopedService](f)
		c.Assert(service.Value(), quicktest.Not(quicktest.Equals), anotherService.Value())
	})
	a.Run("Invalidal", func(c *quicktest.C) {
		service := ioc.Get[*InvalidalService](f)
		c.Assert(service.IsErr(), quicktest.IsTrue)
		c.Logf("invalidal service: %s", service.Error())
	})
	a.Run("InvalidTyped", func(c *quicktest.C) {
		service := ioc.GetNamed[*ScopedService](f, "*ioc_test.Service")
		c.Assert(service.IsErr(), quicktest.IsTrue)
		c.Logf("invalidal service: %s", service.Error())
	})
	a.Run("Must", func(c *quicktest.C) {
		service := ioc.MustGet[*Service](f)
		c.Assert(service.data, quicktest.Equals, 100)
		anotherService := ioc.MustGet[*ScopedService](f)
		c.Assert(anotherService.data, quicktest.Equals, 1000)
	})
}
