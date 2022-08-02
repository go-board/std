package result

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
