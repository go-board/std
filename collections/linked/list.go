package linked

import (
	"github.com/go-board/std/iter"
	"github.com/go-board/std/optional"
)

type List[T any] struct {
	head *listNode[T]
	tail *listNode[T]
}

var _ iter.Iterable[any] = (*List[any])(nil)

type listNode[T any] struct {
	data T
	next *listNode[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func FromSlice[T any](elems ...T) *List[T] {
	list := NewList[T]()
	for _, ele := range elems {
		list.Append(ele)
	}
	return list
}

func FromIterator[T any](iter iter.Iter[T]) *List[T] {
	list := NewList[T]()
	for e := iter.Next(); e.IsSome(); e = iter.Next() {
		list.Append(e.Value())
	}
	return list
}

func (self *List[T]) Iter() iter.Iter[T] {
	return &listIter[T]{nextNode: self.head, prevNode: self.tail}
}

func (self *List[T]) Append(data T) {
	listNode := &listNode[T]{data: data, next: nil}
	if self.tail != nil {
		self.tail = listNode
	}
	self.head = listNode
	self.tail = listNode
}

type listIter[T any] struct {
	nextNode *listNode[T]
	prevNode *listNode[T]
}

var _ iter.Iter[any] = (*listIter[any])(nil)

func (self *listIter[T]) Next() optional.Optional[T] {
	if self.nextNode != nil {
		value := self.nextNode.data
		self.nextNode = self.nextNode.next
		return optional.Some(value)
	}
	return optional.None[T]()
}
