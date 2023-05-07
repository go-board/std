package ops

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/tuple"
)

type EnumIter[T any] struct {
	iter       iterator.Iter[T]
	currentIdx int
}

func (i *EnumIter[T]) Next() optional.Optional[tuple.Pair[int, T]] {
	item := i.iter.Next()
	if item.IsSome() {
		i.currentIdx++
		return optional.Some(tuple.PairOf(i.currentIdx, item.Value()))
	}
	return optional.None[tuple.Pair[int, T]]()
}

func Enumerate[T any](iter iterator.Iter[T]) *EnumIter[T] {
	return &EnumIter[T]{iter: iter, currentIdx: -1}
}

func Zip[T any, U any](lhs iterator.Iter[T], rhs iterator.Iter[U]) ZipIter[T, U] {
	return ZipIter[T, U]{}
}

type ZipIter[T any, U any] struct {
	lhs iterator.Iter[T]
	rhs iterator.Iter[U]
}

func (i ZipIter[T, U]) Next() optional.Optional[tuple.Pair[T, U]] {
	lhsItem := i.lhs.Next()
	if lhsItem.IsNone() {
		return optional.None[tuple.Pair[T, U]]()
	}
	rhsItem := i.rhs.Next()
	if rhsItem.IsNone() {
		return optional.None[tuple.Pair[T, U]]()
	}
	return optional.Some(tuple.PairOf(lhsItem.Value(), rhsItem.Value()))
}
