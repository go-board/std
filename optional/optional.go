package optional

import (
	"encoding/json"
	"fmt"
)

// Optional is a value that may or may not be present.
type Optional[T any] struct{ data *T }

// FromPair returns an Optional from a value.
func FromPair[T any](data T, ok bool) Optional[T] {
	if ok {
		return Some(data)
	}
	return None[T]()
}

// FromPtr returns an Optional from a value.
func FromPtr[T any, P ~*T](data P) Optional[T] { return Optional[T]{data: data} }

// Some returns an Optional from a value.
//
// Oops!! We can't restrict T is not pointer type. Holy shit!!!
func Some[T any](data T) Optional[T] { return Optional[T]{data: &data} }

// None returns an Optional from a value.
func None[T any]() Optional[T] { return Optional[T]{} }

func (self Optional[T]) Compare(other Optional[T], cmp func(T, T) int) int {
	if self.IsSome() && other.IsSome() {
		return cmp(self.Value(), other.Value())
	} else if self.IsSome() {
		return +1
	} else if other.IsSome() {
		return -1
	}
	return 0
}

func (self Optional[T]) String() string {
	if self.IsSome() {
		return fmt.Sprintf("Some(%+v)", *self.data)
	}
	return "None"
}

// IsSome returns true if the Optional is Some.
func (self Optional[T]) IsSome() bool { return self.data != nil }

// IsNone returns true if the Optional is None.
func (self Optional[T]) IsNone() bool { return !self.IsSome() }

// IsSomeAnd returns true if the Optional is Some and satisfies the given predicate.
func (self Optional[T]) IsSomeAnd(f func(T) bool) bool {
	if self.IsSome() {
		return f(*self.data)
	}
	return false
}

func (self Optional[T]) Expect(msg string) T {
	if self.IsSome() {
		return *self.data
	}
	panic(msg)
}

// Value returns the value of the Optional.
func (self Optional[T]) Value() T {
	if self.IsSome() {
		return *self.data
	}
	panic("Unwrap empty value")
}

// ValueOr returns the Optional if it is Some, otherwise returns the given default value.
func (self Optional[T]) ValueOr(v T) T {
	if self.IsSome() {
		return self.Value()
	}
	return v
}

// ValueOrElse returns the Optional if it is Some, otherwise returns the given default value.
func (self Optional[T]) ValueOrElse(f func() T) T {
	if self.IsSome() {
		return self.Value()
	}
	return f()
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

func (self Optional[T]) IfPresent(f func(T)) {
	if self.IsSome() {
		f(self.Value())
	}
}

func (self Optional[T]) Filter(f func(T) bool) Optional[T] {
	if self.IsSome() && f(self.Value()) {
		return Some(self.Value())
	}
	return None[T]()
}

// Map returns None if the option is None, otherwise calls the given function and returns the result.
func (self Optional[T]) Map(f func(T) T) Optional[T] {
	return Map(self, f)
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

func (self Optional[T]) EqBy(opt Optional[T], eq func(T, T) bool) bool {
	if self.IsSome() && opt.IsSome() {
		return eq(self.Value(), opt.Value())
	}
	if self.IsNone() && opt.IsNone() {
		return true
	}
	return false
}

func (self Optional[T]) CloneBy(clone func(T) T) Optional[T] {
	if self.IsNone() {
		return None[T]()
	}
	return Some(clone(self.Value()))
}

func (self Optional[T]) Get() (data T, ok bool) {
	if self.IsSome() {
		data, ok = self.Value(), true
	}
	return
}

func (self Optional[T]) MarshalJSON() ([]byte, error) {
	if self.IsSome() {
		return json.Marshal(*self.data)
	}
	return []byte("null"), nil
}

var _ json.Marshaler = (*Optional[any])(nil)

func (self Optional[T]) UnmarshalJSON(v []byte) error {
	if string(v) != "null" {
		return json.Unmarshal(v, &self.data)
	}
	return nil
}

var _ json.Unmarshaler = (*Optional[any])(nil)

func Map[A, B any](opt Optional[A], f func(A) B) Optional[B] {
	if opt.IsSome() {
		return Some(f(opt.Value()))
	}
	return None[B]()
}

// As try to convert the value to the given type.
//
// if the value is not convertible to the given type,
// then the function returns None,
// otherwise, it returns Some.
//
// Deprecated: use [TypeAssert] instead.
func As[T any](v any) Optional[T] {
	return TypeAssert[T](v)
}

// TypeAssert try to convert the value to the given type.
//
// if the value is not convertible to the given type,
// then the function returns None,
// otherwise, it returns Some.
func TypeAssert[T any](v any) Optional[T] {
	val, ok := v.(T)
	return FromPair(val, ok)
}
