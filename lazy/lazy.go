package lazy

import (
	"sync"

	"github.com/go-board/std/delegate"
)

type Lazy[T any] struct {
	compute func() T
}

func NewLazy[T any](compute func() T) *Lazy[T]            { return &Lazy[T]{compute} }
func (self *Lazy[T]) Get() T                              { return self.compute() }
func (self *Lazy[T]) With(consumer delegate.Consumer1[T]) { consumer(self.Get()) }

type OnceLazy[T any] struct {
	ticket  *sync.Once
	compute func() T
	inner   T
}

func NewOnceLazy[T any](compute func() T) *OnceLazy[T] {
	return &OnceLazy[T]{ticket: &sync.Once{}, compute: compute}
}

func (self *OnceLazy[T]) Get() T {
	self.ticket.Do(func() { self.inner = self.compute() })
	return self.inner
}

func (self *OnceLazy[T]) With(consumer delegate.Consumer1[T]) {
	consumer(self.Get())
}
