package delegate

type Transform[In, Out any] func(In) Out
