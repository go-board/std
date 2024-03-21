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

// MakeMapEntry creates a new MapEntry.
func MakeMapEntry[K, V any](key K, value V) MapEntry[K, V] {
	return MapEntry[K, V]{inner: tuple.MakePair(key, value)}
}

// Key returns the key of the MapEntry.
func (self MapEntry[K, V]) Key() K { return self.inner.First() }

// Value returns the value of the MapEntry.
func (self MapEntry[K, V]) Value() V { return self.inner.Second() }

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
	var v V
	return MapEntry[K, V]{inner: tuple.MakePair(k, v)}
}

// Insert inserts a new MapEntry.
func (self *Map[K, V]) Insert(key K, value V) optional.Optional[V] {
	e, ok := self.inner.Set(MakeMapEntry(key, value))
	return optional.Map(optional.FromPair(e, ok), MapEntry[K, V].Value)
}

func (self *Map[K, V]) InsertIter(it iter.Seq[MapEntry[K, V]]) {
	iter.ForEach(it, self.insertEntry)
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
	iter.ForEach(self.Entries(), newTree.insertEntry)
	return newTree
}

// ContainsKey returns true if the Map contains the given key.
func (self *Map[K, V]) ContainsKey(key K) bool {
	_, ok := self.inner.Get(self.keyEntry(key))
	return ok
}

// ContainsAll tests is all elements in [iter.Seq] is a valid map key.
func (self *Map[K, V]) ContainsAll(it iter.Seq[K]) bool {
	return iter.All(it, self.ContainsKey)
}

// ContainsAny tests is any elements in [Seq] is a valid map key.
func (self *Map[K, V]) ContainsAny(it iter.Seq[K]) bool {
	return iter.Any(it, self.ContainsKey)
}

// Remove removes the MapEntry for the given key.
func (self *Map[K, V]) Remove(key K) {
	self.inner.Delete(self.keyEntry(key))
}

// RemoveIter remove all elements in [iter.Seq].
func (self *Map[K, V]) RemoveIter(it iter.Seq[K]) {
	iter.ForEach(it, self.Remove)
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

// Keys returns an iterator over the keys in the map.
func (self *Map[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		self.inner.Scan(func(item MapEntry[K, V]) bool { return yield(item.Key()) })
	}
}

func (self *Map[K, V]) KeysMut() iter.Seq[*KeyItem[K, V]] {
	return func(yield func(*KeyItem[K, V]) bool) {
		self.inner.Scan(func(item MapEntry[K, V]) bool {
			return yield(&KeyItem[K, V]{key: item.Key(), m: self})
		})
	}
}

// Values returns an iterator over the keys in the map.
func (self *Map[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		self.inner.Scan(func(item MapEntry[K, V]) bool { return yield(item.Value()) })
	}
}

func (self *Map[K, V]) ValuesMut() iter.Seq[*ValueItem[K, V]] {
	return func(yield func(*ValueItem[K, V]) bool) {
		self.inner.Scan(func(item MapEntry[K, V]) bool {
			return yield(&ValueItem[K, V]{key: item.Key(), m: self})
		})
	}
}

// Entries returns an iterator over the keys in the map.
func (self *Map[K, V]) Entries() iter.Seq[MapEntry[K, V]] { return self.inner.Scan }

// Len returns the number of entries in the map.
func (self *Map[K, V]) Len() int {
	return self.inner.Len()
}

func (self *Map[K, V]) MergeKeep(o *Map[K, V]) {
	self.MergeFunc(o, func(_ K, prev V, _ V) V { return prev })
}

func (self *Map[K, V]) MergeOverwrite(o *Map[K, V]) {
	self.MergeFunc(o, func(key K, prev V, current V) V { return current })
}

func (self *Map[K, V]) MergeFunc(o *Map[K, V], solve func(key K, prev V, current V) V) {
	o.inner.Scan(func(item MapEntry[K, V]) bool {
		prev := self.GetEntry(item.Key())
		v := item.Value()
		if prev.IsSome() {
			v = solve(item.Key(), prev.Value().Value(), v)
		}
		self.Insert(item.Key(), v)
		return true
	})
}

type KeyItem[K any, V any] struct {
	key K
	m   *Map[K, V]
}

func (k *KeyItem[K, V]) Key() K   { return k.key }
func (k *KeyItem[K, V]) Value() V { return k.m.Get(k.key).ValueOrZero() }
func (k *KeyItem[K, V]) Set(v V)  { k.m.Insert(k.key, v) }
func (k *KeyItem[K, V]) Remove()  { k.m.Remove(k.key) }

type ValueItem[K any, V any] struct {
	val V
	key K
	m   *Map[K, V]
}

func (v *ValueItem[K, V]) Value() V  { return v.val }
func (v *ValueItem[K, V]) Set(val V) { v.m.Insert(v.key, val) }
func (v *ValueItem[K, V]) Remove()   { v.m.Remove(v.key) }
