package maps

import (
	"sort"

	"github.com/go-board/std/delegate"
)

type keySorter[T any] struct {
	ord  delegate.Ord[T]
	keys []T
}

func (self *keySorter[T]) Len() int           { return len(self.keys) }
func (self *keySorter[T]) Less(i, j int) bool { return self.ord(self.keys[i], self.keys[j]) < 0 }
func (self *keySorter[T]) Swap(i, j int)      { self.keys[i], self.keys[j] = self.keys[j], self.keys[i] }

func Keys[K comparable, V any](m map[K]V) []K {
	rs := make([]K, 0, len(m))
	for k := range m {
		rs = append(rs, k)
	}
	return rs
}

func SortedKeys[K comparable, V any](m map[K]V, ord delegate.Ord[K]) []K {
	keys := Keys(m)
	ksWrapper := &keySorter[K]{keys: keys, ord: ord}
	sort.Sort(ksWrapper)
	return ksWrapper.keys
}

func Values[K comparable, V any](m map[K]V) []V {
	rs := make([]V, 0, len(m))
	for _, v := range m {
		rs = append(rs, v)
	}
	return rs
}
