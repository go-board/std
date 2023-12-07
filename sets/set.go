package sets

import (
	"encoding/json"

	"github.com/go-board/std/core"
	"github.com/go-board/std/iter"
)

var unit = struct{}{}

type HashSet[E comparable] struct{ inner map[E]core.Unit }

var _set HashSet[core.Unit]

var (
	_ json.Marshaler   = _set
	_ json.Unmarshaler = _set
)

// FromSlice returns a new empty hash set.
func FromSlice[E comparable](elements ...E) HashSet[E] {
	inner := make(map[E]core.Unit, len(elements))
	for _, element := range elements {
		inner[element] = unit
	}
	return HashSet[E]{inner: inner}
}

func FromMapKeys[E comparable, TValue any, M ~map[E]TValue](m M) HashSet[E] {
	inner := make(map[E]core.Unit, len(m))
	for key := range m {
		inner[key] = unit
	}
	return HashSet[E]{inner: inner}
}

// Collect create a HashSet from [Seq].
func Collect[E comparable](s iter.Seq[E]) HashSet[E] {
	set := HashSet[E]{inner: make(map[E]core.Unit)}
	set.AddIter(s)
	return set
}

// Add adds the given keys to the set.
func (self HashSet[E]) Add(keys ...E) {
	for _, key := range keys {
		self.inner[key] = unit
	}
}

// AddAll adds all elements from another [HashSet].
func (self HashSet[E]) AddAll(other HashSet[E]) {
	self.AddIter(other.Iter())
}

// AddIter adds all elements in Iter.
func (self HashSet[E]) AddIter(s iter.Seq[E]) {
	iter.ForEach(s, func(e E) { self.inner[e] = unit })
}

// Remove removes the given key from the set.
func (self HashSet[E]) Remove(key E) {
	delete(self.inner, key)
}

func (self HashSet[E]) RemoveIter(it iter.Seq[E]) { it.ForEach(self.Remove) }

// Clear removes all keys from the set.
func (self HashSet[E]) Clear() {
	for k := range self.inner {
		delete(self.inner, k)
	}
}

func (self HashSet[E]) Filter(fn func(E) bool) HashSet[E] {
	return Collect(self.Iter().Filter(fn))
}

func (self HashSet[E]) Retain(fn func(E) bool) {
	self.Iter().ForEach(func(e E) {
		if !fn(e) {
			self.Remove(e)
		}
	})
}

func (self HashSet[E]) Map(fn func(E) E) HashSet[E] {
	m := FromSlice[E]()
	for k := range self.inner {
		m.Add(fn(k))
	}
	return m
}

// Contains returns true if the given key is in the set.
func (self HashSet[E]) Contains(key E) bool {
	_, exists := self.inner[key]
	return exists
}

// ContainsAll returns true if all the given keys are in the set.
func (self HashSet[E]) ContainsAll(keys iter.Seq[E]) bool {
	return iter.All(keys, func(e E) bool { return self.Contains(e) })
}

// ContainsAny returns true if any of the given keys are in the set.
func (self HashSet[E]) ContainsAny(keys iter.Seq[E]) bool {
	return iter.Any(keys, func(e E) bool { return self.Contains(e) })
}

// Size returns the number of elements in the set.
func (self HashSet[E]) Size() int {
	return len(self.inner)
}

// IsEmpty returns true if the set is empty.
func (self HashSet[E]) IsEmpty() bool {
	return self.Size() == 0
}

func (self HashSet[E]) ToMap() map[E]struct{} {
	m := make(map[E]struct{}, len(self.inner))
	for k, v := range self.inner {
		m[k] = v
	}
	return m
}

// Clone returns a copy of the set.
func (self HashSet[E]) Clone() HashSet[E] {
	return self.DeepCloneBy(func(t E) E { return t })
}

// DeepCloneBy returns a copy of the set and clone each element use given clone func.
func (self HashSet[E]) DeepCloneBy(clone func(E) E) HashSet[E] {
	other := FromSlice[E]()
	for key := range self.inner {
		other.Add(clone(key))
	}
	return other
}

// SupersetOf returns true if the given set is a superset of this set.
func (self HashSet[E]) SupersetOf(other HashSet[E]) bool {
	for key := range other.inner {
		if !self.Contains(key) {
			return false
		}
	}
	return true
}

// SubsetOf returns true if the given set is a subset of this set.
func (self HashSet[E]) SubsetOf(other HashSet[E]) bool {
	for key := range self.inner {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// Union returns a new set containing all the elements that are in either set.
func (self HashSet[E]) Union(other HashSet[E]) HashSet[E] {
	union := self.Clone()
	union.AddAll(other)
	return union
}

func (self HashSet[E]) UnionIter(it iter.Seq[E]) HashSet[E] {
	union := self.Clone()
	union.AddIter(it)
	return union
}

// UnionAssign union another [HashSet] into self
func (self HashSet[E]) UnionAssign(other HashSet[E]) {
	self.AddAll(other)
}

// Intersection returns a new set containing all the elements that are in both sets.
func (self HashSet[E]) Intersection(other HashSet[E]) HashSet[E] {
	intersection := FromSlice[E]()
	for key := range self.inner {
		if other.Contains(key) {
			intersection.Add(key)
		}
	}
	return intersection
}

// Difference returns a new set containing all the elements that are in this set but not in the other set.
func (self HashSet[E]) Difference(other HashSet[E]) HashSet[E] {
	diff := FromSlice[E]()
	for key := range self.inner {
		if !other.Contains(key) {
			diff.Add(key)
		}
	}
	return diff
}

// SymmetricDifference returns a new set containing all the elements that are in this set or the other set but not in both.
func (self HashSet[E]) SymmetricDifference(other HashSet[E]) HashSet[E] {
	diff := FromSlice[E]()
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
func (self HashSet[E]) Equal(other HashSet[E]) bool {
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

func (self HashSet[E]) EqualIter(it iter.Seq[E]) bool {
	return self.Equal(Collect(it))
}

// Iter returns a [iter.Seq] that iterate over the keys in the set.
func (self HashSet[E]) Iter() iter.Seq[E] {
	return func(yield func(E) bool) {
		for key := range self.inner {
			if !yield(key) {
				break
			}
		}
	}
}

func (self HashSet[E]) IterMut() iter.Seq[*SetItem[E]] {
	return func(yield func(*SetItem[E]) bool) {
		for key := range self.inner {
			if !yield(&SetItem[E]{item: key, s: self}) {
				break
			}
		}
	}
}

func (self HashSet[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.inner)
}

func (self HashSet[E]) UnmarshalJSON(v []byte) error {
	return json.Unmarshal(v, &self.inner)
}

type SetItem[E comparable] struct {
	item E
	s    HashSet[E]
}

func (s *SetItem[E]) Remove() { s.s.Remove(s.item) }

func (s *SetItem[E]) Value() E { return s.item }
