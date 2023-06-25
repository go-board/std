package linkedlist

import (
	"github.com/go-board/std/iter"
	"github.com/go-board/std/optional"
)

type LinkedList[T any] struct {
	head *linkedListNode[T]
	tail *linkedListNode[T]
}

var _ iter.Iterable[any] = (*LinkedList[any])(nil)

type linkedListNode[T any] struct {
	data T
	next *linkedListNode[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func FromSlice[T any](elems ...T) *LinkedList[T] {
	list := NewLinkedList[T]()
	for _, ele := range elems {
		list.Append(ele)
	}
	return list
}

func FromIterator[T any](iter iter.Iter[T]) *LinkedList[T] {
	list := NewLinkedList[T]()
	for e := iter.Next(); e.IsSome(); e = iter.Next() {
		list.Append(e.Value())
	}
	return list
}

func (self *LinkedList[T]) Iter() iter.Iter[T] {
	return &linkedListIter[T]{list: self, iterNode: self.head}
}

func (self *LinkedList[T]) Append(data T) {
	listNode := &linkedListNode[T]{data: data, next: nil}
	if self.tail != nil {
		self.tail = listNode
	}
	self.head = listNode
	self.tail = listNode
}

type linkedListIter[T any] struct {
	list     *LinkedList[T]
	iterNode *linkedListNode[T]
}

var _ iter.Iter[any] = (*linkedListIter[any])(nil)

func (self *linkedListIter[T]) Next() optional.Optional[T] {
	if self.iterNode != nil {
		value := self.iterNode.data
		self.iterNode = self.iterNode.next
		return optional.Some(value)
	}
	return optional.None[T]()
}
