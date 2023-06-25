package ordered

import (
	"github.com/go-board/std/clone"
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/iter"
	"github.com/go-board/std/optional"
	"github.com/tidwall/btree"
)

// Set is an ordered set based on a B-Tree.
type Set[T any] struct {
	cmp   func(T, T) int
	inner *btree.BTreeG[T]
}

var (
	_ clone.Cloneable[*Set[any]] = (*Set[any])(nil)
	_ iter.Iterable[any]         = (*Set[any])(nil)
)

// NewSet creates a new Set.
func NewSet[T any](cmp func(T, T) int) *Set[T] {
	less := func(a, b T) bool { return cmp(a, b) < 0 }
	return &Set[T]{inner: btree.NewBTreeG(less), cmp: cmp}
}

func NewOrderedSet[T cmp.Ordered]() *Set[T] {
	return NewSet(cmp.Compare[T])
}

func (self *Set[T]) Insert(element T) {
	self.inner.Set(element)
}

// InsertMany adds elements to the set.
func (self *Set[T]) InsertMany(elements ...T) {
	for _, element := range elements {
		self.inner.Set(element)
	}
}

// Remove removes elements from the set.
func (self *Set[T]) Remove(elements ...T) {
	for _, element := range elements {
		self.inner.Delete(element)
	}
}

func (self *Set[T]) Clear() {
	self.inner.Clear()
}

// Reverse returns a reversed view of the set.
func (self *Set[T]) Reverse() *Set[T] {
	newSet := NewSet(invert(self.cmp))
	iter.ForEach(self.Iter(), newSet.Insert)
	return newSet
}

// Contains returns true if the set contains the element.
func (self *Set[T]) Contains(element T) bool {
	_, ok := self.inner.Get(element)
	return ok
}

// Range returns an iter over the set.
func (self *Set[T]) Range(start, end T) *Set[T] {
	newSet := NewSet(self.cmp)
	insert := func(s *Set[T], v T) bool {
		s.Insert(v)
		return true
	}
	self.inner.Ascend(start, func(item T) bool {
		return self.inner.Less(item, end) && insert(newSet, item)
	})
	return newSet
}

// First returns the first element of the set.
func (self *Set[T]) First() optional.Optional[T] {
	return optional.FromPair(self.inner.Min())
}

// Last returns the last element of the set.
func (self *Set[T]) Last() optional.Optional[T] {
	return optional.FromPair(self.inner.Max())
}

// PopFirst removes and returns the first element of the set.
func (self *Set[T]) PopFirst() optional.Optional[T] {
	return optional.FromPair(self.inner.PopMin())
}

// PopLast removes and returns the last element of the set.
func (self *Set[T]) PopLast() optional.Optional[T] {
	return optional.FromPair(self.inner.PopMax())
}

// Clone returns a copy of the set.
func (self *Set[T]) Clone() *Set[T] {
	return &Set[T]{inner: self.inner.Copy(), cmp: self.cmp}
}

// Intersection returns the intersection of the two sets.
func (self *Set[T]) Intersection(o *Set[T]) *Set[T] {
	newSet := NewSet(self.cmp)
	iter.ForEach(self.Iter(), func(t T) {
		if o.Contains(t) {
			newSet.Insert(t)
		}
	})
	return newSet
}

// Difference returns the difference of the two sets.
func (self *Set[T]) Difference(o *Set[T]) *Set[T] {
	cloned := self.Clone()
	iter.ForEach(self.Iter(), func(t T) {
		if o.Contains(t) {
			cloned.Remove(t)
		}
	})
	return cloned
}

// Union returns the union of the two sets.
func (self *Set[T]) Union(o *Set[T]) *Set[T] {
	cloned := self.Clone()
	iter.ForEach(o.Iter(), cloned.Insert)
	return cloned
}

// SubsetOf returns true if the set is a subset of the other set.
func (self *Set[T]) SubsetOf(o *Set[T]) bool {
	return iter.AllOf(o.Iter(), o.Contains)
}

// SupersetOf returns true if the set is a superset of the other set.
func (self *Set[T]) SupersetOf(o *Set[T]) bool {
	return iter.AllOf(o.Iter(), self.Contains)
}

// Len returns the number of elements in the set.
func (self *Set[T]) Len() int {
	return self.inner.Len()
}

// Iter returns an iter over the set.
func (self *Set[T]) Iter() iter.Iter[T] {
	return &setIter[T]{iter: self.inner.Iter()}
}

type setIter[T any] struct {
	iter btree.IterG[T]
}

var _ iter.Iter[any] = (*setIter[any])(nil)

func (self *setIter[T]) Next() optional.Optional[T] {
	if self.iter.Next() {
		return optional.Some(self.iter.Item())
	}
	return optional.None[T]()
}
