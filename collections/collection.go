package collections

import (
	"github.com/go-board/std/collections/btree"
	"github.com/go-board/std/delegate"
)

type TreeMap[TKey, TValue any] struct {
	*btree.TreeMap[TKey, TValue]
}

func NewTreeMap[TKey, TValue any](less delegate.Lt[TKey]) *TreeMap[TKey, TValue] {
	return &TreeMap[TKey, TValue]{btree.NewTreeMap[TKey, TValue](less)}
}

type TreeSet[TElement any] struct {
	*btree.TreeSet[TElement]
}

func NewTreeSet[TElement any](less delegate.Lt[TElement]) *TreeSet[TElement] {
	return &TreeSet[TElement]{btree.NewTreeSet(less)}
}
