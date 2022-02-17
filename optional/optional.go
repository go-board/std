package optional

import (
	"github.com/go-board/std/delegate"
)

type Optional[T any] struct{ data *T }

func From[T any](data T, ok bool) Optional[T] {
	if ok {
		return Some(data)
	}
	return None[T]()
}

func Of[T any](data T) Optional[T]     { return Optional[T]{data: &data} }
func OfPtr[T any](data *T) Optional[T] { return Optional[T]{data: data} }
func Some[T any](data T) Optional[T]   { return Optional[T]{data: &data} }
func None[T any]() Optional[T]         { return Optional[T]{} }

func (self Optional[T]) IsSome() bool { return self.data != nil }

func (self Optional[T]) IsNone() bool { return !self.IsSome() }

func (self Optional[T]) Value() T {
	if self.IsSome() {
		return *self.data
	}
	panic("Unwrap empty value")
}

func (self Optional[T]) IfPresent(consume delegate.Consumer1[T]) {
	if self.IsSome() {
		consume(self.Value())
	}
}

func (self Optional[T]) IfAbsent(consume func()) {
	if self.IsNone() {
		consume()
	}
}
