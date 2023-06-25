package iter_test

import (
	"strconv"
	"testing"

	"github.com/go-board/std/iter"
	"github.com/go-board/std/optional"
)

func assertTrue(t *testing.T, ok bool, args ...any) {
	t.Helper()
	if !ok {
		t.Fatal(args...)
	}
}

func assertFalse(t *testing.T, ok bool, args ...any) {
	t.Helper()
	if ok {
		t.Fatal(args...)
	}
}

func assertEqual[T comparable](t *testing.T, lhs T, rhs T, args ...any) {
	t.Helper()
	if lhs != rhs {
		t.Fatal(args...)
	}
}

func TestAll(t *testing.T) {
	ok1 := iter.AllOf[int](iter.OfSlice([]int{1, 2, 3, 4, 5}), func(i int) bool { return i > 0 })
	assertTrue(t, ok1)
	ok2 := iter.AllOf[int](iter.OfSlice([]int{1, 2, 3, 4, 5}), func(i int) bool { return i > 4 })
	assertFalse(t, ok2)
}

func TestAny(t *testing.T) {
	ok1 := iter.AnyOf[int](iter.OfSlice([]int{1, 2, 3, 4, 5}), func(i int) bool { return i > 1 })
	assertTrue(t, ok1)
	ok2 := iter.AnyOf[int](iter.OfSlice([]int{1, 2, 3, 4, 5}), func(i int) bool { return i < 0 })
	assertFalse(t, ok2)
}

func TestMap(t *testing.T) {
	strings := iter.Map[int](iter.OfSlice([]int{1, 2, 3}), strconv.Itoa)
	other := iter.OfSliceVariadic("1", "2", "3")
	ok := iter.Equal(strings, iter.Iter[string](other))
	assertTrue(t, ok)
}

func TestFilter(t *testing.T) {
	i := iter.Filter[int](iter.OfSlice([]int{1, 2, 3, 4, 5}), func(i int) bool { return i%2 == 0 })
	other := iter.OfSlice([]int{2, 4})
	assertTrue(t, iter.Equal(i, iter.Iter[int](other)))
}

func TestFilterMap(t *testing.T) {
	i := iter.FilterMap[int](iter.OfSlice([]int{1, 2, 3, 4, 5}), func(t int) optional.Optional[string] {
		if t%3 != 0 {
			return optional.Some(strconv.Itoa(t * t))
		}
		return optional.None[string]()
	})
	assertTrue(t, iter.Equal[string](i, iter.OfSlice([]string{"1", "4", "16", "25"})))
}

func TestFoldLeft(t *testing.T) {
	it := iter.OfSliceVariadic(1, 2, 3, 4, 5, 6)
	i := iter.FoldLeft[int](it, 0, func(i int, i2 int) int { return i + i2 })
	assertEqual(t, i, 21)
}

func TestFoldRight(t *testing.T) {
	i := iter.FoldRight[int](iter.OfSlice([]int{1, 2, 3, 4, 5}), 1, func(a int, t int) int {
		return t * a
	})
	assertEqual(t, i, 120)
}
