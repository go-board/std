package iterator

type Iterable[T any] interface {
	Iter() Iterator[T]
}

type PrevIterable[T any] interface {
	PrevIter() PrevIter[T]
}
