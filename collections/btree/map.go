package btree

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/tuple"
	"github.com/tidwall/btree"
)

type MapEntry[TKey, TValue any] struct{ inner tuple.Pair[TKey, TValue] }

func NewMapEntry[TKey, TValue any](key TKey, value TValue) MapEntry[TKey, TValue] {
	return MapEntry[TKey, TValue]{inner: tuple.NewPair(key, value)}
}

func (self MapEntry[K, V]) Key() K   { return self.inner.First() }
func (self MapEntry[K, V]) Value() V { return self.inner.Second() }

type TreeMap[TKey, TValue any] struct {
	inner *btree.Generic[MapEntry[TKey, TValue]]
}

func NewTreeMap[TKey, TValue any](less delegate.Lt[TKey]) *TreeMap[TKey, TValue] {
	entryLessFn := func(a, b MapEntry[TKey, TValue]) bool { return less(a.Key(), b.Key()) }
	return &TreeMap[TKey, TValue]{inner: btree.NewGeneric(entryLessFn)}
}

func (self *TreeMap[TKey, TValue]) Put(key TKey, value TValue) {
	self.inner.Set(NewMapEntry(key, value))
}

func (self *TreeMap[TKey, TValue]) Get(key TKey) optional.Optional[TValue] {
	entry, ok := self.inner.Get(MapEntry[TKey, TValue]{
		inner: tuple.PairFromA[TKey, TValue](key),
	})
	if ok {
		return optional.Some(entry.Value())
	}
	return optional.None[TValue]()
}

func (self *TreeMap[TKey, TValue]) GetDefault(key TKey, value TValue) TValue {
	return self.Get(key).OrElse(value)
}

func (self *TreeMap[TKey, TValue]) IterKey() iterator.Iterator[TKey] {
	return &treeMapKeyIter[TKey, TValue]{
		inner: &treeMapIter[TKey, TValue]{inner: self.inner.Iter()},
	}
}

func (self *TreeMap[TKey, TValue]) IterValue() iterator.Iterator[TValue] {
	return &treeMapValueIter[TKey, TValue]{
		inner: &treeMapIter[TKey, TValue]{inner: self.inner.Iter()},
	}
}

func (self *TreeMap[TKey, TValue]) Size() uint {
	return uint(self.inner.Len())
}

func (self *TreeMap[TKey, TValue]) Iter() iterator.Iterator[MapEntry[TKey, TValue]] {
	return &treeMapIter[TKey, TValue]{inner: self.inner.Iter()}
}

type treeMapIter[TKey, TValue any] struct {
	inner btree.GenericIter[MapEntry[TKey, TValue]]
}

var _ iterator.Iterator[MapEntry[any, any]] = (*treeMapIter[any, any])(nil)

func (self *treeMapIter[TKey, TValue]) Next() optional.Optional[MapEntry[TKey, TValue]] {
	if self.inner.Next() {
		return optional.Some(self.inner.Item())
	}
	return optional.None[MapEntry[TKey, TValue]]()
}

type treeMapKeyIter[TKey, TValue any] struct {
	inner *treeMapIter[TKey, TValue]
}

func (self *treeMapKeyIter[TKey, TValue]) Next() optional.Optional[TKey] {
	return optional.Map(self.inner.Next(), func(e MapEntry[TKey, TValue]) TKey { return e.Key() })
}

type treeMapValueIter[TKey, TValue any] struct {
	inner *treeMapIter[TKey, TValue]
}

func (self *treeMapValueIter[TKey, TValue]) Next() optional.Optional[TValue] {
	return optional.Map(self.inner.Next(), func(e MapEntry[TKey, TValue]) TValue { return e.Value() })
}
