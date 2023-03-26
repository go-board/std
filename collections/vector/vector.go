package vector

import (
	"encoding/json"

	"github.com/go-board/std/core"
	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/slices"
)

type Vector[T any] struct{ vec []T }

// New create a [Vector]
func New[T any](elems ...T) *Vector[T] {
	return &Vector[T]{vec: elems}
}

// FromSlice create a new [Vector] from given slice
func FromSlice[T any](slice []T) *Vector[T] {
	return New(slice...)
}

// Iter implements [iterator.Iterable] and returns an [iterator.Iterator]
func (self *Vector[T]) Iter() iterator.Iterator[T] {
	return &vectorIter[T]{vec: self, index: -1, total: self.Len()}
}

type vectorIter[T any] struct {
	vec   *Vector[T]
	index int
	total int
}

func (i *vectorIter[T]) Next() optional.Optional[T] {
	i.index++
	if i.index < i.total {
		return i.vec.At(i.index)
	}
	return optional.None[T]()
}

var _ iterator.Iterator[any] = (*vectorIter[any])(nil)

// Len returns length of [Vector]
func (self *Vector[T]) Len() int { return len(self.vec) }

// Cap returns capacity of [Vector]
func (self *Vector[T]) Cap() int { return cap(self.vec) }

// Available returns how many available slots
func (self *Vector[T]) Available() int { return self.Cap() - self.Len() }

// At return element or None at specific position.
func (self *Vector[T]) At(n int) optional.Optional[T] {
	if n < 0 || n >= self.Len() {
		return optional.None[T]()
	}
	return optional.Some(self.vec[n])
}

func (self *Vector[T]) All(predicate func(T) bool) bool {
	return slices.All(self.vec, predicate)
}

func (self *Vector[T]) Any(predicate func(T) bool) bool {
	return slices.Any(self.vec, predicate)
}

func (self *Vector[T]) None(predicate func(T) bool) bool {
	return slices.None(self.vec, predicate)
}

func (self *Vector[T]) ReduceLeft(by func(T, T) T) optional.Optional[T] {
	return slices.Reduce(self.vec, by)
}

func (self *Vector[T]) ReduceRight(by func(T, T) T) optional.Optional[T] {
	return slices.ReduceRight(self.vec, by)
}

func (self *Vector[T]) Reverse() *Vector[T] {
	slices.Reverse(self.vec)
	return self
}

func (self *Vector[T]) Clone() *Vector[T] {
	self.vec = slices.Clone(self.vec)
	return self
}

func (self *Vector[T]) DeepCloneBy(by func(T) T) *Vector[T] {
	self.vec = slices.DeepCloneBy(self.vec, by)
	return self
}

func (self *Vector[T]) ForEach(predicate func(T)) {
	slices.ForEach(self.vec, predicate)
}

func (self *Vector[T]) ForEachIndexed(predicate func(elem T, index int)) {
	slices.ForEachIndexed(self.vec, predicate)
}

func (self *Vector[T]) Filter(predicate func(T) bool) *Vector[T] {
	return FromSlice(slices.Filter(self.vec, predicate))
}

func (self *Vector[T]) Map(transform func(T) T) *Vector[T] {
	return FromSlice(slices.Map(self.vec, transform))
}

func (self *Vector[T]) MaxBy(less func(T, T) bool) optional.Optional[T] {
	return slices.MaxBy(self.vec, less)
}

func (self *Vector[T]) MinBy(less func(T, T) bool) optional.Optional[T] {
	return slices.MinBy(self.vec, less)
}

func (self *Vector[T]) Raw() []T {
	return self.vec
}

func (self *Vector[T]) Append(elem T) *Vector[T] {
	self.vec = append(self.vec, elem)
	return self
}

func (self *Vector[T]) AppendAll(elems ...T) *Vector[T] {
	self.vec = append(self.vec, elems...)
	return self
}

func (self *Vector[T]) Prepend(elem T) *Vector[T] {
	self.vec = append([]T{elem}, self.vec...)
	return self
}

func (self *Vector[T]) Extend(vec *Vector[T]) *Vector[T] {
	self.vec = append(self.vec, vec.vec...)
	return self
}

func (self *Vector[T]) EqualBy(vec *Vector[T], by func(T, T) bool) bool {
	return slices.EqualBy(self.vec, vec.vec, by)
}

func (self *Vector[T]) ContainsBy(e T, by func(T, T) bool) bool {
	return slices.ContainsBy(self.vec, e, by)
}

func (self *Vector[T]) IndexBy(e T, by func(T, T) bool) optional.Optional[int] {
	return slices.IndexBy(self.vec, e, by)
}

func (self *Vector[T]) LastIndexBy(e T, by func(T, T) bool) optional.Optional[int] {
	return slices.LastIndexBy(self.vec, e, by)
}

func (self *Vector[T]) SortBy(less func(T, T) bool) *Vector[T] {
	slices.SortBy(self.vec, less)
	return self
}

func (self *Vector[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.vec)
}

func (self *Vector[T]) UnmarshalJSON(v []byte) error {
	return json.Unmarshal(v, &self.vec)
}

var (
	_ json.Marshaler   = (*Vector[any])(nil)
	_ json.Unmarshaler = (*Vector[any])(nil)
)

// ComparableVector wraps Vector where element type is comparable
//
// This is specific version of Vector, aka: specialization.
type ComparableVector[T comparable] struct{ *Vector[T] }

// NewComparable create a new [ComparableVector] which element is [comparable]
func NewComparable[T comparable](elems ...T) *ComparableVector[T] {
	return &ComparableVector[T]{Vector: New(elems...)}
}

func FromComparable[T comparable](slice []T) *ComparableVector[T] {
	return NewComparable(slice...)
}

func (self *ComparableVector[T]) Contains(e T) bool {
	return slices.Contains(self.vec, e)
}

func (self *ComparableVector[T]) Equal(vec *ComparableVector[T]) bool {
	return slices.Equal(self.vec, vec.vec)
}

func (self *ComparableVector[T]) Index(e T) optional.Optional[int] {
	return slices.Index(self.vec, e)
}

func (self *ComparableVector[T]) LastIndex(e T) optional.Optional[int] {
	return slices.LastIndex(self.vec, e)
}

func (self *ComparableVector[T]) Uniq() *ComparableVector[T] {
	m := make(map[T]struct{}, len(self.vec))
	vec := make([]T, 0, len(self.vec))

	for _, v := range self.vec {
		if _, ok := m[v]; !ok {
			vec = append(vec, v)
			m[v] = struct{}{}
		}
	}
	return NewComparable(vec...)
}

// OrderedVector consist of ordered element
type OrderedVector[T core.Ordered] struct{ *Vector[T] }

func NewOrdered[T core.Ordered](elems ...T) *OrderedVector[T] {
	return &OrderedVector[T]{Vector: New(elems...)}
}

func FromOrdered[T core.Ordered](slice []T) *OrderedVector[T] {
	return NewOrdered(slice...)
}

func (self *OrderedVector[T]) Sort() *OrderedVector[T] {
	slices.Sort(self.vec)
	return self
}

func (self *OrderedVector[T]) Sorted() bool {
	return slices.IsSorted(self.vec)
}

func (self *OrderedVector[T]) Max() optional.Optional[T] {
	return slices.Max(self.vec)
}

func (self *OrderedVector[T]) Min() optional.Optional[T] {
	return slices.Min(self.vec)
}
