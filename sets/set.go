package sets

import "github.com/go-board/std/clone"

type HashSet[TKey comparable] struct {
	inner map[TKey]struct{}
}

// NewHashSet returns a new empty hash set.
func NewHashSet[TKey comparable](elements ...TKey) HashSet[TKey] {
	inner := make(map[TKey]struct{})
	for _, element := range elements {
		inner[element] = struct{}{}
	}
	return HashSet[TKey]{inner: inner}
}

// Add adds the given key to the set.
func (self HashSet[TKey]) Add(key TKey) {
	self.inner[key] = struct{}{}
}

// AddAll adds all the given keys to the set.
func (self HashSet[TKey]) AddAll(keys []TKey) {
	for _, k := range keys {
		self.Add(k)
	}
}

// Remove removes the given key from the set.
func (self HashSet[TKey]) Remove(key TKey) {
	self.RemoveBy(func(k TKey) bool { return k == key })
}

// RemoveBy remove keys from the set if the given predicate returns true.
func (self HashSet[TKey]) RemoveBy(predicate func(TKey) bool) {
	for k := range self.inner {
		if predicate(k) {
			delete(self.inner, k)
		}
	}
}

// Clear removes all keys from the set.
func (self HashSet[TKey]) Clear() {
	for k := range self.inner {
		delete(self.inner, k)
	}
}

func (self HashSet[TKey]) ForEach(fn func(TKey)) {
	for k := range self.inner {
		fn(k)
	}
}

// Contains returns true if the given key is in the set.
func (self HashSet[TKey]) Contains(key TKey) bool {
	_, exists := self.inner[key]
	return exists
}

// ContainsAll returns true if all of the given keys are in the set.
func (self HashSet[TKey]) ContainsAll(keys []TKey) bool {
	for _, key := range keys {
		if !self.Contains(key) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if any of the given keys are in the set.
func (self HashSet[TKey]) ContainsAny(keys []TKey) bool {
	for _, key := range keys {
		if self.Contains(key) {
			return true
		}
	}
	return false
}

// Size returns the number of elements in the set.
func (self HashSet[TKey]) Size() int {
	return len(self.inner)
}

// IsEmpty returns true if the set is empty.
func (self HashSet[TKey]) IsEmpty() bool {
	return self.Size() == 0
}

// ToSlice returns a slice containing all of the elements in the set.
func (self HashSet[TKey]) ToSlice() []TKey {
	keys := make([]TKey, self.Size())
	i := 0
	for key := range self.inner {
		keys[i] = key
		i++
	}
	return keys
}

// ToMap returns a hash map containing all of the elements in the set.
func (self HashSet[TKey]) ToMap() map[TKey]struct{} {
	// todo: deep clone? or shallow clone?
	result := make(map[TKey]struct{})
	for key := range self.inner {
		result[key] = struct{}{}
	}
	return result
}

// Clone returns a copy of the set.
func (self HashSet[TKey]) Clone() HashSet[TKey] {
	other := NewHashSet[TKey]()
	for key := range self.inner {
		other.Add(key)
	}
	return other
}

// DeepCloneBy returns a copy of the set and clone each element use given clone func.
func (self HashSet[TKey]) DeepCloneBy(clone func(TKey) TKey) HashSet[TKey] {
	other := NewHashSet[TKey]()
	for key := range self.inner {
		other.Add(clone(key))
	}
	return other
}

// SupersetOf returns true if the given set is a superset of this set.
func (self HashSet[TKey]) SupersetOf(other HashSet[TKey]) bool {
	for key := range other.inner {
		if !self.Contains(key) {
			return false
		}
	}
	return true
}

// SubsetOf returns true if the given set is a subset of this set.
func (self HashSet[TKey]) SubsetOf(other HashSet[TKey]) bool {
	for key := range self.inner {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// Union returns a new set containing all of the elements that are in either set.
func (self HashSet[TKey]) Union(other HashSet[TKey]) HashSet[TKey] {
	union := NewHashSet[TKey]()
	for k := range self.inner {
		union.Add(k)
	}
	for key := range other.inner {
		union.Add(key)
	}
	return union
}

// Intersection returns a new set containing all of the elements that are in both sets.
func (self HashSet[TKey]) Intersection(other HashSet[TKey]) HashSet[TKey] {
	intersection := NewHashSet[TKey]()
	for key := range self.inner {
		if other.Contains(key) {
			intersection.Add(key)
		}
	}
	return intersection
}

// Difference returns a new set containing all of the elements that are in this set but not in the other set.
func (self HashSet[TKey]) Difference(other HashSet[TKey]) HashSet[TKey] {
	diff := NewHashSet[TKey]()
	for key := range self.inner {
		if !other.Contains(key) {
			diff.Add(key)
		}
	}
	return diff
}

// SymmetricDifference returns a new set containing all of the elements that are in this set or the other set but not in both.
func (self HashSet[TKey]) SymmetricDifference(other HashSet[TKey]) HashSet[TKey] {
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
func (self HashSet[TKey]) Equal(other HashSet[TKey]) bool {
	if other.inner == nil || self.inner == nil {
		return false
	}
	if self.Size() != other.Size() {
		return false
	}
	for key := range other.inner {
		if !self.Contains(key) {
			return false
		}
	}
	return true
}

// DeepClone returns a copy of the set and clone each element.
func DeepClone[TKey interface {
	comparable
	clone.Cloneable[TKey]
}](
	set HashSet[TKey],
) HashSet[TKey] {
	return set.DeepCloneBy(func(t TKey) TKey { return t.Clone() })
}

// Map returns a new set containing all of the elements that are in this set and the given map.
func Map[TKey, TNewKey comparable](
	set HashSet[TKey], fn func(TKey) TNewKey,
) HashSet[TNewKey] {
	newSet := NewHashSet[TNewKey]()
	for key := range set.inner {
		newSet.Add(fn(key))
	}
	return newSet
}
