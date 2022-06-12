package lazy

import (
	"sync"
)

// Lazy is a lazy value.
type Lazy[T any] struct {
	compute func() T
}

func NewLazy[T any](compute func() T) *Lazy[T] { return &Lazy[T]{compute} }
func (self *Lazy[T]) Get() T                   { return self.compute() }
func (self *Lazy[T]) With(consumer func(T))    { consumer(self.Get()) }

// OnceLazy is a lazy value that is computed only once.
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

func (self *OnceLazy[T]) With(consumer func(T)) {
	consumer(self.Get())
}
