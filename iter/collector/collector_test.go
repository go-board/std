package collector_test

import (
	"strconv"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-board/std/iter"
	"github.com/go-board/std/iter/collector"
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

func TestToSlice(t *testing.T) {
	ints := collector.Collect(seq(1, 2, 3, 4, 5), collector.ToSlice[int]())
	qt.Assert(t, ints, qt.DeepEquals, []int{1, 2, 3, 4, 5})
}

func TestToMap(t *testing.T) {
	mapFn := func(e int) (string, bool) { return strconv.Itoa(e), e%2 == 0 }
	m := collector.Collect(seq(1, 2, 3, 4, 5), collector.ToMap(mapFn))

	qt.Assert(t, m, qt.DeepEquals, map[string]bool{
		"1": false,
		"2": true,
		"3": false,
		"4": true,
		"5": false,
	})
}

func TestToOrderedSet(t *testing.T) {
	s := collector.Collect(seq(1, 2, 3, 4, 5), collector.ToOrderedSet(func(e int) int { return e % 3 }))
	qt.Assert(t, s.Len(), qt.Equals, 3)
	x := collector.Collect(s.AscendIter(), collector.ToSlice[int]())
	qt.Assert(t, x, qt.DeepEquals, []int{0, 1, 2})
}

func TestToOrderedMap(t *testing.T) {
	s := collector.Collect(seq(1, 2, 3, 4, 5), collector.ToOrderedMap(func(e int) (int, string) {
		return e, strconv.Itoa(e)
	}))
	keys := collector.Collect(s.Keys(), collector.ToSlice[int]())
	values := collector.Collect(s.Values(), collector.ToSlice[string]())

	qt.Assert(t, keys, qt.DeepEquals, []int{1, 2, 3, 4, 5})

	qt.Assert(t, values, qt.DeepEquals, []string{"1", "2", "3", "4", "5"})

}

func TestDistinct(t *testing.T) {
	m := collector.Collect(seq(1, 1), collector.Distinct[int]())
	x := collector.Collect(m, collector.ToSlice[int]())
	qt.Assert(t, x, qt.DeepEquals, []int{1})
}

func TestDistinctFunc(t *testing.T) {
	m := collector.Collect(seq(1, 2, 3, 2, 1), collector.DistinctFunc(func(x, y int) int {
		return x - y
	}))
	x := collector.Collect(m, collector.ToSlice[int]())
	qt.Assert(t, x, qt.DeepEquals, []int{1, 2, 3})
}

func TestCollectOne(t *testing.T) {
	c := collector.ToSlice[int]()
	c.Collect(1)
	c.Collect(2)
	qt.Assert(t, c.Finish(), qt.DeepEquals, []int{1, 2})
}

func TestChunk(t *testing.T) {
	m := collector.Collect(seq(1, 2, 3, 4, 5), collector.Chunk[int](2))
	x := collector.Collect(iter.Map(m, func(e iter.Seq[int]) []int {
		return collector.Collect(e, collector.ToSlice[int]())
	}), collector.ToSlice[[]int]())
	qt.Assert(t, x, qt.DeepEquals, [][]int{{1, 2}, {3, 4}, {5}})
}
