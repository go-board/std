package btree

import (
	"github.com/go-board/std/clone"
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/iterator/stream"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/tuple"
	"github.com/tidwall/btree"
)

// MapEntry is a tuple of key and value.
type MapEntry[TKey, TValue any] struct{ inner tuple.Pair[TKey, TValue] }

// NewMapEntry creates a new MapEntry.
func NewMapEntry[TKey, TValue any](key TKey, value TValue) MapEntry[TKey, TValue] {
	return MapEntry[TKey, TValue]{inner: tuple.NewPair(key, value)}
}

// Key returns the key of the MapEntry.
func (self MapEntry[K, V]) Key() K { return self.inner.First() }

// Value returns the value of the MapEntry.
func (self MapEntry[K, V]) Value() V { return self.inner.Second() }

// TreeMap is a map based on a B-Tree.
type TreeMap[TKey, TValue any] struct {
	less  delegate.Lt[TKey]
	inner *btree.Generic[MapEntry[TKey, TValue]]
}

var (
	_ clone.Cloneable[*TreeMap[any, any]]   = (*TreeMap[any, any])(nil)
	_ iterator.Iterable[MapEntry[any, any]] = (*TreeMap[any, any])(nil)
)

// NewTreeMap creates a new TreeMap.
func NewTreeMap[TKey, TValue any](less delegate.Lt[TKey]) *TreeMap[TKey, TValue] {
	entryLessFn := func(a, b MapEntry[TKey, TValue]) bool { return less(a.Key(), b.Key()) }
	return &TreeMap[TKey, TValue]{inner: btree.NewGeneric(entryLessFn)}
}

// Put inserts a new MapEntry.
func (self *TreeMap[TKey, TValue]) Put(key TKey, value TValue) {
	self.inner.Set(NewMapEntry(key, value))
}

// PutEntry inserts a new MapEntry.
func (self *TreeMap[TKey, TValue]) PutEntry(entry MapEntry[TKey, TValue]) {
	self.inner.Set(entry)
}

// Get returns the value for the given key.
func (self *TreeMap[TKey, TValue]) Get(key TKey) optional.Optional[TValue] {
	entry, ok := self.inner.Get(MapEntry[TKey, TValue]{
		inner: tuple.PairFromA[TKey, TValue](key),
	})
	if ok {
		return optional.Some(entry.Value())
	}
	return optional.None[TValue]()
}

// GetDefault returns the value for the given key or the default value.
func (self *TreeMap[TKey, TValue]) GetDefault(key TKey, value TValue) TValue {
	return self.Get(key).OrElse(value)
}

// GetEntry returns the MapEntry for the given key.
func (self *TreeMap[TKey, TValue]) GetEntry(key TKey) optional.Optional[MapEntry[TKey, TValue]] {
	value := self.Get(key)
	if value.IsNone() {
		return optional.None[MapEntry[TKey, TValue]]()
	}
	return optional.Some(NewMapEntry(key, value.Value()))
}

// Clone returns a copy of the TreeMap.
func (self *TreeMap[TKey, TValue]) Clone() *TreeMap[TKey, TValue] {
	return &TreeMap[TKey, TValue]{inner: self.inner.Copy()}
}

// Reversed returns a reversed copy of the TreeMap.
func (self *TreeMap[TKey, TValue]) Reversed() *TreeMap[TKey, TValue] {
	newTree := NewTreeMap[TKey, TValue](func(t1, t2 TKey) bool { return !self.less(t1, t2) })
	stream.FromIterable[MapEntry[TKey, TValue]](self).ForEach(func(entry MapEntry[TKey, TValue]) {
		newTree.Put(entry.Key(), entry.Value())
	})
	return newTree
}

// ContainsKey returns true if the TreeMap contains the given key.
func (self *TreeMap[TKey, TValue]) ContainsKey(key TKey) bool {
	_, ok := self.inner.Get(MapEntry[TKey, TValue]{
		inner: tuple.PairFromA[TKey, TValue](key),
	})
	return ok
}

// Remove removes the MapEntry for the given key.
func (self *TreeMap[TKey, TValue]) Remove(key TKey) {
	self.inner.Delete(MapEntry[TKey, TValue]{
		inner: tuple.PairFromA[TKey, TValue](key),
	})
}

// First returns the first MapEntry.
func (self *TreeMap[TKey, TValue]) First() optional.Optional[MapEntry[TKey, TValue]] {
	return optional.From(self.inner.Min())
}

// Last returns the last MapEntry.
func (self *TreeMap[TKey, TValue]) Last() optional.Optional[MapEntry[TKey, TValue]] {
	return optional.From(self.inner.Max())
}

// PopFirst removes and returns the first MapEntry.
func (self *TreeMap[TKey, TValue]) PopFirst() optional.Optional[MapEntry[TKey, TValue]] {
	return optional.From(self.inner.PopMin())
}

// PopLast removes and returns the last MapEntry.
func (self *TreeMap[TKey, TValue]) PopLast() optional.Optional[MapEntry[TKey, TValue]] {
	return optional.From(self.inner.PopMax())
}

// IterRange returns an iterator over the entries in the map that are within the range from
func (self *TreeMap[TKey, TValue]) IterRange(from, to TKey, iter func(e MapEntry[TKey, TValue])) {
	stopEntry := MapEntry[TKey, TValue]{inner: tuple.PairFromA[TKey, TValue](to)}
	self.inner.Ascend(MapEntry[TKey, TValue]{
		inner: tuple.PairFromA[TKey, TValue](from),
	}, func(e MapEntry[TKey, TValue]) bool {
		if !self.inner.Less(e, stopEntry) {
			return false
		}
		iter(e)
		return true
	})
}

// IterKey returns an iterator over the keys in the map.
func (self *TreeMap[TKey, TValue]) IterKey() iterator.Iterator[TKey] {
	return &treeMapKeyIter[TKey, TValue]{
		inner: &treeMapIter[TKey, TValue]{inner: self.inner.Iter()},
	}
}

// IterValue returns an iterator over the values in the map.
func (self *TreeMap[TKey, TValue]) IterValue() iterator.Iterator[TValue] {
	return &treeMapValueIter[TKey, TValue]{
		inner: &treeMapIter[TKey, TValue]{inner: self.inner.Iter()},
	}
}

// Size returns the number of entries in the map.
func (self *TreeMap[TKey, TValue]) Size() uint {
	return uint(self.inner.Len())
}

// Iter returns an iterator over the entries in the map.
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
