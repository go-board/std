package optional

import (
	"fmt"
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

// IsSomeAnd returns true if the Optional is Some and satisfies the given predicate.
func (self Optional[T]) IsSomeAnd(predicate func(T) bool) bool {
	if self.IsSome() {
		return predicate(*self.data)
	}
	return false
}

// Value returns the value of the Optional.
func (self Optional[T]) Value() T {
	if self.IsSome() {
		return *self.data
	}
	panic("Unwrap empty value")
}

// ValueOr returns the Optional if it is Some, otherwise returns the given default value.
func (self Optional[T]) ValueOr(defaultValue T) T {
	if self.IsSome() {
		return self.Value()
	}
	return defaultValue
}

// ValueOrElse returns the Optional if it is Some, otherwise returns the given default value.
func (self Optional[T]) ValueOrElse(defaultFunc func() T) T {
	if self.IsSome() {
		return self.Value()
	}
	return defaultFunc()
}

// ValueOrZero returns the Optional if it is Some, otherwise returns the zero value of the type.
func (self Optional[T]) ValueOrZero() T {
	var zero T
	return self.ValueOr(zero)
}

// And returns None if the option is None, otherwise returns given opt.
func (self Optional[T]) And(opt Optional[T]) Optional[T] {
	if self.IsSome() {
		return opt
	}
	return None[T]()
}

// AndThen returns None if the option is None, otherwise calls f with the wrapped value and returns the result.
func (self Optional[T]) AndThen(f func(T) Optional[T]) Optional[T] {
	if self.IsSome() {
		return f(*self.data)
	}
	return None[T]()
}

// Or returns the given opt if the option is None, otherwise returns the option.
func (self Optional[T]) Or(opt Optional[T]) Optional[T] {
	if self.IsNone() {
		return opt
	}
	return self
}

// OrElse returns the Optional if it contains a value, otherwise calls f and returns the result.
func (self Optional[T]) OrElse(f func() Optional[T]) Optional[T] {
	if self.IsSome() {
		return Some(*self.data)
	}
	return f()
}

func (self Optional[T]) IfPresent(consume func(T)) {
	if self.IsSome() {
		consume(self.Value())
	}
}

func (self Optional[T]) Filter(fn func(T) bool) Optional[T] {
	if self.IsSome() && fn(self.Value()) {
		return Some(self.Value())
	}
	return None[T]()
}

// Map returns None if the option is None, otherwise calls the given function and returns the result.
func (self Optional[T]) Map(f func(T) T) Optional[T] {
	return Map(self, f)
}

// MapOr returns None if the option is None, otherwise calls the given function and returns the result.
func (self Optional[T]) MapOr(defaultValue T, f func(T) T) T {
	if self.IsSome() {
		return f(self.Value())
	}
	return defaultValue
}

// MapOrElse returns None if the option is None, otherwise calls the given function and returns the result.
func (self Optional[T]) MapOrElse(df func() T, f func(T) T) T {
	if self.IsSome() {
		return f(self.Value())
	}
	return df()
}

// Xor returns None if the option is None, otherwise returns the given opt.
func (self Optional[T]) Xor(opt Optional[T]) Optional[T] {
	if self.IsSome() && opt.IsNone() {
		return Some(*self.data)
	}
	if self.IsNone() && opt.IsSome() {
		return Some(*opt.data)
	}
	return None[T]()
}
