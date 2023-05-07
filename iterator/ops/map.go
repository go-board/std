package ops

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

func Map[T, U any](it iterator.Iter[T], transform func(T) U) iterator.Iter[U] {
	return iterator.IterFunc[U](func() optional.Optional[U] {
		return optional.Map(it.Next(), transform)
	})
}
