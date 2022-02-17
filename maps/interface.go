package maps

import (
	"github.com/go-board/std/delegate"
	"github.com/go-board/std/optional"
)

type Map[TKey, TValue any] interface {
	Put(key TKey, value TValue)
	PutIfAbsent(key TKey, value TValue)
	Replace(key TKey, value TValue) optional.Optional[TValue]
	ReplaceWhen(key TKey, old TValue, new TValue, eq delegate.Equal[TValue])
	Size() uint
	Iterate(fn delegate.Consumer2[TKey, TValue])
	Contains(key TKey) bool
	Del(key TKey) bool
	Get(key TKey) optional.Optional[TValue]
	GetOrPut(key TKey, value TValue)
}
