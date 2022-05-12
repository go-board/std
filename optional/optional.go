package optional

import (
	"fmt"

	"github.com/go-board/std/delegate"
)

// Optional is a value that may or may not be present.
type Optional[T any] struct{ data *T }

func (self Optional[T]) String() string {
	if self.IsSome() {
		return fmt.Sprintf("Some(%+v)", *self.data)
	}
	return "None"
}

// From returns an Optional from a value.
func From[T any](data T, ok bool) Optional[T] {
	if ok {
		return Some(data)
	}
	return None[T]()
}

// Of returns an Optional from a value.
func Of[T any](data *T) Optional[T] { return Optional[T]{data: data} }

// Some returns an Optional from a value.
func Some[T any](data T) Optional[T] { return Optional[T]{data: &data} }

// None returns an Optional from a value.
func None[T any]() Optional[T] { return Optional[T]{} }

// IsSome returns true if the Optional is Some.
func (self Optional[T]) IsSome() bool { return self.data != nil }

// IsNone returns true if the Optional is None.
func (self Optional[T]) IsNone() bool { return !self.IsSome() }

// Value returns the value of the Optional.
func (self Optional[T]) Value() T {
	if self.IsSome() {
		return *self.data
	}
	panic("Unwrap empty value")
}

func (self Optional[T]) And(opt Optional[T]) Optional[T] {
	if self.IsSome() {
		return opt
	}
	return None[T]()
}

func (self Optional[T]) Or(opt Optional[T]) Optional[T] {
	if self.IsNone() {
		return opt
	}
	return self
}

func (self Optional[T]) OrElse(defaultValue T) T {
	if self.IsSome() {
		return self.Value()
	}
	return defaultValue
}

func (self Optional[T]) IfPresent(consume delegate.Consumer1[T]) {
	if self.IsSome() {
		consume(self.Value())
	}
}

func (self Optional[T]) Filter(fn delegate.Predicate[T]) Optional[T] {
	if self.IsSome() && fn(self.Value()) {
		return Some(self.Value())
	}
	return None[T]()
}

func (self Optional[T]) Map(mapFn delegate.Function1[T, T]) Optional[T] {
	return Map(self, mapFn)
}

func (self Optional[T]) OrZero() T {
	if self.IsSome() {
		return self.Value()
	}
	var zero T
	return zero
}
