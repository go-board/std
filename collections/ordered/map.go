package ordered

import (
	"github.com/go-board/std/clone"
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/iter"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/tuple"
	"github.com/tidwall/btree"
)

func invert[T any](cmp func(T, T) int) func(T, T) int {
	return func(lhs, rhs T) int { return -cmp(lhs, rhs) }
}

// MapEntry is a tuple of key and value.
type MapEntry[K, V any] struct{ inner tuple.Pair[K, V] }

// MapEntryOf creates a new MapEntry.
func MapEntryOf[K, V any](key K, value V) MapEntry[K, V] {
	return MapEntry[K, V]{inner: tuple.NewPair(key, value)}
}

// Key returns the key of the MapEntry.
func (self MapEntry[K, V]) Key() K { return self.inner.First }

// Value returns the value of the MapEntry.
func (self MapEntry[K, V]) Value() V { return self.inner.Second }

// Map is an ordered map based on a B-Tree.
type Map[K, V any] struct {
	cmp   func(K, K) int
	inner *btree.BTreeG[MapEntry[K, V]]
}

var dummyMap *Map[any, any]

var _ clone.Cloneable[*Map[any, any]] = dummyMap

// NewMap creates a new Map.
func NewMap[K, V any](cmp func(K, K) int) *Map[K, V] {
	less := func(a, b MapEntry[K, V]) bool { return cmp(a.Key(), b.Key()) < 0 }
	return &Map[K, V]{inner: btree.NewBTreeG(less), cmp: cmp}
}

// NewOrderedMap creates a new Map from Ordered type
func NewOrderedMap[K cmp.Ordered, V any]() *Map[K, V] {
	return NewMap[K, V](cmp.Compare[K])
}

func (self *Map[K, V]) keyEntry(k K) MapEntry[K, V] {
	return MapEntry[K, V]{inner: tuple.Pair[K, V]{First: k}}
}

// Insert inserts a new MapEntry.
func (self *Map[K, V]) Insert(key K, value V) optional.Optional[V] {
	e, ok := self.inner.Set(MapEntryOf(key, value))
	return optional.Map(optional.FromPair(e, ok), MapEntry[K, V].Value)
}

func (self *Map[K, V]) InsertIter(it iter.Seq[MapEntry[K, V]]) {
	iter.ForEach(it, func(me MapEntry[K, V]) {
		self.insertEntry(me)
	})
}

func (self *Map[K, V]) insertEntry(entry MapEntry[K, V]) {
	self.inner.Set(entry)
}

// Get returns the value for the given key.
func (self *Map[K, V]) Get(key K) optional.Optional[V] {
	return optional.Map(self.GetEntry(key), MapEntry[K, V].Value)
}

// GetDefault returns the value for the given key or the default value.
func (self *Map[K, V]) GetDefault(key K, value V) V {
	return self.Get(key).ValueOr(value)
}

// GetEntry returns the MapEntry for the given key.
func (self *Map[K, V]) GetEntry(key K) optional.Optional[MapEntry[K, V]] {
	return optional.FromPair(self.inner.Get(self.keyEntry(key)))
}

// Clone returns a copy of the Map.
func (self *Map[K, V]) Clone() *Map[K, V] {
	return &Map[K, V]{inner: self.inner.Copy()}
}

// Reverse returns a reversed copy of the Map.
func (self *Map[K, V]) Reverse() *Map[K, V] {
	newTree := NewMap[K, V](invert(self.cmp))
	iter.ForEach(self.EntryIter(), newTree.insertEntry)
	return newTree
}

// ContainsKey returns true if the Map contains the given key.
func (self *Map[K, V]) ContainsKey(key K) bool {
	_, ok := self.inner.Get(self.keyEntry(key))
	return ok
}

// Remove removes the MapEntry for the given key.
func (self *Map[K, V]) Remove(key K) {
	self.inner.Delete(self.keyEntry(key))
}

// First returns the first MapEntry.
func (self *Map[K, V]) First() optional.Optional[MapEntry[K, V]] {
	return optional.FromPair(self.inner.Min())
}

// Last returns the last MapEntry.
func (self *Map[K, V]) Last() optional.Optional[MapEntry[K, V]] {
	return optional.FromPair(self.inner.Max())
}

// PopFirst removes and returns the first MapEntry.
func (self *Map[K, V]) PopFirst() optional.Optional[MapEntry[K, V]] {
	return optional.FromPair(self.inner.PopMin())
}

// PopLast removes and returns the last MapEntry.
func (self *Map[K, V]) PopLast() optional.Optional[MapEntry[K, V]] {
	return optional.FromPair(self.inner.PopMax())
}

// KeyIter returns an iterator over the keys in the map.
func (self *Map[K, V]) KeyIter() iter.Seq[K] {
	return iter.Map(self.inner.Scan, func(e MapEntry[K, V]) K { return e.Key() })
}

// ValueIter returns an iterator over the keys in the map.
func (self *Map[K, V]) ValueIter() iter.Seq[V] {
	return iter.Map(self.inner.Scan, func(e MapEntry[K, V]) V { return e.Value() })
}

// EntryIter returns an iterator over the keys in the map.
func (self *Map[K, V]) EntryIter() iter.Seq[MapEntry[K, V]] {
	return self.inner.Scan
}

// Len returns the number of entries in the map.
func (self *Map[K, V]) Len() int {
	return self.inner.Len()
}

func (self *Map[K, V]) Merge(o *Map[K, V]) {
	o.inner.Scan(func(item MapEntry[K, V]) bool {
		self.Insert(item.Key(), item.Value())
		return true
	})
}
