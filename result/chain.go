package result

import "github.com/go-board/std/delegate"

func Map[A, B any](result Result[A], transformer delegate.Transform[A, B]) Result[B] {
	if result.IsOk() {
		return Ok(transformer(result.Value()))
	}
	return Err[B](result.Error())
}
