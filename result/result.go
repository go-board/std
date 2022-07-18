package result

import (
	"fmt"
)

// Result is a type that represents either a value or an error.
type Result[Ok any] struct {
	data Ok
	err  error
}

func (self Result[Ok]) String() string {
	if self.IsOk() {
		return fmt.Sprintf("Ok(%+v)", self.data)
	}
	return fmt.Sprintf("Err(%+v)", self.err)
}

// Of create a new Result from a value or an error.
func Of[Ok any](data Ok, err error) Result[Ok] {
	return Result[Ok]{data, err}
}

// Ok create a new Result from a value.
func Ok[Ok any](data Ok) Result[Ok] {
	return Of(data, nil)
}

// Err create a new Result from an error.
func Err[Ok any](err error) Result[Ok] {
	var ok Ok
	return Of(ok, err)
}

// Errorf create a new Result from a formatted error.
func Errorf[Ok any](format string, args ...any) Result[Ok] {
	return Err[Ok](fmt.Errorf(format, args...))
}

// IsOk returns true if the Result is Ok.
func (self Result[Ok]) IsOk() bool { return self.err == nil }

// IsErr returns true if the Result is Err.
func (self Result[Ok]) IsErr() bool { return !self.IsOk() }

// And return self if self is error, otherwise return other.
func (self Result[Ok]) And(res Result[Ok]) Result[Ok] {
	if self.IsOk() {
		return res
	}
	return Err[Ok](self.err)
}

// Or return self if self is ok, otherwise return other.
func (self Result[T]) Or(res Result[T]) Result[T] {
	if self.IsOk() {
		return Ok(self.data)
	}
	return res
}

// Value unwrap the value of the Result, panic if the Result is Err.
func (self Result[Ok]) Value() Ok {
	if self.IsOk() {
		return self.data
	}
	panic("unwrap error value")
}

// Error unwrap the error of the Result, panic if the Result is Ok.
func (self Result[Ok]) Error() error {
	if self.IsErr() {
		return self.err
	}
	panic("unwrap ok value")
}

// IfOk call the function if the Result is Ok.
func (self Result[Ok]) IfOk(consume func(Ok)) {
	if self.IsOk() {
		consume(self.Value())
	}
}

// IfErr call the function if the Result is Err.
func (self Result[Ok]) IfErr(consume func(error)) {
	if self.IsErr() {
		consume(self.Error())
	}
}

// Match call consumeOk if the Result is Ok, otherwise call consumeErr.
func (self Result[Ok]) Match(consumeOk func(Ok), consumeErr func(error)) {
	if self.IsErr() {
		consumeErr(self.Error())
	} else {
		consumeOk(self.Value())
	}
}

// AsRawParts return tuple of value ptr and error.
func (self Result[Ok]) AsRawParts() (*Ok, error) {
	if self.IsOk() {
		return &self.data, nil
	}
	return nil, self.err
}
