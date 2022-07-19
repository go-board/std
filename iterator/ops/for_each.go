package ops

import (
	"github.com/go-board/std/iterator"
)

func ForEach[T any](iter iterator.Iterator[T], consumer func(T)) {
	iterate := func(x T) bool { consumer(x); return true }
	ForEachUntil(iter, iterate)
}

func ForEachUntil[T any](iter iterator.Iterator[T], predicate func(T) bool) {
	for e := iter.Next(); e.IsSome(); e = iter.Next() {
		if !predicate(e.Value()) {
			return
		}
	}
}
