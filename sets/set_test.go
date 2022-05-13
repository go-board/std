package sets

import (
	"testing"

	"github.com/frankban/quicktest"
)

type item struct{ key int }

func (i item) Clone() item { return item{key: i.key} }

func TestHashSet_Add(t *testing.T) {
	a := quicktest.New(t)
	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	a.Assert(true, quicktest.Equals, s.Contains(1))
	a.Assert(true, quicktest.Equals, s.Contains(2))
	a.Assert(false, quicktest.Equals, s.Contains(3))
}

func TestHashSet_AddAll(t *testing.T) {
	a := quicktest.New(t)
	s := NewHashSet[int]()
	s.AddAll([]int{1, 2})
	a.Assert(true, quicktest.Equals, s.Contains(1))
	a.Assert(true, quicktest.Equals, s.Contains(2))
	a.Assert(false, quicktest.Equals, s.Contains(3))
	a.Assert(2, quicktest.Equals, s.Size())
}

func TestHashSet_Remove(t *testing.T) {
	a := quicktest.New(t)
	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	s.Remove(1)
	a.Assert(false, quicktest.Equals, s.Contains(1))
	a.Assert(true, quicktest.Equals, s.Contains(2))
	a.Assert(1, quicktest.Equals, s.Size())
}

func TestHashSet_RemoveBy(t *testing.T) {
	a := quicktest.New(t)
	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.RemoveBy(func(i int) bool { return i == 1 })
	a.Assert(false, quicktest.Equals, s.Contains(1))
	a.Assert(true, quicktest.Equals, s.Contains(2))
	a.Assert(true, quicktest.Equals, s.Contains(3))
	a.Assert(2, quicktest.Equals, s.Size())
}

func TestHashSet_Clear(t *testing.T) {
	a := quicktest.New(t)
	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	s.Clear()
	a.Assert(false, quicktest.Equals, s.Contains(1))
	a.Assert(false, quicktest.Equals, s.Contains(2))
	a.Assert(0, quicktest.Equals, s.Size())
}

func TestHashSet_ForEach(t *testing.T) {
	a := quicktest.New(t)
	s := NewHashSet[item]()
	s.Add(item{key: 1})
	s.Add(item{key: 2})
	s.Add(item{key: 3})
	s.ForEach(func(i item) {
		a.Assert(i.key, quicktest.Equals, i.key)
	})
}

func TestHashSet_Contains(t *testing.T) {
	a := quicktest.New(t)
	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	a.Assert(true, quicktest.Equals, s.Contains(1))
	a.Assert(true, quicktest.Equals, s.Contains(2))
	a.Assert(false, quicktest.Equals, s.Contains(3))
}

func TestHashSet_ContainsAll(t *testing.T) {
	a := quicktest.New(t)
	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	a.Assert(true, quicktest.Equals, s.ContainsAll([]int{1, 2, 3}))
	a.Assert(false, quicktest.Equals, s.ContainsAll([]int{1, 2, 4}))
}

func TestHashSet_ContainsAny(t *testing.T) {
	a := quicktest.New(t)
	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	a.Assert(true, quicktest.Equals, s.ContainsAny([]int{1, 2, 4}))
	a.Assert(true, quicktest.Equals, s.ContainsAny([]int{1, 2, 3}))
	a.Assert(false, quicktest.Equals, s.ContainsAny([]int{5, 6, 7}))
}

func TestHashSet_Size(t *testing.T) {
	a := quicktest.New(t)
	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	a.Assert(3, quicktest.Equals, s.Size())
	s.Add(5)
	a.Assert(4, quicktest.Equals, s.Size())
	s.Remove(3)
	a.Assert(3, quicktest.Equals, s.Size())
}

func TestHashSet_IsEmpty(t *testing.T) {
	a := quicktest.New(t)
	s := NewHashSet[int]()
	a.Assert(true, quicktest.Equals, s.IsEmpty())
	s.Add(1)
	a.Assert(false, quicktest.Equals, s.IsEmpty())
}

func TestHashSet_ToSlice(t *testing.T) {
	a := quicktest.New(t)
	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	a.Assert([]int{1, 2, 3}, quicktest.DeepEquals, s.ToSlice())
}

func TestHashSet_Equals(t *testing.T) {
	a := quicktest.New(t)
	s1 := NewHashSet[int]()
	s2 := NewHashSet[int]()
	a.Assert(true, quicktest.Equals, s1.Equals(s2))
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	a.Assert(true, quicktest.Equals, s1.Equals(s2))
	s2.Add(4)
	a.Assert(false, quicktest.Equals, s1.Equals(s2))
}

func TestHashSet_Clone(t *testing.T) {
	a := quicktest.New(t)
	s1 := NewHashSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := s1.Clone()
	a.Assert(true, quicktest.Equals, s1.Equals(s2))
	s2.Add(4)
	a.Assert(false, quicktest.Equals, s1.Equals(s2))
}

func TestHashSet_DeepCloneBy(t *testing.T) {
	a := quicktest.New(t)
	s1 := NewHashSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := s1.DeepCloneBy(func(i int) int { return i })
	a.Assert(true, quicktest.Equals, s1.Equals(s2))
	s2.Add(4)
	a.Assert(false, quicktest.Equals, s1.Equals(s2))
}

func TestHashSet_SupersetOf(t *testing.T) {
	a := quicktest.New(t)
	s1 := NewHashSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := NewHashSet[int]()
	a.Assert(true, quicktest.Equals, s1.SupersetOf(s2))
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	a.Assert(true, quicktest.Equals, s1.SupersetOf(s2))
	s2.Add(4)
	a.Assert(false, quicktest.Equals, s1.SupersetOf(s2))
}

func TestHashSet_SubsetOf(t *testing.T) {
	a := quicktest.New(t)
	s1 := NewHashSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := NewHashSet[int]()
	a.Assert(true, quicktest.Equals, s2.SubsetOf(s1))
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	a.Assert(true, quicktest.Equals, s2.SubsetOf(s1))
	s2.Add(4)
	a.Assert(false, quicktest.Equals, s2.SubsetOf(s1))
}

func TestHashSet_Union(t *testing.T) {
	a := quicktest.New(t)
	s1 := NewHashSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := NewHashSet[int]()
	s2.Add(1)
	s2.Add(4)
	s2.Add(5)
	s3 := s1.Union(s2)
	a.Assert(true, quicktest.Equals, s3.ContainsAll([]int{1, 2, 3, 4, 5}))
}

func TestHashSet_Intersection(t *testing.T) {
	a := quicktest.New(t)
	s1 := NewHashSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := NewHashSet[int]()
	s2.Add(1)
	s2.Add(4)
	s2.Add(5)
	s3 := s1.Intersection(s2)
	a.Assert(true, quicktest.Equals, s3.ContainsAll([]int{1}))
	a.Assert(1, quicktest.Equals, s3.Size())
}

func TestHashSet_Difference(t *testing.T) {
	a := quicktest.New(t)
	s1 := NewHashSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := NewHashSet[int]()
	s2.Add(1)
	s2.Add(4)
	s2.Add(5)
	s3 := s1.Difference(s2)
	a.Assert(true, quicktest.Equals, s3.ContainsAll([]int{2, 3}))
	a.Assert(2, quicktest.Equals, s3.Size())
}

func TestHashSet_SymmetricDifference(t *testing.T) {
	a := quicktest.New(t)
	s1 := NewHashSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := NewHashSet[int]()
	s2.Add(1)
	s2.Add(4)
	s2.Add(5)
	s3 := s1.SymmetricDifference(s2)
	a.Assert(true, quicktest.Equals, s3.ContainsAll([]int{2, 3, 4, 5}))
	a.Assert(4, quicktest.Equals, s3.Size())
}

func TestHashSet_Equal(t *testing.T) {
	a := quicktest.New(t)
	s1 := NewHashSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := NewHashSet[int]()
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	a.Assert(true, quicktest.Equals, s1.Equal(s2))
	s2.Add(4)
	a.Assert(false, quicktest.Equals, s1.Equal(s2))
}

func TestDeepClone(t *testing.T) {
	a := quicktest.New(t)
	s1 := NewHashSet[item]()
	s1.Add(item{key: 1})
	s1.Add(item{key: 2})
	s1.Add(item{key: 3})
	s2 := DeepClone(s1)
	a.Assert(true, quicktest.Equals, s1.Equal(s2))
	s2.Add(item{key: 4})
	a.Assert(false, quicktest.Equals, s1.Equal(s2))
}

func TestMap(t *testing.T) {
	a := quicktest.New(t)
	s1 := NewHashSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := Map(s1, func(i int) int { return i * 2 })
	a.Assert(true, quicktest.Equals, s2.ContainsAll([]int{2, 4, 6}))
}
