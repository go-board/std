package result

import "github.com/go-board/std/delegate"

type Result[Ok any] struct {
	data Ok
	err  error
}

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

func (self Result[Ok]) IfOk(consume delegate.Consumer1[Ok]) {
	if self.IsOk() {
		consume(self.Value())
	}
}

func (self Result[Ok]) IfErr(consume delegate.Consumer1[error]) {
	if self.IsErr() {
		consume(self.Error())
	}
}
