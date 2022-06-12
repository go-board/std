package collections

import (
	"github.com/go-board/std/collections/btree"
)

type TreeMap[TKey, TValue any] struct {
	*btree.TreeMap[TKey, TValue]
}

func NewTreeMap[TKey, TValue any](less func(TKey, TKey) bool) *TreeMap[TKey, TValue] {
	return &TreeMap[TKey, TValue]{btree.NewTreeMap[TKey, TValue](less)}
}

type TreeSet[TElement any] struct {
	*btree.TreeSet[TElement]
}

func NewTreeSet[TElement any](less func(TElement, TElement) bool) *TreeSet[TElement] {
	return &TreeSet[TElement]{btree.NewTreeSet(less)}
}
