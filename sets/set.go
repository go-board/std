package sets

import (
	"encoding/json"

	"github.com/go-board/std/core"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/iterator/adapters"
)

var unit = struct{}{}

type HashSet[TKey comparable] struct{ inner map[TKey]core.Unit }

var (
	_ iterator.Iterable[core.Unit] = (*HashSet[core.Unit])(nil)
	_ json.Marshaler               = (*HashSet[core.Unit])(nil)
	_ json.Unmarshaler             = (*HashSet[core.Unit])(nil)
)

// NewHashSet returns a new empty hash set.
func NewHashSet[TKey comparable](elements ...TKey) *HashSet[TKey] {
	inner := make(map[TKey]core.Unit, len(elements))
	for _, element := range elements {
		inner[element] = unit
	}
	return &HashSet[TKey]{inner: inner}
}

func FromMapKey[TKey comparable, TValue any, M ~map[TKey]TValue](m M) *HashSet[TKey] {
	inner := make(map[TKey]core.Unit, len(m))
	for key := range m {
		inner[key] = unit
	}
	return &HashSet[TKey]{inner: inner}
}

// Add adds the given keys to the set.
func (self *HashSet[TKey]) Add(keys ...TKey) {
	for _, key := range keys {
		self.inner[key] = unit
	}
}

// AddAll adds all elements in other [HashSet]
func (self *HashSet[TKey]) AddAll(other *HashSet[TKey]) {
	self.Add(other.ToSlice()...)
}

// Remove removes the given key from the set.
func (self *HashSet[TKey]) Remove(key TKey) {
	self.RemoveBy(func(k TKey) bool { return k == key })
}

// RemoveBy remove keys from the set if the given predicate returns true.
func (self *HashSet[TKey]) RemoveBy(predicate func(TKey) bool) {
	for k := range self.inner {
		if predicate(k) {
			delete(self.inner, k)
		}
	}
}

// Clear removes all keys from the set.
func (self *HashSet[TKey]) Clear() {
	for k := range self.inner {
		delete(self.inner, k)
	}
}

func (self *HashSet[TKey]) ForEach(fn func(TKey)) {
	for k := range self.inner {
		fn(k)
	}
}

func (self *HashSet[TKey]) Filter(fn func(TKey) bool) *HashSet[TKey] {
	other := NewHashSet[TKey]()
	self.ForEach(func(t TKey) {
		if fn(t) {
			other.Add(t)
		}
	})
	return other
}

// Contains returns true if the given key is in the set.
func (self *HashSet[TKey]) Contains(key TKey) bool {
	_, exists := self.inner[key]
	return exists
}

// ContainsAll returns true if all the given keys are in the set.
func (self *HashSet[TKey]) ContainsAll(keys []TKey) bool {
	for _, key := range keys {
		if !self.Contains(key) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if any of the given keys are in the set.
func (self *HashSet[TKey]) ContainsAny(keys []TKey) bool {
	for _, key := range keys {
		if self.Contains(key) {
			return true
		}
	}
	return false
}

// Size returns the number of elements in the set.
func (self *HashSet[TKey]) Size() int {
	return len(self.inner)
}

// IsEmpty returns true if the set is empty.
func (self *HashSet[TKey]) IsEmpty() bool {
	return self.Size() == 0
}

// ToSlice returns a slice containing all elements in the set.
func (self *HashSet[TKey]) ToSlice() []TKey {
	keys := make([]TKey, self.Size())
	i := 0
	for key := range self.inner {
		keys[i] = key
		i++
	}
	return keys
}

func (self *HashSet[TKey]) ToMap() map[TKey]struct{} {
	m := make(map[TKey]struct{}, len(self.inner))
	for k, v := range self.inner {
		m[k] = v
	}
	return m
}

// Clone returns a copy of the set.
func (self *HashSet[TKey]) Clone() *HashSet[TKey] {
	return self.DeepCloneBy(func(t TKey) TKey { return t })
}

// DeepCloneBy returns a copy of the set and clone each element use given clone func.
func (self *HashSet[TKey]) DeepCloneBy(clone func(TKey) TKey) *HashSet[TKey] {
	other := NewHashSet[TKey]()
	for key := range self.inner {
		other.Add(clone(key))
	}
	return other
}

// SupersetOf returns true if the given set is a superset of this set.
func (self *HashSet[TKey]) SupersetOf(other *HashSet[TKey]) bool {
	for key := range other.inner {
		if !self.Contains(key) {
			return false
		}
	}
	return true
}

// SubsetOf returns true if the given set is a subset of this set.
func (self *HashSet[TKey]) SubsetOf(other *HashSet[TKey]) bool {
	for key := range self.inner {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// Union returns a new set containing all the elements that are in either set.
func (self *HashSet[TKey]) Union(other *HashSet[TKey]) *HashSet[TKey] {
	union := NewHashSet[TKey]()
	union.AddAll(self)
	union.AddAll(other)
	return union
}

// UnionAssign union another [HashSet] into self
func (self *HashSet[TKey]) UnionAssign(other *HashSet[TKey]) {
	self.AddAll(other)
}

// Intersection returns a new set containing all the elements that are in both sets.
func (self *HashSet[TKey]) Intersection(other *HashSet[TKey]) *HashSet[TKey] {
	intersection := NewHashSet[TKey]()
	for key := range self.inner {
		if other.Contains(key) {
			intersection.Add(key)
		}
	}
	return intersection
}

// Difference returns a new set containing all the elements that are in this set but not in the other set.
func (self *HashSet[TKey]) Difference(other *HashSet[TKey]) *HashSet[TKey] {
	diff := NewHashSet[TKey]()
	for key := range self.inner {
		if !other.Contains(key) {
			diff.Add(key)
		}
	}
	return diff
}

// SymmetricDifference returns a new set containing all the elements that are in this set or the other set but not in both.
func (self *HashSet[TKey]) SymmetricDifference(other *HashSet[TKey]) *HashSet[TKey] {
	diff := NewHashSet[TKey]()
	for key := range self.inner {
		if !other.Contains(key) {
			diff.Add(key)
		}
	}
	for key := range other.inner {
		if !self.Contains(key) {
			diff.Add(key)
		}
	}
	return diff
}

// Equal returns true if the given set is equal to this set.
func (self *HashSet[TKey]) Equal(other *HashSet[TKey]) bool {
	if self.Size() != other.Size() {
		return false
	}
	for key := range other.inner {
		if _, ok := self.inner[key]; !ok {
			return false
		}
	}
	return true
}

func (self *HashSet[TKey]) Iter() iterator.Iterator[TKey] {
	return adapters.OfSlice(self.ToSlice()...)
}

func (self *HashSet[TKey]) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.inner)
}

func (self *HashSet[TKey]) UnmarshalJSON(v []byte) error {
	return json.Unmarshal(v, &self.inner)
}

// Map maps each element in current [HashSet] to a new [HashSet].
func Map[TKey, TNewKey comparable](
	set *HashSet[TKey],
	fn func(TKey) TNewKey,
) *HashSet[TNewKey] {
	newSet := NewHashSet[TNewKey]()
	for key := range set.inner {
		newSet.Add(fn(key))
	}
	return newSet
}
