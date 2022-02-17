package linkedlist

import (
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
)

type LinkedList[T any] struct {
	head *linkedListNode[T]
	tail *linkedListNode[T]
}

var _ iterator.Iterable[any] = (*LinkedList[any])(nil)

type linkedListNode[T any] struct {
	data T
	next *linkedListNode[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func FromSlice[T any](eles ...T) *LinkedList[T] {
	list := NewLinkedList[T]()
	for _, ele := range eles {
		list.Append(ele)
	}
	return list
}

func FromIterator[T any](iter iterator.Iterator[T]) *LinkedList[T] {
	list := NewLinkedList[T]()
	for e := iter.Next(); e.IsSome(); e = iter.Next() {
		list.Append(e.Value())
	}
	return list
}

func (l *LinkedList[T]) Iter() iterator.Iterator[T] {
	return &linkedListIter[T]{list: l, iterNode: l.head}
}

func (l *LinkedList[T]) Append(data T) {
	listNode := &linkedListNode[T]{data: data, next: nil}
	if l.tail != nil {
		l.tail = listNode
	}
	l.head = listNode
	l.tail = listNode
}

type linkedListIter[T any] struct {
	list     *LinkedList[T]
	iterNode *linkedListNode[T]
}

var _ iterator.Iterator[any] = (*linkedListIter[any])(nil)

func (l *linkedListIter[T]) Next() optional.Optional[T] {
	if l.iterNode != nil {
		value := l.iterNode.data
		l.iterNode = l.iterNode.next
		return optional.Some(value)
	}
	return optional.None[T]()
}
