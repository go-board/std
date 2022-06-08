package result

func Map[A, B any](result Result[A], transformer func(A) B) Result[B] {
	if result.IsOk() {
		return Ok(transformer(result.Value()))
	}
	return Err[B](result.Error())
}
