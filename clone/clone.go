package clone

import (
	"github.com/go-board/std/delegate"
)

type Cloneable[T any] interface{ Clone() T }

func Clone[T Cloneable[T]](o T) T { return o.Clone() }

type Cloner[T any] delegate.Function1[T, T]
