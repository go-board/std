package iter

type Iterable[T any] interface {
	Iter() Iter[T]
}

type PrevIterable[T any] interface {
	PrevIter() PrevIter[T]
}
