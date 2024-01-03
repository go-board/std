package sets

import (
	"testing"

	"github.com/frankban/quicktest"
	"github.com/go-board/std/iter"
)

func seq[E any](elems ...E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, e := range elems {
			if !yield(e) {
				break
			}
		}
	}
}

type item struct{ key int }

func (i item) Clone() item { return item{key: i.key} }

func TestHashSet_Add(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int]()
	s.AddIter(seq(1, 2))
	a.Assert(s.Contains(1), quicktest.IsTrue)
	a.Assert(s.Contains(2), quicktest.IsTrue)
	a.Assert(s.Contains(3), quicktest.IsFalse)
}

func TestHashSet_Remove(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int](1, 2)
	s.Remove(1)
	a.Assert(s.Contains(1), quicktest.IsFalse)
	a.Assert(s.Contains(2), quicktest.IsTrue)
	a.Assert(s.Size(), quicktest.Equals, 1)
}

func TestHashSet_Clear(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int](1, 2)
	s.Clear()
	a.Assert(s.Contains(1), quicktest.IsFalse)
	a.Assert(s.Contains(2), quicktest.IsFalse)
	a.Assert(s.Size(), quicktest.Equals, 0)
}

func TestHashSet_Contains(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int](1, 2)
	a.Assert(s.Contains(1), quicktest.IsTrue)
	a.Assert(s.Contains(2), quicktest.IsTrue)
	a.Assert(s.Contains(3), quicktest.IsFalse)
}

func TestHashSet_ContainsAll(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int](1, 2, 3)
	a.Assert(s.ContainsAll(seq(1, 2, 3)), quicktest.IsTrue)
	a.Assert(s.ContainsAll(seq(1, 2, 4)), quicktest.IsFalse)
}

func TestHashSet_ContainsAny(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int](1, 2, 3)
	a.Assert(true, quicktest.Equals, s.ContainsAny(seq(1, 2, 4)))
	a.Assert(true, quicktest.Equals, s.ContainsAny(seq(1, 2, 3)))
	a.Assert(false, quicktest.Equals, s.ContainsAny(seq(5, 6, 7)))
}

func TestHashSet_Size(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int](1, 2, 3)
	a.Assert(s.Size(), quicktest.Equals, 3)
	s.Add(5)
	a.Assert(s.Size(), quicktest.Equals, 4)
	s.Remove(3)
	a.Assert(s.Size(), quicktest.Equals, 3)
}

func TestHashSet_IsEmpty(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int]()
	a.Assert(s.IsEmpty(), quicktest.IsTrue)
	s.Add(1)
	a.Assert(s.IsEmpty(), quicktest.IsFalse)
}

func TestHashSet_Clone(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int](1, 2, 3)
	s2 := s1.Clone()
	a.Assert(s1.Equal(s2), quicktest.IsTrue)
	s2.Add(4)
	a.Assert(s1.Equal(s2), quicktest.IsFalse)
}

func TestHashSet_DeepCloneBy(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int](1, 2, 3)
	s2 := s1.DeepCloneBy(func(i int) int { return i })
	a.Assert(s1.Equal(s2), quicktest.IsTrue)
	s2.Add(4)
	a.Assert(s1.Equal(s2), quicktest.IsFalse)
}

func TestHashSet_SupersetOf(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int]()
	s1.AddIter(seq[int](1, 2, 3))
	s2 := FromSlice[int]()
	a.Assert(s1.SupersetOf(s2), quicktest.IsTrue)
	s2.AddIter(seq(1, 2, 3))
	a.Assert(s1.SupersetOf(s2), quicktest.IsTrue)
	s2.Add(4)
	a.Assert(s1.SupersetOf(s2), quicktest.IsFalse)
}

func TestHashSet_SubsetOf(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int](1, 2, 3)
	s2 := FromSlice[int]()
	a.Assert(s2.SubsetOf(s1), quicktest.IsTrue)
	s2.AddIter(seq(1, 2, 3))
	a.Assert(s2.SubsetOf(s1), quicktest.IsTrue)
	s2.Add(4)
	a.Assert(s2.SubsetOf(s1), quicktest.IsFalse)
}

func TestHashSet_Union(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int](1, 2, 3)
	s2 := FromSlice[int](1, 4, 5)
	s3 := s1.Union(s2)
	a.Assert(s3.ContainsAll(seq(1, 2, 3, 4, 5)), quicktest.IsTrue)
	a.Assert(s3.Size(), quicktest.Equals, 5)
}

func TestHashSet_Intersection(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := FromSlice[int]()
	s2.Add(1)
	s2.Add(4)
	s2.Add(5)
	s3 := s1.Intersection(s2)
	a.Assert(s3.ContainsAll(seq(1)), quicktest.IsTrue)
	a.Assert(s3.Size(), quicktest.Equals, 1)
}

func TestHashSet_Difference(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := FromSlice[int]()
	s2.Add(1)
	s2.Add(4)
	s2.Add(5)
	s3 := s1.Difference(s2)
	a.Assert(s3.ContainsAll(seq(2, 3)), quicktest.IsTrue)
	a.Assert(s3.Size(), quicktest.Equals, 2)
}

func TestHashSet_SymmetricDifference(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int](1, 2, 3)
	s2 := FromSlice[int](1, 4, 5)
	s3 := s1.SymmetricDifference(s2)
	a.Assert(s3.ContainsAll(seq(2, 3, 4, 5)), quicktest.IsTrue)
	a.Assert(s3.Size(), quicktest.Equals, 4)
}

func TestHashSet_Equal(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int](1, 2, 3)
	s2 := FromSlice[int](1, 2, 3)
	a.Assert(s1.Equal(s2), quicktest.IsTrue)
	s2.Add(4)
	a.Assert(s1.Equal(s2), quicktest.IsFalse)
}

func TestHashSet_Iter(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice(1, 2, 3, 4, 5)

	a.Assert(iter.Size(s1.Iter()), quicktest.Equals, 5)
}

func TestHashSet_IterMut(t *testing.T) {
	s := FromSlice[int](1, 2, 3, 4)
	iter.ForEach(s.IterMut(), func(s *SetItem[int]) {
		if s.Value()%2 == 0 {
			s.Remove()
		}
	})
	x := s.ToMap()
	t.Logf("result is %+v\n", x)
}
