package slices_test

import (
	"fmt"

	"github.com/go-board/std/slices"
)

type user struct {
	Id   int64  `json:"id"`
	Name string `json:"Name"`
}

func (u user) Clone() user {
	return user{Id: u.Id, Name: u.Name}
}

func ExampleSortBy() {
	slice := []int{3, 1, 2}
	slices.SortBy(slice, func(a, b int) bool { return a < b })
	fmt.Println(slice)
	// Output:
	// [1 2 3]
}

func ExampleSortBy_key() {
	slice := []user{
		{Id: 3, Name: "Jack"},
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
	}
	slices.SortBy(slice, func(a, b user) bool { return a.Id < b.Id })
	fmt.Println(slice)
	// Output:
	// [{1 John} {2 Jane} {3 Jack}]
}

func ExampleSort() {
	slice := []int{3, 1, 2}
	slices.Sort(slice)
	fmt.Println(slice)
	// Output:
	// [1 2 3]
}

func ExampleIsSortedBy() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 3, Name: "Jack"},
	}
	isSortedBy := slices.IsSortedBy(slice, func(a, b user) bool { return a.Id < b.Id })
	fmt.Println(isSortedBy)
	// Output:
	// true
}

func ExampleIsSortedBy_not() {
	slice := []user{
		{Id: 2, Name: "Jane"},
		{Id: 3, Name: "Jack"},
		{Id: 1, Name: "John"},
	}
	isSortedBy := slices.IsSortedBy(slice, func(a, b user) bool { return a.Id < b.Id })
	fmt.Println(isSortedBy)
	// Output:
	// false
}

func ExampleIsSorted() {
	slice := []int{1, 2, 3}
	isSorted := slices.IsSorted(slice)
	fmt.Println(isSorted)
	// Output:
	// true
}

func ExampleIsSorted_not() {
	slice := []int{3, 2, 1}
	isSorted := slices.IsSorted(slice)
	fmt.Println(isSorted)
	// Output:
	// false
}

func ExampleMap() {
	slice := []int{1, 2, 3}
	result := slices.Map(slice, func(i int) int { return i * 2 })
	fmt.Println(result)
	// Output:
	// [2 4 6]
}

func ExampleMap_user() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 3, Name: "Jack"},
	}
	ids := slices.Map(slice, func(u user) int64 { return u.Id })
	names := slices.Map(slice, func(u user) string { return u.Name })
	fmt.Println(ids)
	fmt.Println(names)
	// Output:
	// [1 2 3]
	// [John Jane Jack]
}

func ExampleForEach() {
	slice := []int{1, 2, 3}
	slices.ForEach(slice, func(i int) { fmt.Println(i) })
	// Output:
	// 1
	// 2
	// 3
}

func ExampleForEachIndexed() {
	slice := []int{1, 2, 3}
	slices.ForEachIndexed(slice, func(i int, v int) { fmt.Println(i, v) })
	// Output:
	// 1 0
	// 2 1
	// 3 2
}

func ExampleFilter() {
	slice := []int{1, 2, 3}
	result := slices.Filter(slice, func(i int) bool { return i > 1 })
	fmt.Println(result)
	// Output:
	// [2 3]
}

func ExampleFilter_user() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 3, Name: "Jack"},
	}
	result := slices.Filter(slice, func(u user) bool { return u.Id&1 == 1 })
	fmt.Println(result)
	// Output:
	// [{1 John} {3 Jack}]
}

func ExampleFold() {
	slice := []int{1, 2, 3}
	result := slices.Fold(slice, 0, func(acc, i int) int { return acc + i })
	fmt.Println(result)
	// Output:
	// 6
}

func ExampleFold_diff() {
	slice := []int{1, 2, 3}
	result := slices.Fold(slice, "", func(acc string, i int) string {
		return acc + fmt.Sprintf("%d", i)
	})
	fmt.Println(result)
	// Output:
	// 123
}

func ExampleFold_user() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 3, Name: "Jack"},
	}
	result := slices.Fold(slice, "", func(acc string, u user) string { return acc + u.Name })
	fmt.Println(result)
	// Output:
	// JohnJaneJack
}

func ExampleReduce() {
	slice := []int{1, 2, 3}
	result := slices.Reduce(slice, func(acc, i int) int { return acc + i })
	fmt.Println(result)
	// Output:
	// Some(6)
}

func ExampleReduce_user() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 3, Name: "Jack"},
	}
	result := slices.Reduce(slice, func(acc, u user) user { return u })
	fmt.Println(result)
	// Output:
	// Some({Id:3 Name:Jack})
}

func ExampleReduce_none() {
	slice := []int{}
	result := slices.Reduce(slice, func(acc, i int) int { return acc + i })
	fmt.Println(result)
	// Output:
	// None
}

func ExampleAny() {
	slice := []int{1, 2, 3}
	result := slices.Any(slice, func(i int) bool { return i > 1 })
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAny_user() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
	}
	result := slices.Any(slice, func(u user) bool { return u.Id == 2 })
	fmt.Println(result)
	// Output:
	// true
}

func ExampleAny_false() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
	}
	result := slices.Any(slice, func(u user) bool { return u.Id == 3 })
	fmt.Println(result)
	// Output:
	// false
}

func ExampleAll() {
	slice := []int{1, 2, 3}
	result := slices.All(slice, func(i int) bool { return i > 0 })
	fmt.Println(result)
	// Output:
	// true
}

func ExampleNone() {
	slice := []int{1, 2, 3}
	result := slices.None(slice, func(i int) bool { return i > 5 })
	fmt.Println(result)
	// Output:
	// true
}

func ExampleFindIndexBy() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
	}
	result := slices.FindIndexBy(slice, user{Id: 2}, func(t, u user) bool { return u.Id == t.Id })
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleFindIndex() {
	slice := []int{1, 2, 3}
	result := slices.FindIndex(slice, 2)
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleContainsBy() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
	}
	result := slices.ContainsBy(slice, user{Id: 2}, func(t, u user) bool { return u.Id == t.Id })
	fmt.Println(result)
	// Output:
	// true
}

func ExampleContains() {
	slice := []int{1, 2, 3}
	result := slices.Contains(slice, 6)
	fmt.Println(result)
	// Output:
	// false
}

func ExampleMaxBy() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 3, Name: "Jack"},
		{Id: 4, Name: "Bob"},
	}
	maxId := slices.MaxBy(slice, func(a, b user) bool { return a.Id < b.Id })
	maxName := slices.MaxBy(slice, func(a, b user) bool { return a.Name < b.Name })
	fmt.Println(maxId, maxName)
	// Output:
	// Some({Id:4 Name:Bob}) Some({Id:1 Name:John})
}

func ExampleMax() {
	slice := []int{1, 2, 3}
	result := slices.Max(slice)
	fmt.Println(result)
	// Output:
	// Some(3)
}

func ExampleMinBy() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 3, Name: "Jack"},
		{Id: 4, Name: "Bob"},
	}
	minId := slices.MinBy(slice, func(a, b user) bool { return a.Id < b.Id })
	minName := slices.MinBy(slice, func(a, b user) bool { return a.Name < b.Name })
	fmt.Println(minId, minName)
	// Output:
	// Some({Id:1 Name:John}) Some({Id:4 Name:Bob})
}

func ExampleMin() {
	slice := []int{1, 2, 3}
	result := slices.Min(slice)
	fmt.Println(result)
	// Output:
	// Some(1)
}

func ExampleNth_in() {
	slice := []int{1, 2, 3}
	result := slices.Nth(slice, 1)
	fmt.Println(result)
	// Output:
	// Some(2)
}

func ExampleNth_not_in() {
	slice := []int{1, 2, 3}
	result := slices.Nth(slice, 4)
	fmt.Println(result)
	// Output:
	// None
}

func ExampleFlattenBy() {
	slice := []int{1, 2, 3}
	result := slices.FlattenBy(slice, func(i int) []int {
		s := make([]int, i)
		for j := 0; j < i; j++ {
			s[j] = j
		}
		return s
	})
	fmt.Println(result)
	// Output:
	// [0 0 1 0 1 2]
}

func ExampleFlatten() {
	slice := [][]int{{1, 2, 3}, {4, 5, 6}}
	result := slices.Flatten(slice)
	fmt.Println(result)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleFlatten_empty() {
	slice := [][]int{}
	result := slices.Flatten(slice)
	fmt.Println(result)
	// Output:
	// []
}

func ExampleChunk() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := slices.Chunk(slice, 3)
	fmt.Println(result)
	// Output:
	// [[1 2 3] [4 5 6] [7 8 9] [10]]
}

func ExampleChunk_empty() {
	slice := []int{}
	result := slices.Chunk(slice, 3)
	fmt.Println(result)
	// Output:
	// []
}

func ExampleGroupBy() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 2, Name: "Jack"},
		{Id: 4, Name: "Bob"},
	}
	result := slices.GroupBy(slice, func(u user) (int64, user) { return u.Id, u })
	fmt.Println(result)
	// Output:
	// map[1:[{1 John}] 2:[{2 Jane} {2 Jack}] 4:[{4 Bob}]]
}

func ExampleEqualBy_full() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 2, Name: "Jack"},
		{Id: 4, Name: "Bob"},
	}
	slice2 := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 2, Name: "Jack"},
		{Id: 4, Name: "Bob"},
	}
	result := slices.EqualBy(slice, slice2, func(a, b user) bool { return a.Id == b.Id && a.Name == b.Name })
	fmt.Println(result)
	// Output:
	// true
}

func ExampleEqualBy_partial() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 2, Name: "Jack"},
		{Id: 4, Name: "Bob"},
	}
	slice2 := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 2, Name: "Jack"},
		{Id: 4, Name: "Bob"},
	}
	result := slices.EqualBy(slice, slice2, func(a, b user) bool { return a.Id == b.Id })
	fmt.Println(result)
	// Output:
	// true
}

func ExampleEqual() {
	slice := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	result := slices.Equal(slice, slice2)
	fmt.Println(result)
	slice3 := []int{1, 2, 3, 4}
	result2 := slices.Equal(slice, slice3)
	fmt.Println(result2)
	slice4 := []int{1, 2, 4}
	result3 := slices.Equal(slice, slice4)
	fmt.Println(result3)
	// Output:
	// true
	// false
	// false
}

func ExampleDeepClone() {
	slice := []int{1, 2, 3}
	result := slices.DeepCloneBy(slice, func(i int) int { return i })
	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleDeepClone_user() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 2, Name: "Jack"},
	}
	result := slices.DeepClone(slice)
	fmt.Println(result)
	// Output:
	// [{1 John} {2 Jane} {2 Jack}]
}

func ExampleClone() {
	slice := []int{1, 2, 3}
	result := slices.Clone(slice)
	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleToSet() {
	slice := []int{1, 2, 3}
	result := slices.ToHashSet(slice)
	fmt.Println(result)
	// Output:
	// map[1:{} 2:{} 3:{}]
}

func ExampleToSet_duplicate() {
	slice := []int{1, 2, 3, 3}
	result := slices.ToHashSet(slice)
	fmt.Println(result)
	// Output:
	// map[1:{} 2:{} 3:{}]
}

func ExampleIntersectionBy() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 3, Name: "Jack"},
		{Id: 4, Name: "Bob"},
	}
	slice2 := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 5, Name: "Bob"},
		{Id: 6, Name: "Jack"},
	}
	result := slices.IntersectionBy(slice, slice2, func(a, b user) bool { return a.Id == b.Id && a.Name == b.Name })
	fmt.Println(result)
	// Output:
	// [{1 John} {2 Jane}]
}

func ExampleDifferenceBy() {
	slice := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 3, Name: "Jack"},
		{Id: 4, Name: "Bob"},
	}
	slice2 := []user{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 5, Name: "Bob"},
		{Id: 6, Name: "Jack"},
	}
	result := slices.DifferenceBy(slice, slice2, func(a, b user) bool { return a.Id == b.Id && a.Name == b.Name })
	fmt.Println(result)
	// Output:
	// [{3 Jack} {4 Bob}]
}

func ExamplePartition() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	lhs, rhs := slices.Partition(slice, func(x int) bool { return x%2 == 0 })
	fmt.Println(lhs)
	fmt.Println(rhs)
	// Output:
	// [2 4 6 8 10]
	// [1 3 5 7 9]
}
