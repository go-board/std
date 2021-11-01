package maps

import (
	"sort"
)

type keySorter[T any] struct {
	Keys []T
	cmp  func(T, T) bool
}

func (ks *keySorter) Len() int           { return len(ks.Keys) }
func (ks *keySorter) Less(i, j int) bool { return ks.cmp(ks.Keys[i], ks.Keys[j]) }
func (ks *keySorter) Swap(i, j int)      { ks.Keys[i], ks.Keys[j] = ks.Keys[j], ks.Keys[i] }

func Keys[K comparable, V any](m map[K]V) []K {
	rs := make([]K, 0, len(m))
	for k := range m {
		rs = append(rs, k)
	}
	return rs
}

func SortedKeys[K comparable, V any](m map[K]V, cmp func(K, K) int) []K {
	keys := Keys(m)
	ksWrapper := &keySorter[K]{Keys: keys, cmp: cmp}
	sort.Sort(ksWrapper)
	return keySorter.Keys
}

func Values[K comparable, V any](m map[K]V) []V {
	rs := make([]V, 0, len(m))
	for _, v := range m {
		rs = append(rs, v)
	}
	return rs
}
