package optional

import (
	"github.com/go-board/std/clone"
	"github.com/go-board/std/delegate"
)

func Clone[T clone.Cloneable[T]](opt Optional[T]) Optional[T] {
	if opt.IsSome() {
		return Some(opt.Value().Clone())
	}
	return None[T]()
}

func Map[A, B any](opt Optional[A], mapFn delegate.Function1[A, B]) Optional[B] {
	if opt.IsSome() {
		return Some(mapFn(opt.Value()))
	}
	return None[B]()
}
