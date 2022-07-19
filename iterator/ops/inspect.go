package ops

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type inspectIter[T any] struct {
	iter    iterator.Iterator[T]
	inspect func(T)
}

func (self *inspectIter[T]) Next() optional.Optional[T] {
	o := self.iter.Next()
	if o.IsSome() {
		self.inspect(o.Value())
	}
	return o
}

func Inspect[T any](
	iter iterator.Iterator[T],
	inspect func(T),
) iterator.Iterator[T] {
	return &inspectIter[T]{iter, inspect}
}

type inspectIterUntil[T any] struct {
	iter    iterator.Iterator[T]
	inspect func(T) bool
}

func (self *inspectIterUntil[T]) Next() optional.Optional[T] {
	o := self.iter.Next()
	if o.IsSome() && self.inspect(o.Value()) {
		return o
	}
	return optional.None[T]()
}

func InspectUntil[T any](
	iter iterator.Iterator[T],
	inspect func(T) bool,
) iterator.Iterator[T] {
	return &inspectIterUntil[T]{iter, inspect}
}
