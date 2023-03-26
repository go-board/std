package btree

import (
	"github.com/go-board/std/clone"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/iterator/stream"
	"github.com/go-board/std/optional"
	"github.com/tidwall/btree"
)

// TreeSet is a set based on a B-Tree.
type TreeSet[TElement any] struct {
	inner *btree.BTreeG[TElement]
}

var (
	_ clone.Cloneable[*TreeSet[any]] = (*TreeSet[any])(nil)
	_ iterator.Iterable[any]         = (*TreeSet[any])(nil)
)

// NewTreeSet creates a new TreeSet.
func NewTreeSet[TElement any](less func(TElement, TElement) bool) *TreeSet[TElement] {
	return &TreeSet[TElement]{inner: btree.NewBTreeG(less)}
}

// Add adds elements to the set.
func (self *TreeSet[TElement]) Add(elements ...TElement) {
	for _, element := range elements {
		self.inner.Set(element)
	}
}

// Remove removes elements from the set.
func (self *TreeSet[TElement]) Remove(elements ...TElement) {
	for _, element := range elements {
		self.inner.Delete(element)
	}
}

// Reversed returns a reversed view of the set.
func (self *TreeSet[TElement]) Reversed() *TreeSet[TElement] {
	newSet := NewTreeSet(func(t1, t2 TElement) bool { return !self.inner.Less(t1, t2) })
	stream.FromIterable[TElement](self).ForEach(func(element TElement) {
		newSet.Add(element)
	})
	return newSet
}

// Contains returns true if the set contains the element.
func (self *TreeSet[TElement]) Contains(element TElement) bool {
	_, ok := self.inner.Get(element)
	return ok
}

// Range returns an iterator over the set.
func (self *TreeSet[TElement]) Range(start, end TElement) *TreeSet[TElement] {
	if self.inner.Less(start, end) {
		newSet := NewTreeSet(self.inner.Less)
		self.inner.Ascend(start, func(item TElement) bool {
			newSet.Add(item)
			return newSet.inner.Less(item, end)
		})
	}
	return NewTreeSet(self.inner.Less)
}

// First returns the first element of the set.
func (self *TreeSet[TElement]) First() optional.Optional[TElement] {
	return optional.FromPair(self.inner.Min())
}

// Last returns the last element of the set.
func (self *TreeSet[TElement]) Last() optional.Optional[TElement] {
	return optional.FromPair(self.inner.Max())
}

// PopFirst removes and returns the first element of the set.
func (self *TreeSet[TElement]) PopFirst() optional.Optional[TElement] {
	return optional.FromPair(self.inner.PopMin())
}

// PopLast removes and returns the last element of the set.
func (self *TreeSet[TElement]) PopLast() optional.Optional[TElement] {
	return optional.FromPair(self.inner.PopMax())
}

// Clone returns a copy of the set.
func (self *TreeSet[TElement]) Clone() *TreeSet[TElement] {
	return &TreeSet[TElement]{inner: self.inner.Copy()}
}

// Intersection returns the intersection of the two sets.
func (self *TreeSet[TElement]) Intersection(o *TreeSet[TElement]) *TreeSet[TElement] {
	newSet := NewTreeSet(self.inner.Less)
	stream.FromIterable[TElement](o).ForEach(func(t TElement) {
		if self.Contains(t) {
			newSet.Add(t)
		}
	})
	return newSet
}

// Difference returns the difference of the two sets.
func (self *TreeSet[TElement]) Difference(o *TreeSet[TElement]) *TreeSet[TElement] {
	cloned := self.Clone()
	stream.FromIterable[TElement](self).ForEach(func(t TElement) {
		if o.Contains(t) {
			cloned.Remove(t)
		}
	})
	return cloned
}

// Union returns the union of the two sets.
func (self *TreeSet[TElement]) Union(o *TreeSet[TElement]) {
	self.Add(stream.FromIterable[TElement](o).Collect()...)
}

// SubsetOf returns true if the set is a subset of the other set.
func (self *TreeSet[TElement]) SubsetOf(o *TreeSet[TElement]) bool {
	return stream.FromIterable[TElement](self).All(func(t TElement) bool {
		return o.Contains(t)
	})
}

// SupersetOf returns true if the set is a superset of the other set.
func (self *TreeSet[TElement]) SupersetOf(o *TreeSet[TElement]) bool {
	return stream.FromIterable[TElement](o).All(func(t TElement) bool {
		return self.Contains(t)
	})
}

// Size returns the number of elements in the set.
func (self *TreeSet[TElement]) Size() uint {
	return uint(self.inner.Len())
}

// Iter returns an iterator over the set.
func (self *TreeSet[TElement]) Iter() iterator.Iterator[TElement] {
	return &treeSetIter[TElement]{iter: self.inner.Iter()}
}

type treeSetIter[TElement any] struct {
	iter btree.IterG[TElement]
}

var _ iterator.Iterator[any] = (*treeSetIter[any])(nil)

func (self *treeSetIter[TElement]) Next() optional.Optional[TElement] {
	if self.iter.Next() {
		return optional.Some(self.iter.Item())
	}
	return optional.None[TElement]()
}
