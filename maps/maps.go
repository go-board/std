package maps

import "github.com/go-board/std/tuple"

type MapEntry[K, V any] struct {
	inner tuple.Pair[K, V]
}

func entry[K, V any](key K, value V) MapEntry[K, V] {
	return MapEntry[K, V]{inner: tuple.PairOf(key, value)}
}
