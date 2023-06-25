package slices_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/slices"
)

type item struct {
	Value int
	Name  string
	Tags  []string
}

func (i item) Compare(o item) int {
	return cmp.Compare(i.Value, o.Value)
}

func TestSort(t *testing.T) {
	t.Run("sort_by", func(t *testing.T) {
		t.Run("comparable", func(t *testing.T) {
			a := qt.New(t)
			slice := []int{1, 2, 4, 3, 9, 7, 15, 11}
			slices.SortBy(slice, cmp.Compare[int])
			a.Assert(slice, qt.DeepEquals, []int{1, 2, 3, 4, 7, 9, 11, 15})
		})
		t.Run("any", func(t *testing.T) {
			a := qt.New(t)
			slice := []item{
				{Value: 3, Name: "Jack", Tags: []string{"b", "c"}},
				{Value: 2, Name: "Jane", Tags: []string{"a", "c"}},
				{Value: 4, Name: "Jill", Tags: []string{"b", "c"}},
				{Value: 1, Name: "John", Tags: []string{"a", "b"}},
			}
			slices.SortBy(slice, item.Compare)
			a.Assert(slice, qt.DeepEquals, []item{
				{Value: 1, Name: "John", Tags: []string{"a", "b"}},
				{Value: 2, Name: "Jane", Tags: []string{"a", "c"}},
				{Value: 3, Name: "Jack", Tags: []string{"b", "c"}},
				{Value: 4, Name: "Jill", Tags: []string{"b", "c"}},
			})
		})
	})
	t.Run("sort", func(t *testing.T) {
		a := qt.New(t)
		slice := []int{1, 15, 3, 2, 11, 4, 9, 7}
		slices.Sort(slice)
		a.Assert(slice, qt.DeepEquals, []int{1, 2, 3, 4, 7, 9, 11, 15})
	})
}

func TestIsSortedBy(t *testing.T) {
	a := qt.New(t)
	a.Run("sorted_by", func(c *qt.C) {
		sorted := slices.IsSortedBy([]item{
			{Value: 1, Name: "John", Tags: []string{"a", "b"}},
			{Value: 2, Name: "Jane", Tags: []string{"a", "c"}},
			{Value: 3, Name: "Jack", Tags: []string{"b", "c"}},
			{Value: 4, Name: "Jill", Tags: []string{"b", "c"}},
		}, item.Compare)
		c.Assert(sorted, qt.IsTrue)
	})
	a.Run("sorted", func(c *qt.C) {
		sorted := slices.IsSorted([]int{1, 2, 3, 4, 7, 9, 11, 15})
		c.Assert(sorted, qt.IsTrue)
	})
}

func TestMap(t *testing.T) {
	a := qt.New(t)
	a.Run("map", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		mapped := slices.Map(slice, func(i int) int { return i * 2 })
		c.Assert(mapped, qt.DeepEquals, []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20})
	})
	a.Run("complex map", func(c *qt.C) {
		slice := []item{
			{Value: 1, Name: "John", Tags: []string{"a", "b"}},
			{Value: 2, Name: "Jane", Tags: []string{"a", "c"}},
			{Value: 3, Name: "Jack", Tags: []string{"b", "c"}},
			{Value: 4, Name: "Jill", Tags: []string{"b", "c"}},
		}
		mapped := slices.Map(slice, func(i item) []string { return i.Tags })
		c.Assert(mapped, qt.DeepEquals, [][]string{
			{"a", "b"},
			{"a", "c"},
			{"b", "c"},
			{"b", "c"},
		})
	})
	a.Run("indexMap", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		v := slices.MapIndexed(slice, func(i int, v int) int { return v * 2 })
		c.Assert(v, qt.DeepEquals, []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18})
	})
}

func TestFilter(t *testing.T) {
	a := qt.New(t)
	a.Run("filter", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		filtered := slices.Filter(slice, func(i int) bool { return i%2 == 0 })
		c.Assert(filtered, qt.DeepEquals, []int{2, 4, 6, 8, 10})
	})
	a.Run("complex filter", func(c *qt.C) {
		slice := []item{
			{Value: 1, Name: "John", Tags: []string{"a", "b"}},
			{Value: 2, Name: "Jane", Tags: []string{"a", "c"}},
			{Value: 3, Name: "Jack", Tags: []string{"b", "c"}},
			{Value: 4, Name: "Jill", Tags: []string{"b", "c"}},
		}
		filtered := slices.Filter(slice, func(i item) bool { return i.Value%2 == 1 })
		c.Assert(filtered, qt.DeepEquals, []item{
			{Value: 1, Name: "John", Tags: []string{"a", "b"}},
			{Value: 3, Name: "Jack", Tags: []string{"b", "c"}},
		})
	})
	a.Run("filterIndexed", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		filtered := slices.FilterIndexed(slice, func(i int, v int) bool { return v%2 == 0 })
		c.Assert(filtered, qt.DeepEquals, []int{1, 3, 5, 7, 9})
	})
}

func TestFold(t *testing.T) {
	a := qt.New(t)
	a.Run("fold", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		folded := slices.Fold(slice, 0, func(acc, i int) int { return acc + i })
		c.Assert(folded, qt.Equals, 55)
	})
	a.Run("complex fold", func(c *qt.C) {
		slice := []item{
			{Value: 1, Name: "John", Tags: []string{"a", "b"}},
			{Value: 2, Name: "Jane", Tags: []string{"a", "c"}},
			{Value: 3, Name: "Jack", Tags: []string{"b", "c"}},
			{Value: 4, Name: "Jill", Tags: []string{"b", "c"}},
		}
		foldedValue := slices.Fold(slice, 0, func(acc int, i item) int { return acc + i.Value })
		foldedName := slices.Fold(slice, "", func(acc string, i item) string { return acc + " " + i.Name })
		foldedTags := slices.Fold(slice, []string{}, func(acc []string, i item) []string { return append(acc, i.Tags...) })
		c.Assert(foldedValue, qt.Equals, 10)
		c.Assert(foldedName, qt.Equals, " John Jane Jack Jill")
		c.Assert(foldedTags, qt.DeepEquals, []string{"a", "b", "a", "c", "b", "c", "b", "c"})
	})
	a.Run("foldRight", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		folded := slices.FoldRight(slice, 1, func(i int, acc int) int { return acc * i })
		c.Assert(folded, qt.Equals, 3628800)
	})
}

func TestReduce(t *testing.T) {
	a := qt.New(t)
	a.Run("reduce", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		reduced := slices.Reduce(slice, func(acc, i int) int { return acc + i })
		c.Assert(reduced.IsSome(), qt.IsTrue)
		c.Assert(reduced.Value(), qt.DeepEquals, 55)
	})
	a.Run("complex reduce", func(c *qt.C) {
		slice := []item{
			{Value: 1, Name: "John", Tags: []string{"a", "b"}},
			{Value: 2, Name: "Jane", Tags: []string{"a", "c"}},
			{Value: 3, Name: "Jack", Tags: []string{"b", "c"}},
			{Value: 4, Name: "Jill", Tags: []string{"b", "c"}},
		}
		reducedValue := slices.Reduce(slice, func(acc, i item) item { return item{Value: acc.Value + i.Value} })
		reducedName := slices.Reduce(slice, func(acc, i item) item { return item{Name: acc.Name + " " + i.Name} })
		reducedTags := slices.Reduce(slice, func(acc, i item) item { return item{Tags: append(acc.Tags, i.Tags...)} })

		c.Assert(reducedValue.IsSome(), qt.IsTrue)
		c.Assert(reducedValue.Value(), qt.DeepEquals, item{Value: 10})
		c.Assert(reducedName.IsSome(), qt.IsTrue)
		c.Assert(reducedName.Value(), qt.DeepEquals, item{Name: "John Jane Jack Jill"})
		c.Assert(reducedTags.IsSome(), qt.IsTrue)
		c.Assert(reducedTags.Value(), qt.DeepEquals, item{Tags: []string{"a", "b", "a", "c", "b", "c", "b", "c"}})
	})
	a.Run("reduceRight", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		reduced := slices.ReduceRight(slice, func(i int, acc int) int { return acc * i })
		c.Assert(reduced.IsSome(), qt.IsTrue)
		c.Assert(reduced.Value(), qt.DeepEquals, 3628800)

		slice2 := []int{}
		reduced2 := slices.ReduceRight(slice2, func(i int, acc int) int { return acc * i })
		c.Assert(reduced2.IsNone(), qt.IsTrue)
	})
}

func TestOps(t *testing.T) {
	a := qt.New(t)
	a.Run("any", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		any := slices.Any(slice, func(i int) bool { return i%2 == 0 })
		c.Assert(any, qt.IsTrue)
		notAny := slices.Any(slice, func(i int) bool { return i%2 == 100 })
		c.Assert(notAny, qt.IsFalse)
	})
	a.Run("all", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		all := slices.All(slice, func(i int) bool { return i >= 0 })
		c.Assert(all, qt.IsTrue)
		notAll := slices.All(slice, func(i int) bool { return i > 3 })
		c.Assert(notAll, qt.IsFalse)
	})
	a.Run("none", func(c *qt.C) {
		c.Assert(slices.None([]int{1, 2, 3, 4, 5}, func(i int) bool { return i < 0 }), qt.IsTrue)
		c.Assert(slices.None([]int{1, 2, 3, 4, 5}, func(i int) bool { return i > 3 }), qt.IsFalse)
	})
}

func TestIndex(t *testing.T) {
	a := qt.New(t)
	a.Run("index", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		found := slices.Index(slice, 2)
		c.Assert(found.IsSome(), qt.IsTrue)
		c.Assert(found.Value(), qt.Equals, 1)

		notFound := slices.Index(slice, 11)
		c.Assert(notFound.IsNone(), qt.IsTrue)
	})
	a.Run("indexBy", func(c *qt.C) {
		slice := []item{
			{Value: 1, Name: "John", Tags: []string{"a", "b"}},
			{Value: 2, Name: "Jane", Tags: []string{"a", "c"}},
			{Value: 3, Name: "Jack", Tags: []string{"b", "c"}},
			{Value: 4, Name: "Jill", Tags: []string{"b", "c"}},
		}
		found := slices.IndexBy(slice, func(i item) bool { return i.Value == 2 })
		c.Assert(found.IsSome(), qt.IsTrue)
		c.Assert(found.Value(), qt.Equals, 1)

		notFound := slices.IndexBy(slice, func(i item) bool { return i.Value == 11 })
		c.Assert(notFound.IsNone(), qt.IsTrue)
	})
	a.Run("lastIndex", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 2, 10}
		found := slices.LastIndex(slice, 2)
		c.Assert(found.IsSome(), qt.IsTrue)
		c.Assert(found.Value(), qt.Equals, 8)

		notFound := slices.LastIndex(slice, 11)
		c.Assert(notFound.IsNone(), qt.IsTrue)
	})
	a.Run("lastIndexBy", func(c *qt.C) {
		slice := []item{
			{Value: 1, Name: "John", Tags: []string{"a", "b"}},
			{Value: 2, Name: "Jane", Tags: []string{"a", "c"}},
			{Value: 3, Name: "Jack", Tags: []string{"b", "c"}},
			{Value: 4, Name: "Jill", Tags: []string{"b", "c"}},
		}
		by := func(lhs, rhs item) bool { return lhs.Value == rhs.Value }

		found := slices.LastIndexBy(slice, slice[2], by)
		c.Assert(found.IsSome(), qt.IsTrue)
		c.Assert(found.Value(), qt.DeepEquals, 2)

		notFound := slices.LastIndexBy(slice, item{Value: 11}, by)
		c.Assert(notFound.IsNone(), qt.IsTrue)
	})
	a.Run("contains", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		c.Assert(slices.Contains(slice, 2), qt.IsTrue)

		c.Assert(slices.Contains(slice, 11), qt.IsFalse)
	})

	a.Run("containsBy", func(c *qt.C) {
		slice := []item{
			{Value: 1, Name: "John", Tags: []string{"a", "b"}},
			{Value: 2, Name: "Jane", Tags: []string{"a", "c"}},
			{Value: 3, Name: "Jack", Tags: []string{"b", "c"}},
			{Value: 4, Name: "Jill", Tags: []string{"b", "c"}},
		}

		c.Assert(slices.ContainsBy(slice, func(i item) bool { return i.Value == 2 }), qt.IsTrue)
		c.Assert(slices.ContainsBy(slice, func(i item) bool { return i.Value == 11 }), qt.IsFalse)
	})
}

func TestNth(t *testing.T) {
	a := qt.New(t)
	in := slices.Nth([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2)
	a.Assert(in.IsSome(), qt.IsTrue)
	a.Assert(in.Value(), qt.Equals, 3)
	notIn := slices.Nth([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11)
	a.Assert(notIn.IsNone(), qt.IsTrue)
}

func TestFlatten(t *testing.T) {
	a := qt.New(t)
	a.Run("flat", func(c *qt.C) {
		slice := [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}
		flat := slices.Flatten(slice)
		c.Assert(flat, qt.DeepEquals, []int{1, 2, 3, 4, 5, 6})
	})
	a.Run("flatBy", func(c *qt.C) {
		slice := []int{1, 2, 3}
		flat := slices.FlattenBy(slice, func(i int) []int {
			return []int{i, i + 1}
		})
		a.Assert(flat, qt.DeepEquals, []int{1, 2, 2, 3, 3, 4})
	})
}

func TestChunk(t *testing.T) {
	a := qt.New(t)
	chunked := slices.Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	a.Assert(chunked, qt.DeepEquals, [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10},
	})
	chunked2 := slices.Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	a.Assert(chunked2, qt.DeepEquals, [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	})
}

func TestGroupBy(t *testing.T) {
}

func TestEqual(t *testing.T) {
}

func TestDeepClone(t *testing.T) {
}

func TestToSet(t *testing.T) {
}

func TestIntersectionBy(t *testing.T) {
}

func TestDifferenceBy(t *testing.T) {
}

func TestDistinct(t *testing.T) {
	a := qt.New(t)
	a.Run("distinct", func(c *qt.C) {
		x := slices.Distinct([]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1})
		slices.Sort(x)
		c.Assert(x, qt.DeepEquals, []int{1, 2, 3, 4, 5})
	})
	a.Run("distinctBy", func(c *qt.C) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		distincted := slices.DistinctBy(slice, func(i int) int { return i % 2 })
		a.Assert(distincted, qt.CmpEquals(), []int{9, 10})
	})
}

func TestSingle(t *testing.T) {
	a := qt.New(t)
	s := slices.Single([]int{1, 2, 3})
	a.Assert(s.IsErr(), qt.IsTrue)
	a.Assert(s.Error(), qt.ErrorMatches, "slice is not scalar")
	s = slices.Single([]int{1})
	a.Assert(s.IsOk(), qt.IsTrue)
	a.Assert(s.Value(), qt.Equals, 1)
	s = slices.Single([]int{})
	a.Assert(s.IsErr(), qt.IsTrue)
	a.Assert(s.Error(), qt.ErrorMatches, "slice is not scalar")
}

func TestReverse(t *testing.T) {
	a := qt.New(t)
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slices.Reverse(slice)
	a.Assert(slice, qt.DeepEquals, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func TestMapSet(t *testing.T) {
	a := qt.New(t)
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a.Run("toSet", func(c *qt.C) {
		set := slices.ToHashSet(slice)
		a.Assert(set, qt.DeepEquals, map[int]struct{}{
			1:  {},
			2:  {},
			3:  {},
			4:  {},
			5:  {},
			6:  {},
			7:  {},
			8:  {},
			9:  {},
			10: {},
		})
	})
	a.Run("toHashMap", func(c *qt.C) {
		m := slices.ToHashMap(slice, func(i int, idx int) (int, int) { return i * 2, idx })
		a.Assert(m, qt.DeepEquals, map[int]int{
			2:  0,
			4:  1,
			6:  2,
			8:  3,
			10: 4,
			12: 5,
			14: 6,
			16: 7,
			18: 8,
			20: 9,
		})
	})
	a.Run("toIndexedMap", func(c *qt.C) {
		m := slices.ToIndexedMap(slice)
		a.Assert(m, qt.DeepEquals, map[int]int{
			0: 1,
			1: 2,
			2: 3,
			3: 4,
			4: 5,
			5: 6,
			6: 7,
			7: 8,
			8: 9,
			9: 10,
		})
	})
}

func TestPartition(t *testing.T) {
	a := qt.New(t)
	a.Run("int", func(c *qt.C) {
		lhs, rhs := slices.Partition([]int{1, 2, 3, 4, 5, 6}, func(i int) bool { return i%2 == 0 })
		c.Assert(lhs, qt.DeepEquals, []int{2, 4, 6})
		c.Assert(rhs, qt.DeepEquals, []int{1, 3, 5})
	})
	a.Run("user", func(c *qt.C) {
		lhs, rhs := slices.Partition([]user{{Id: 100}, {Id: 2}, {Id: 4000}, {Id: 5}}, func(u user) bool { return u.Id < 18 })
		c.Assert(lhs, qt.DeepEquals, []user{{Id: 2}, {Id: 5}})
		c.Assert(rhs, qt.DeepEquals, []user{{Id: 100}, {Id: 4000}})
	})
}
