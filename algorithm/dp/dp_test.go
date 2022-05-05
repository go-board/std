package dp_test

import (
	"reflect"
	"testing"

	"github.com/go-board/std/algorithm/dp"
)

func TestLcs(t *testing.T) {
	t.Run("lcs", func(t *testing.T) {
		testCases := []struct {
			left  []string
			right []string
			want  []string
		}{
			{},
			{left: []string{"a"}},
			{[]string{"a"}, []string{"a"}, []string{"a"}},
			{[]string{"a", "c", "d"}, []string{"b", "c", "a", "d"}, []string{"a", "d"}},
			{[]string{"a", "c", "d", "e"}, []string{"b", "c", "a", "d", "e"}, []string{"a", "d", "e"}},
		}
		for _, tc := range testCases {
			got := dp.Lcs(tc.left, tc.right)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got:%v, want:%v", got, tc.want)
			}
		}
	})
	t.Run("lcs_by", func(t *testing.T) {
		type A struct {
			a int
			b []int
		}
		testCases := []struct {
			left  []A
			right []A
			want  []A
		}{
			{},
			{left: []A{{a: 1}}},
			{[]A{{a: 1}}, []A{{a: 1}}, []A{{a: 1}}},
			{[]A{{a: 1}, {a: 2}, {a: 3}}, []A{{a: 1}, {a: 2}, {a: 3}}, []A{{a: 1}, {a: 2}, {a: 3}}},
			{[]A{{a: 1}, {a: 2}, {a: 3}}, []A{{a: 1}, {a: 2}, {a: 3}, {a: 4}}, []A{{a: 1}, {a: 2}, {a: 3}}},
		}
		for _, tc := range testCases {
			got := dp.LcsBy(tc.left, tc.right, func(a, b A) bool { return a.a == b.a })
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got:%v, want:%v", got, tc.want)
			}
		}
	})
	t.Run("lcs_by_pointer", func(t *testing.T) {
		type A struct {
			a int
			b []int
		}
		testCases := []struct {
			left  []*A
			right []*A
			want  []*A
		}{
			{},
			{left: []*A{{a: 1}}},
			{[]*A{{a: 1}}, []*A{{a: 1}}, []*A{{a: 1}}},
			{[]*A{{a: 1}, {a: 2}, {a: 3}}, []*A{{a: 1}, {a: 2}, {a: 3}}, []*A{{a: 1}, {a: 2}, {a: 3}}},
			{[]*A{{a: 1}, {a: 2}, {a: 3}}, []*A{{a: 1}, {a: 2}, {a: 3}, {a: 4}}, []*A{{a: 1}, {a: 2}, {a: 3}}},
		}
		for _, tc := range testCases {
			got := dp.LcsBy(tc.left, tc.right, func(a, b *A) bool { return a.a == b.a })
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got:%v, want:%v", got, tc.want)
			}
		}
	})
}
