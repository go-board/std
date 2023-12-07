package sets

import (
	"encoding/json"
	"testing"

	"github.com/go-board/std/iter"

	"github.com/frankban/quicktest"
)

func seq[E any, S ~[]E](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, e := range s {
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
	s.Add(1, 2)
	a.Assert(s.Contains(1), quicktest.IsTrue)
	a.Assert(s.Contains(2), quicktest.IsTrue)
	a.Assert(s.Contains(3), quicktest.IsFalse)
}

func TestHashSet_AddAll(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int]()
	s.AddAll(FromSlice(1, 2))
	a.Assert(s.Contains(1), quicktest.IsTrue)
	a.Assert(s.Contains(2), quicktest.IsTrue)
	a.Assert(s.Contains(3), quicktest.IsFalse)
	a.Assert(s.Size(), quicktest.Equals, 2)
}

func TestHashSet_Remove(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int]()
	s.Add(1)
	s.Add(2)
	s.Remove(1)
	a.Assert(s.Contains(1), quicktest.IsFalse)
	a.Assert(s.Contains(2), quicktest.IsTrue)
	a.Assert(s.Size(), quicktest.Equals, 1)
}

func TestHashSet_Clear(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int]()
	s.Add(1)
	s.Add(2)
	s.Clear()
	a.Assert(s.Contains(1), quicktest.IsFalse)
	a.Assert(s.Contains(2), quicktest.IsFalse)
	a.Assert(s.Size(), quicktest.Equals, 0)
}

func TestHashSet_Contains(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int]()
	s.Add(1)
	s.Add(2)
	a.Assert(s.Contains(1), quicktest.IsTrue)
	a.Assert(s.Contains(2), quicktest.IsTrue)
	a.Assert(s.Contains(3), quicktest.IsFalse)
}

func TestHashSet_ContainsAll(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	a.Assert(s.ContainsAll(seq([]int{1, 2, 3})), quicktest.IsTrue)
	a.Assert(s.ContainsAll(seq([]int{1, 2, 4})), quicktest.IsFalse)
}

func TestHashSet_ContainsAny(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	a.Assert(true, quicktest.Equals, s.ContainsAny(seq([]int{1, 2, 4})))
	a.Assert(true, quicktest.Equals, s.ContainsAny(seq([]int{1, 2, 3})))
	a.Assert(false, quicktest.Equals, s.ContainsAny(seq([]int{5, 6, 7})))
}

func TestHashSet_Size(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
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
	s1 := FromSlice[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := s1.Clone()
	a.Assert(s1.Equal(s2), quicktest.IsTrue)
	s2.Add(4)
	a.Assert(s1.Equal(s2), quicktest.IsFalse)
}

func TestHashSet_DeepCloneBy(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := s1.DeepCloneBy(func(i int) int { return i })
	a.Assert(s1.Equal(s2), quicktest.IsTrue)
	s2.Add(4)
	a.Assert(s1.Equal(s2), quicktest.IsFalse)
}

func TestHashSet_SupersetOf(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := FromSlice[int]()
	a.Assert(s1.SupersetOf(s2), quicktest.IsTrue)
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	a.Assert(s1.SupersetOf(s2), quicktest.IsTrue)
	s2.Add(4)
	a.Assert(s1.SupersetOf(s2), quicktest.IsFalse)
}

func TestHashSet_SubsetOf(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := FromSlice[int]()
	a.Assert(s2.SubsetOf(s1), quicktest.IsTrue)
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	a.Assert(s2.SubsetOf(s1), quicktest.IsTrue)
	s2.Add(4)
	a.Assert(s2.SubsetOf(s1), quicktest.IsFalse)
}

func TestHashSet_Union(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := FromSlice[int]()
	s2.Add(1)
	s2.Add(4)
	s2.Add(5)
	s3 := s1.Union(s2)
	a.Assert(s3.ContainsAll(seq([]int{1, 2, 3, 4, 5})), quicktest.IsTrue)
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
	a.Assert(s3.ContainsAll(seq([]int{1})), quicktest.IsTrue)
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
	a.Assert(s3.ContainsAll(seq([]int{2, 3})), quicktest.IsTrue)
	a.Assert(s3.Size(), quicktest.Equals, 2)
}

func TestHashSet_SymmetricDifference(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := FromSlice[int]()
	s2.Add(1)
	s2.Add(4)
	s2.Add(5)
	s3 := s1.SymmetricDifference(s2)
	a.Assert(s3.ContainsAll(seq([]int{2, 3, 4, 5})), quicktest.IsTrue)
	a.Assert(s3.Size(), quicktest.Equals, 4)
}

func TestHashSet_Equal(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2 := FromSlice[int]()
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	a.Assert(s1.Equal(s2), quicktest.IsTrue)
	s2.Add(4)
	a.Assert(s1.Equal(s2), quicktest.IsFalse)
}

func TestHashSet_Iter(t *testing.T) {
	a := quicktest.New(t)
	s1 := FromSlice(1, 2, 3, 4, 5)

	a.Assert(iter.Size(s1.Iter()), quicktest.Equals, 5)
}

func TestHashSet_Marshal(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice(5, 1, 4, 2, 3, 1, 2, 3)
	b, err := json.Marshal(s)
	a.Assert(err, quicktest.IsNil)
	a.Logf("%s\n ", b)
}

func TestHashSet_UnmarshalJSON(t *testing.T) {
	a := quicktest.New(t)
	s := FromSlice[int]()
	err := json.Unmarshal([]byte(`{"1":{},"2":{},"3":{},"4":{},"5":{}}`), &s)
	a.Assert(err, quicktest.IsNil)
	a.Logf("%+v\n", s)
	a.Assert(s.Equal(FromSlice(1, 2, 3, 4, 5)), quicktest.IsTrue)
}

func TestIterMut(t *testing.T) {
	s := FromSlice[int](1, 2, 3, 4)
	s.IterMut().ForEach(func(s *SetItem[int]) {
		if s.Value()%2 == 0 {
			s.Remove()
		}
	})
	x := s.ToMap()
	t.Logf("result is %+v\n", x)
}
