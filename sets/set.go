package sets

import (
	"github.com/go-board/std/core"
	"github.com/go-board/std/iter"
)

var unit = struct{}{}

type HashSet[E comparable] struct{ inner map[E]core.Unit }

func New[E comparable]() HashSet[E] {
	return HashSet[E]{inner: make(map[E]core.Unit)}
}

// FromSlice returns a new empty hash set.
func FromSlice[E comparable](elements ...E) HashSet[E] {
	set := New[E]()
	for _, elem := range elements {
		set.Add(elem)
	}
	return set
}

func FromMapKeys[E comparable, V any, M ~map[E]V](m M) HashSet[E] {
	set := New[E]()
	for key := range m {
		set.Add(key)
	}
	return set
}

// FromIter create a HashSet from [iter.Seq].
func FromIter[E comparable](s iter.Seq[E]) HashSet[E] {
	set := New[E]()
	set.AddIter(s)
	return set
}

// Add adds the given key to the set.
func (self HashSet[E]) Add(key E) {
	self.inner[key] = unit
}

// AddIter adds all elements in [iter.Seq] to the set.
func (self HashSet[E]) AddIter(it iter.Seq[E]) {
	iter.ForEach(it, self.Add)
}

// Remove removes the given key from the set.
func (self HashSet[E]) Remove(key E) {
	delete(self.inner, key)
}

// RemoveIter removes all elements in [iter.Seq].
func (self HashSet[E]) RemoveIter(it iter.Seq[E]) { iter.ForEach(it, self.Remove) }

// Clear removes all keys from the set.
func (self HashSet[E]) Clear() {
	for k := range self.inner {
		delete(self.inner, k)
	}
}

// Retain keep element that match the given predicate function.
//
// Otherwise, remove from [HashSet].
func (self HashSet[E]) Retain(fn func(E) bool) {
	iter.ForEach(iter.Filter(self.Iter(), func(e E) bool { return !fn(e) }), self.Remove)
}

// Contains returns true if the given key is in the set.
func (self HashSet[E]) Contains(key E) bool {
	_, exists := self.inner[key]
	return exists
}

// ContainsAll returns true if all the given keys are in the set.
func (self HashSet[E]) ContainsAll(keys iter.Seq[E]) bool {
	return iter.All(keys, self.Contains)
}

// ContainsAny returns true if any of the given keys are in the set.
func (self HashSet[E]) ContainsAny(keys iter.Seq[E]) bool {
	return iter.Any(keys, self.Contains)
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
	m := make(map[E]struct{})
	for k := range self.inner {
		m[k] = struct{}{}
	}
	return m
}

// Clone returns a copy of the set.
func (self HashSet[E]) Clone() HashSet[E] {
	return self.DeepCloneBy(func(t E) E { return t })
}

// DeepCloneBy returns a copy of the set and clone each element use given clone func.
func (self HashSet[E]) DeepCloneBy(clone func(E) E) HashSet[E] {
	return FromIter(iter.Map(self.Iter(), clone))
}

// SupersetOf returns true if the given set is a superset of this set.
func (self HashSet[E]) SupersetOf(other HashSet[E]) bool {
	return iter.All(other.Iter(), self.Contains)
}

func (self HashSet[E]) SupersetOfIter(it iter.Seq[E]) bool {
	return iter.All(it, self.Contains)
}

// SubsetOf returns true if the given set is a subset of this set.
func (self HashSet[E]) SubsetOf(other HashSet[E]) bool {
	return iter.All(self.Iter(), other.Contains)
}

func (self HashSet[E]) SubsetOfIter(it iter.Seq[E]) bool {
	return iter.All(self.Iter(), FromIter(it).Contains)
}

// Union returns a new set containing all the elements that are in either set.
func (self HashSet[E]) Union(other HashSet[E]) HashSet[E] {
	union := self.Clone()
	union.AddIter(other.Iter())
	return union
}

func (self HashSet[E]) UnionIter(it iter.Seq[E]) HashSet[E] {
	union := self.Clone()
	union.AddIter(it)
	return union
}

// Intersection returns a new set containing all the elements that are in both sets.
func (self HashSet[E]) Intersection(other HashSet[E]) HashSet[E] {
	return FromIter(iter.Filter(self.Iter(), other.Contains))
}

func (self HashSet[E]) IntersectionIter(it iter.Seq[E]) HashSet[E] {
	return self.Intersection(FromIter(it))
}

// Difference returns a new set containing all the elements that are in this set but not in the other set.
func (self HashSet[E]) Difference(other HashSet[E]) HashSet[E] {
	return FromIter(iter.Filter(self.Iter(), func(e E) bool { return !other.Contains(e) }))
}

func (self HashSet[E]) DifferenceIter(it iter.Seq[E]) HashSet[E] {
	return self.Difference(FromIter(it))
}

// SymmetricDifference returns a new set containing all the elements that are in this set or the other set but not in both.
func (self HashSet[E]) SymmetricDifference(other HashSet[E]) HashSet[E] {
	return self.Union(other).Difference(self.Intersection(other))
}

func (self HashSet[E]) SymmetricDifferenceIter(it iter.Seq[E]) HashSet[E] {
	return self.SymmetricDifference(FromIter(it))
}

// Equal returns true if the given set is equal to this set.
func (self HashSet[E]) Equal(other HashSet[E]) bool {
	if self.Size() != other.Size() {
		return false
	}
	return self.SupersetOf(other) && self.SubsetOf(other)
}

func (self HashSet[E]) EqualIter(it iter.Seq[E]) bool {
	return self.Equal(FromIter(it))
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
			if !yield(&SetItem[E]{item: key, set: self}) {
				break
			}
		}
	}
}

type SetItem[E comparable] struct {
	item E
	set  HashSet[E]
}

func (s *SetItem[E]) Value() E { return s.item }
func (s *SetItem[E]) Remove()  { s.set.Remove(s.item) }
