package iter_test

import (
	"errors"
	"strconv"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/iter"
	"github.com/go-board/std/tuple"
	cmp2 "github.com/google/go-cmp/cmp"
)

func seq[E any](slice ...E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, x := range slice {
			if !yield(x) {
				break
			}
		}
	}
}

func collect[E any](s iter.Seq[E]) []E {
	var slice []E
	iter.CollectFunc[E](s, func(e E) bool { slice = append(slice, e); return true })
	return slice
}

func TestEnumerate(t *testing.T) {
	x := iter.Enumerate(seq(1, 2, 3))
	qt.Assert(t, collect(x), qt.CmpEquals(cmp2.Comparer(func(l, r tuple.Pair[int, int]) bool {
		return l.First() == r.First() && l.Second() == r.Second()
	})), []tuple.Pair[int, int]{tuple.MakePair(0, 1), tuple.MakePair(1, 2), tuple.MakePair(2, 3)})
}

func TestAppend(t *testing.T) {
	x := iter.Append(seq(1), 2, 3)
	qt.Assert(t, collect(x), qt.DeepEquals, []int{1, 2, 3})
}

func TestPrepend(t *testing.T) {
	x := iter.Prepend(seq(1), 2, 3)
	qt.Assert(t, collect(x), qt.DeepEquals, []int{2, 3, 1})
}

func TestTryForEach(t *testing.T) {
	err := iter.TryForEach(seq(1, 2, 3), func(i int) error {
		if i < 2 {
			return nil
		}
		return errors.New("err: elem must less than 2")
	})
	qt.Assert(t, err, qt.ErrorMatches, "err: elem must less than 2")
	err2 := iter.TryForEach(seq(1, 2, 3), func(i int) error { return nil })
	qt.Assert(t, err2, qt.IsNil)
}

func TestForEach(t *testing.T) {
	iter.ForEach(seq(1, 2, 3), func(i int) { t.Logf("foreach, %d", i) })
}

func TestAll(t *testing.T) {
	ok := iter.All(seq(1, 2, 3), func(i int) bool { return i > 0 })
	qt.Assert(t, ok, qt.IsTrue)
	ok2 := iter.All(seq(1, 2, 3), func(i int) bool { return i > 2 })
	qt.Assert(t, ok2, qt.IsFalse)
}

func TestAny(t *testing.T) {
	ok := iter.Any(seq(1, 2, 3), func(i int) bool { return i > 0 })
	qt.Assert(t, ok, qt.IsTrue)
	ok2 := iter.Any(seq(1, 2, 3), func(i int) bool { return i < 0 })
	qt.Assert(t, ok2, qt.IsFalse)
}

func TestFind(t *testing.T) {
	t.Run("find", func(t *testing.T) {
		_, ok := iter.Find(seq(1, 2, 3), func(i int) bool {
			return i > 0
		})
		qt.Assert(t, ok, qt.IsTrue)
	})
	t.Run("find map", func(t *testing.T) {
		e, ok := iter.FindMap(seq(1, 2, 3), func(i int) (string, bool) {
			if i == 2 {
				return strconv.Itoa(i), true
			}
			return "", false
		})
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, e, qt.Equals, "2")
	})
	t.Run("first", func(t *testing.T) {
		e, ok := iter.First(seq(1, 2, 3, -1, 0, 5), func(i int) bool {
			return i < 0
		})
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, e, qt.Equals, -1)
	})
	t.Run("first option", func(t *testing.T) {
		x := iter.FirstOption(seq(1, 2, -1, 0, 3), func(i int) bool {
			return i < 0
		})
		qt.Assert(t, x.IsSomeAnd(func(i int) bool {
			return i == -1
		}), qt.IsTrue)
	})
	t.Run("last", func(t *testing.T) {
		e, ok := iter.Last(seq(1, 2, 3, -1, 0, 5, -9, 0), func(i int) bool {
			return i < 0
		})
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, e, qt.Equals, -9)
	})
}

func TestIndex(t *testing.T) {
	t.Run("index first", func(t *testing.T) {
		x := iter.IndexFirst(seq(1, 2, 3), 2)
		qt.Assert(t, x, qt.Equals, 1)
	})
	t.Run("index first func", func(t *testing.T) {
		x := iter.IndexFirstFunc(seq(1, 2, 3), func(i int) bool { return i > 2 })
		qt.Assert(t, x, qt.Equals, 2)
	})
	t.Run("index last", func(t *testing.T) {
		x := iter.IndexLast(seq(1, 2, 3, 2, 1), 2)
		qt.Assert(t, x, qt.Equals, 3)
	})
	t.Run("index last func", func(t *testing.T) {
		x := iter.IndexLastFunc(seq(1, 2, 3, 2, 1), func(i int) bool {
			return i == 2
		})
		qt.Assert(t, x, qt.Equals, 3)
	})
}

func TestNth(t *testing.T) {
	t.Run("nth", func(t *testing.T) {
		x, ok := iter.Nth(seq(1, 2, 3), 1)
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, x, qt.Equals, 2)
		y, ok := iter.Nth(seq(1, 2, 3), 3)
		qt.Assert(t, ok, qt.IsFalse)
		qt.Assert(t, y, qt.Equals, 0)
	})
	t.Run("nth option", func(t *testing.T) {
		x := iter.NthOption(seq(1, 2, 3), 0)
		qt.Assert(t, x.IsSomeAnd(func(i int) bool { return i == 1 }), qt.IsTrue)
		y := iter.NthOption(seq(1, 2, 3), 1)
		qt.Assert(t, y.IsSomeAnd(func(i int) bool { return i == 2 }), qt.IsTrue)
	})
}

func TestMap(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		s := seq(1, 2, 3)
		s = iter.Map(s, func(i int) int { return i + 1 })
		qt.Assert(t, collect(s), qt.DeepEquals, []int{2, 3, 4})
	})
	t.Run("map while", func(t *testing.T) {
		s := seq(1, 2, -3, 3, -1)
		x := iter.MapWhile(s, func(e int) (string, bool) {
			if e > 0 {
				return strconv.Itoa(e), true
			}
			return "", false
		})
		qt.Assert(t, collect(x), qt.DeepEquals, []string{"1", "2"})
	})
}

func TestFold(t *testing.T) {
	s := seq(1, 2, 3)
	x := iter.Fold(s, 0, func(a int, i int) int { return a + i })
	qt.Assert(t, x, qt.Equals, 6)
}

func TestReduce(t *testing.T) {
	t.Run("has elems", func(t *testing.T) {
		s := seq(1, 2, 3)
		x, ok := iter.Reduce(s, func(a, b int) int { return a + b })
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, x, qt.Equals, 6)
	})
	t.Run("no elems", func(t *testing.T) {
		s := seq[int]()
		_, ok := iter.Reduce(s, func(a, b int) int { return a + b })
		qt.Assert(t, ok, qt.IsFalse)
	})
	t.Run("has elems option", func(t *testing.T) {
		s := seq(1, 2, 3)
		x := iter.ReduceOption(s, func(i int, i2 int) int { return i + i2 })
		qt.Assert(t, x.IsSome(), qt.IsTrue)
		qt.Assert(t, x.Value(), qt.Equals, 6)
	})
	t.Run("no elems option", func(t *testing.T) {
		s := seq[int]()
		x := iter.ReduceOption(s, func(i int, i2 int) int { return i + i2 })
		qt.Assert(t, x.IsNone(), qt.IsTrue)
	})
}

func TestFilter(t *testing.T) {
	t.Run("filter", func(t *testing.T) {
		s := seq(1, 2, 3)
		s = iter.Filter(s, func(i int) bool { return i > 1 })
		qt.Assert(t, collect(s), qt.DeepEquals, []int{2, 3})
	})
	t.Run("filter_map", func(t *testing.T) {
		s := seq(1, 2, 3)
		x := iter.FilterMap(s, func(e int) (string, bool) {
			if e > 1 {
				return strconv.Itoa(e), true
			}
			return "", false
		})
		qt.Assert(t, collect(x), qt.DeepEquals, []string{"2", "3"})
	})
	t.Run("filter_zero", func(t *testing.T) {
		s := seq(0, 2, 4, 5, 6, 0, 3)
		s = iter.FilterZero(s)
		qt.Assert(t, collect(s), qt.DeepEquals, []int{2, 4, 5, 6, 3})
	})
}

func TestScan(t *testing.T) {
	x := iter.Scan(seq(1, 2, 3, 4, 5), 1, func(state int, elem int) int { return state * elem })
	qt.Assert(t, collect(x), qt.DeepEquals, []int{1, 2, 6, 24, 120})
}

func TestTryFold(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		s := seq("1", "2", "3")
		x, err := iter.TryFold(s, 0, func(a int, i string) (int, error) {
			x, err := strconv.Atoi(i)
			if err != nil {
				return a, err
			}
			return a + x, nil
		})
		qt.Assert(t, err, qt.IsNil)
		qt.Assert(t, x, qt.Equals, 6)
	})
	t.Run("error", func(t *testing.T) {
		s := seq("a", "b", "c")
		_, err := iter.TryFold(s, 0, func(a int, i string) (int, error) {
			x, err := strconv.Atoi(i)
			if err != nil {
				return a, err
			}
			return a + x, nil
		})
		qt.Assert(t, err, qt.IsNotNil)
	})
}

func TestMax(t *testing.T) {
	s := seq(1, 2, 3)
	t.Run("max", func(t *testing.T) {
		x, ok := iter.Max(s)
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, x, qt.Equals, 3)
	})
	t.Run("max option", func(t *testing.T) {
		x := iter.MaxOption(s)
		qt.Assert(t, x.IsSome(), qt.IsTrue)
		qt.Assert(t, x.Value(), qt.Equals, 3)
	})
	t.Run("max_func", func(t *testing.T) {
		x, ok := iter.MaxFunc(s, cmp.Compare[int])
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, x, qt.Equals, 3)
	})
	t.Run("max_by_key", func(t *testing.T) {
		x, ok := iter.MaxByKey(s, func(e int) int { return e })
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, x, qt.Equals, 3)
	})
}

func TestMin(t *testing.T) {
	s := seq(1, 2, 3)
	t.Run("min", func(t *testing.T) {
		x, ok := iter.Min(s)
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, x, qt.Equals, 1)
	})
	t.Run("min_func", func(t *testing.T) {
		x, ok := iter.MinFunc(s, cmp.Compare[int])
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, x, qt.Equals, 1)
	})
	t.Run("min_by_key", func(t *testing.T) {
		x, ok := iter.MinByKey(s, func(e int) int { return e })
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, x, qt.Equals, 1)
	})
}

func TestMinMax(t *testing.T) {
	s := seq(1, 2, 3)
	t.Run("minmax", func(t *testing.T) {
		x, ok := iter.MinMax(s)
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, x.First(), qt.Equals, 1)
		qt.Assert(t, x.Second(), qt.Equals, 3)
	})
	t.Run("minmax_func", func(t *testing.T) {
		x, ok := iter.MinMaxFunc(s, cmp.Compare[int])
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, x.First(), qt.Equals, 1)
		qt.Assert(t, x.Second(), qt.Equals, 3)
	})
	t.Run("minmax_by_key", func(t *testing.T) {
		x, ok := iter.MinMaxByKey(s, func(e int) int { return e })
		qt.Assert(t, ok, qt.IsTrue)
		qt.Assert(t, x.First(), qt.Equals, 1)
		qt.Assert(t, x.Second(), qt.Equals, 3)
	})
}

func TestCountFunc(t *testing.T) {
	s := seq(1, 2, 3)
	c := iter.CountFunc(s, func(i int) bool { return i > 1 })
	qt.Assert(t, c, qt.Equals, 2)
}

func TestCount(t *testing.T) {
	s := seq(1, 2, 3, 3, 3, 4, 5, 6, 7)
	c := iter.Count(s, 3)
	qt.Assert(t, c, qt.Equals, 3)
}

func TestSize(t *testing.T) {
	s := seq(1, 2, 3)
	c := iter.Size(s)
	qt.Assert(t, c, qt.Equals, 3)
}

func TestIsSorted(t *testing.T) {
	t.Run("sorted", func(t *testing.T) {
		s := seq(1, 2, 3)
		ok := iter.IsSorted(s)
		qt.Assert(t, ok, qt.IsTrue)
	})
	t.Run("not sorted", func(t *testing.T) {
		s := seq(1, 2, 3, 2)
		ok := iter.IsSorted(s)
		qt.Assert(t, ok, qt.IsFalse)
	})
}

func TestTake(t *testing.T) {
	s := seq(1, 2, 3, 4, 5)
	s = iter.Take(s, 2)
	qt.Assert(t, collect(s), qt.DeepEquals, []int{1, 2})
}

func TestTakeWhile(t *testing.T) {
	s := seq(1, 2, 3, 2, 1)
	s = iter.TakeWhile(s, func(i int) bool { return i < 3 })
	qt.Assert(t, collect(s), qt.DeepEquals, []int{1, 2})
}

func TestSkip(t *testing.T) {
	s := seq(1, 2, 3, 4, 5)
	s = iter.Skip(s, 2)
	qt.Assert(t, collect(s), qt.DeepEquals, []int{3, 4, 5})
	t.Run("nest skip", func(t *testing.T) {
		s := seq(1, 2, 3, 4, 5)
		s = iter.Skip(s, 1)
		s = iter.Skip(s, 1)
		qt.Assert(t, collect(s), qt.DeepEquals, []int{3, 4, 5})
	})
	t.Run("recursive", func(t *testing.T) {
		x := iter.Skip(iter.Skip(seq(1, 2, 3), 1), 1)
		qt.Assert(t, collect(x), qt.DeepEquals, []int{3})
	})
}

func TestSkipWhile(t *testing.T) {
	s := seq(1, 2, 3, 4, 5)
	s = iter.SkipWhile(s, func(i int) bool { return i < 3 })
	qt.Assert(t, collect(s), qt.DeepEquals, []int{3, 4, 5})
}

func TestFlatten(t *testing.T) {
	s := seq(seq(1, 2), seq(3, 4, 5))
	x := iter.Flatten(s)
	qt.Assert(t, collect(x), qt.DeepEquals, []int{1, 2, 3, 4, 5})
}

func TestCombine(t *testing.T) {
	x := iter.Flatten(seq(seq(1, 2, 3), seq(4, 5, 6)))
	y := iter.Map(x, func(e int) int { return e * e })
	qt.Assert(t, collect(y), qt.DeepEquals, []int{1, 4, 9, 16, 25, 36})
}

func TestDedup(t *testing.T) {
	x := seq(1, 2, 3, 3, 2, 2, 1)
	x = iter.Dedup(x)
	qt.Assert(t, collect(x), qt.DeepEquals, []int{1, 2, 3, 2, 1})
}

func TestStepBy(t *testing.T) {
	x := iter.StepBy(seq(1, 2, 3, 4, 5, 6, 7, 8, 9), 2)
	qt.Assert(t, collect(x), qt.DeepEquals, []int{1, 3, 5, 7, 9})
}

func TestIntersperse(t *testing.T) {
	x := iter.Intersperse(seq(1, 2, 3), 0)
	qt.Assert(t, collect(x), qt.DeepEquals, []int{1, 0, 2, 0, 3})
}

func TestChain(t *testing.T) {
	x := iter.Chain(seq(1, 2, 3), seq(4, 5, 6))
	qt.Assert(t, collect(x), qt.DeepEquals, []int{1, 2, 3, 4, 5, 6})
}

func TestInspect(t *testing.T) {
	x := iter.Inspect(seq(1, 2, 3), func(i int) { t.Log(i) })
	qt.Assert(t, collect(x), qt.DeepEquals, []int{1, 2, 3})
}

func TestCollectFunc(t *testing.T) {
	var s []int
	iter.CollectFunc(seq(1, 2, 3), func(i int) bool {
		s = append(s, i*i+i)
		return true
	})
	qt.Assert(t, s, qt.DeepEquals, []int{2, 6, 12})
}

func TestCollect(t *testing.T) {
	x := iter.Collect(seq(1, 2, 3), collect[int])
	qt.Assert(t, x, qt.DeepEquals, []int{1, 2, 3})
}

func TestPartition(t *testing.T) {
	s := seq(1, 2, 3, 4, 5, 6)
	x := iter.Partition(s, func(i int) bool {
		return i%2 == 0
	})
	qt.Assert(t, collect(x.First()), qt.DeepEquals, []int{2, 4, 6})
	qt.Assert(t, collect(x.Second()), qt.DeepEquals, []int{1, 3, 5})
}

func TestIsPartitioned(t *testing.T) {
	s := seq(2, 4, 6)
	qt.Assert(t, iter.IsPartitioned(s, func(i int) bool {
		return i%2 == 0
	}), qt.IsTrue)
	qt.Assert(t, iter.IsPartitioned(s, func(i int) bool {
		return i%3 == 0
	}), qt.IsFalse)
}

func TestUnzip(t *testing.T) {
	t.Run("unzip", func(t *testing.T) {
		x := seq(tuple.MakePair(1, 2), tuple.MakePair(3, 4), tuple.MakePair(5, 6))
		z := iter.Unzip(x)
		qt.Assert(t, collect(z.First()), qt.DeepEquals, []int{1, 3, 5})
		qt.Assert(t, collect(z.Second()), qt.DeepEquals, []int{2, 4, 6})
	})
	t.Run("unzip func", func(t *testing.T) {
		x := seq(1, 2, 3)
		z := iter.UnzipFunc(x, func(a int) (int, int) { return a, a - 2 })
		qt.Assert(t, collect(z.First()), qt.DeepEquals, []int{1, 2, 3})
		qt.Assert(t, collect(z.Second()), qt.DeepEquals, []int{-1, 0, 1})
	})
}

func TestContains(t *testing.T) {
	t.Run("contains", func(t *testing.T) {
		x := iter.Contains(seq(1, 2, 3), 2)
		qt.Assert(t, x, qt.IsTrue)
		y := iter.Contains(seq(1, 2, 3), 4)
		qt.Assert(t, y, qt.IsFalse)
	})
	t.Run("contains func", func(t *testing.T) {
		x := iter.ContainsFunc(seq(1, 2, 3), func(e int) bool {
			return e == 2
		})
		qt.Assert(t, x, qt.IsTrue)
		y := iter.ContainsFunc(seq(1, 2, 3), func(e int) bool {
			return e == 5
		})
		qt.Assert(t, y, qt.IsFalse)
	})
}

// Go 中没有不可变参数，因此在迭代过程中会存在对原始seq修改的可能。
// 真是糟糕的设计呀。
func TestMutate(t *testing.T) {
	type user struct{ Age int }
	x := seq(&user{1}, &user{2})
	iter.ForEach(x, func(user *user) { user.Age *= 2 })
	qt.Assert(t, collect(x), qt.DeepEquals, []*user{{2}, {4}})
}

func TestHead(t *testing.T) {
	t.Run("option", func(t *testing.T) {
		x := seq(1, 2, 3)
		h := iter.HeadOption(x)
		qt.Assert(t, h.IsSomeAnd(func(i int) bool { return i == 1 }), qt.IsTrue)
		x = iter.Tail(x)
		h = iter.HeadOption(x)
		qt.Assert(t, h.IsSomeAnd(func(i int) bool { return i == 2 }), qt.IsTrue)
	})
}

func TestTail(t *testing.T) {
	x := seq(1, 2, 3)
	x = iter.Tail(x)
	qt.Assert(t, collect(x), qt.DeepEquals, []int{2, 3})
	x = iter.Tail(x)
	qt.Assert(t, collect(x), qt.DeepEquals, []int{3})
}
