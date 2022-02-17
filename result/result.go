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

func (r Result[Ok]) IsOk() bool  { return r.err == nil }
func (r Result[Ok]) IsErr() bool { return !r.IsOk() }

func (r Result[Ok]) Value() Ok {
	if r.IsOk() {
		return r.data
	}
	panic("unwrap error value")
}

func (r Result[Ok]) Error() error {
	if r.IsErr() {
		return r.err
	}
	panic("unwrap ok value")
}

func (r Result[Ok]) IfOk(consume delegate.Consumer1[Ok]) {
	if r.IsOk() {
		consume(r.Value())
	}
}

func (r Result[Ok]) IfErr(consume delegate.Consumer1[error]) {
	if r.IsErr() {
		consume(r.Error())
	}
}
