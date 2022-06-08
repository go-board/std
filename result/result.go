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

// Of creates a new Result from a value.
func Of[Ok any](data Ok, err error) Result[Ok] {
	return Result[Ok]{data, err}
}

func Ok[Ok any](data Ok) Result[Ok] {
	return Of(data, nil)
}

func Err[Ok any](err error) Result[Ok] {
	var ok Ok
	return Of(ok, err)
}

func (self Result[Ok]) IsOk() bool  { return self.err == nil }
func (self Result[Ok]) IsErr() bool { return !self.IsOk() }

func (self Result[Ok]) And(res Result[Ok]) Result[Ok] {
	if self.IsOk() {
		return res
	}
	return Err[Ok](self.err)
}

func (self Result[T]) Or(res Result[T]) Result[T] {
	if self.IsOk() {
		return Ok(self.data)
	}
	return res
}

func (self Result[Ok]) Value() Ok {
	if self.IsOk() {
		return self.data
	}
	panic("unwrap error value")
}

func (self Result[Ok]) Error() error {
	if self.IsErr() {
		return self.err
	}
	panic("unwrap ok value")
}

func (self Result[Ok]) IfOk(consume func(Ok)) {
	if self.IsOk() {
		consume(self.Value())
	}
}

func (self Result[Ok]) IfErr(consume func(error)) {
	if self.IsErr() {
		consume(self.Error())
	}
}

func (self Result[Ok]) Match(consumeOk func(Ok), consumeErr func(error)) {
	if self.IsErr() {
		consumeErr(self.Error())
	} else {
		consumeOk(self.Value())
	}
}
