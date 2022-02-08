package internal

import "github.com/go-board/std/iterator"

func Chain[T any, I iterator.Iterator[T], IA iterator.Iterable[T]](
	iter I,
	iterable IA,
) I {
	return iter
}
