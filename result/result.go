package result

import (
	"fmt"
)

// Result is a type that represents either a value or an error.
type Result[T any] struct {
	data T
	err  error
}

// String implements fmt.Stringer.
func (self Result[T]) String() string {
	if self.IsOk() {
		return fmt.Sprintf("Ok(%+v)", self.data)
	}
	return fmt.Sprintf("Err(%+v)", self.err)
}

// FromPair create a new Result from a value or an error.
func FromPair[T any](data T, err error) Result[T] { return Result[T]{data, err} }

// Ok create a new Result from a value.
func Ok[T any](data T) Result[T] { return FromPair(data, nil) }

// Err create a new Result from an error.
func Err[T any](err error) Result[T] {
	var ok T
	return FromPair(ok, err)
}

// Errorf create a new Result from a formatted error.
func Errorf[T any](format string, args ...any) Result[T] {
	return Err[T](fmt.Errorf(format, args...))
}

// IsOk returns true if the Result is Ok.
func (self Result[T]) IsOk() bool { return self.err == nil }

// IsErr returns true if the Result is Err.
func (self Result[T]) IsErr() bool { return !self.IsOk() }

// And return self if self is error, otherwise return other.
func (self Result[T]) And(res Result[T]) Result[T] {
	if self.IsOk() {
		return res
	}
	return Err[T](self.err)
}

// Or return self if self is ok, otherwise return other.
func (self Result[T]) Or(res Result[T]) Result[T] {
	if self.IsOk() {
		return Ok(self.data)
	}
	return res
}

// Map returns a new Result from the result of applying
// the given function to the value of self if ok, else return err.
//
// FIXME: for now, go doesn't support type parameter on method, so we have to use the same type parameter on the method.
func (self Result[T]) Map(fn func(T) T) Result[T] {
	if self.IsOk() {
		return Ok(fn(self.data))
	}
	return Err[T](self.err)
}

// Value unwrap the value of the Result, panic if the Result is Err.
func (self Result[T]) Value() T {
	if self.IsOk() {
		return self.data
	}
	panic("unwrap error value")
}

// ValueOr returns the value of the Result if it is Ok, otherwise return the given value.
func (self Result[T]) ValueOr(defaultValue T) T {
	if self.IsOk() {
		return self.data
	}
	return defaultValue
}

// ValueOrElse returns the value of the Result if it is Ok,
// otherwise return the result of calling the given function.
func (self Result[T]) ValueOrElse(defaultFunc func() T) T {
	if self.IsOk() {
		return self.data
	}
	return defaultFunc()
}

// ValueOrZero returns the value of the Result if it is Ok,
// otherwise return the zero value of the type of the Result.
func (self Result[T]) ValueOrZero() T {
	if self.IsOk() {
		return self.data
	}
	var t T
	return t
}

// Error unwrap the error of the Result, panic if the Result is Ok.
func (self Result[T]) Error() error {
	if self.IsErr() {
		return self.err
	}
	panic("unwrap ok value")
}

// IfOk call the function if the Result is Ok.
func (self Result[T]) IfOk(consume func(T)) {
	if self.IsOk() {
		consume(self.Value())
	}
}

// IfErr call the function if the Result is Err.
func (self Result[T]) IfErr(consume func(error)) {
	if self.IsErr() {
		consume(self.Error())
	}
}

// Match call onOk if the Result is Ok, otherwise call onErr.
func (self Result[T]) Match(onOk func(T), onErr func(error)) {
	if self.IsErr() {
		onErr(self.err)
	} else {
		onOk(self.data)
	}
}

// AsRawParts return tuple of value ptr and error.
func (self Result[T]) AsRawParts() (data T, err error) {
	if self.IsOk() {
		data = self.data
	} else {
		err = self.err
	}
	return
}

func Map[A, B any](result Result[A], transformer func(A) B) Result[B] {
	if result.IsOk() {
		return Ok(transformer(result.Value()))
	}
	return Err[B](result.Error())
}

func Flatten[A any](result Result[Result[A]]) Result[A] {
	if result.IsOk() {
		return result.data
	}
	return Err[A](result.err)
}
