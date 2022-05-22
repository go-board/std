package optional

import (
	"github.com/go-board/std/clone"
)

func Clone[T clone.Cloneable[T]](opt Optional[T]) Optional[T] {
	if opt.IsSome() {
		return Some(opt.Value().Clone())
	}
	return None[T]()
}

func Map[A, B any](opt Optional[A], mapFn func(A) B) Optional[B] {
	if opt.IsSome() {
		return Some(mapFn(opt.Value()))
	}
	return None[B]()
}
