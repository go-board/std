package iter_test

import (
	"strconv"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-board/std/iter"
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

func collectSlice[E any](s iter.Seq[E]) []E {
	var slice []E
	iter.ForEach(s, func(x E) { slice = append(slice, x) })
	return slice
}

func TestEnumerate(t *testing.T) {
	x := iter.Enumerate(seq(1, 2, 3))
	qt.Assert(t, collectSlice(x), qt.DeepEquals, []iter.Tuple[int, int]{{0, 1}, {1, 2}, {2, 3}})
}

func TestAll(t *testing.T) {
	ok := iter.All(seq(1, 2, 3), func(i int) bool {
		return i > 0
	})
	qt.Assert(t, ok, qt.IsTrue)
}

func TestAny(t *testing.T) {
	ok := iter.Any(seq(1, 2, 3), func(i int) bool {
		return i > 0
	})
	qt.Assert(t, ok, qt.IsTrue)
}

func TestFind(t *testing.T) {
	_, ok := iter.Find(seq(1, 2, 3), func(i int) bool {
		return i > 0
	})
	qt.Assert(t, ok, qt.IsTrue)
}

func TestIndex(t *testing.T) {
	x, ok := iter.Index(seq(1, 2, 3), func(i int) bool {
		return i > 2
	})
	qt.Assert(t, ok, qt.IsTrue)
	qt.Assert(t, x, qt.Equals, 2)
}

func TestMap(t *testing.T) {
	s := seq(1, 2, 3)
	s = iter.Map(s, func(i int) int { return i + 1 })
	qt.Assert(t, collectSlice(s), qt.DeepEquals, []int{2, 3, 4})
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
}

func TestFilter(t *testing.T) {
	s := seq(1, 2, 3)
	s = iter.Filter(s, func(i int) bool { return i > 1 })
	qt.Assert(t, collectSlice(s), qt.DeepEquals, []int{2, 3})
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
	x, ok := iter.Max(s)
	qt.Assert(t, ok, qt.IsTrue)
	qt.Assert(t, x, qt.Equals, 3)
}

func TestMin(t *testing.T) {
	s := seq(1, 2, 3)
	x, ok := iter.Min(s)
	qt.Assert(t, ok, qt.IsTrue)
	qt.Assert(t, x, qt.Equals, 1)
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
	s := seq(1, 2, 3)
	ok := iter.IsSorted(s)
	qt.Assert(t, ok, qt.IsTrue)
}

func TestTake(t *testing.T) {
	s := seq(1, 2, 3, 4, 5)
	s = iter.Take(s, 2)
	qt.Assert(t, collectSlice(s), qt.DeepEquals, []int{1, 2})
}

func TestSkip(t *testing.T) {
	s := seq(1, 2, 3, 4, 5)
	s = iter.Skip(s, 2)
	qt.Assert(t, collectSlice(s), qt.DeepEquals, []int{3, 4, 5})
}

func TestFlatten(t *testing.T) {
	s := seq(seq(1, 2), seq(3, 4, 5))
	x := iter.Flatten(s)
	qt.Assert(t, collectSlice(x), qt.DeepEquals, []int{1, 2, 3, 4, 5})
}

func TestCombine(t *testing.T) {
	x := iter.Flatten(seq(seq(1, 2, 3), seq(4, 5, 6)))
	y := iter.Map(x, func(e int) int { return e * e })
	qt.Assert(t, collectSlice(y), qt.DeepEquals, []int{1, 4, 9, 16, 25, 36})
}
