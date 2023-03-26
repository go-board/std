package lazy

import "sync"

type OnceCell[T any] struct {
	once sync.Once
	val  *T
}

func NewOnceCell[T any]() OnceCell[T] {
	return OnceCell[T]{}
}

func (self *OnceCell[T]) Get() *T {
	if self.val == nil {
		panic("")
	}
	return self.val
}
