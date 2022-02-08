package iterator

type Iterable[T any] interface {
	Iter() Iterator[T]
}

type DoubleEndedIterable[T any] interface {
	DoubleEndedIter() DoubleEndedIterator[T]
}
