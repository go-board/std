//go:build 1.21

package lazy

import "sync"

func OnceFunc(f func()) func() {
	return sync.OnceFunc(f)
}

func OnceValue[T any](f func() T) func() T {
	return sync.OnceValue(f)
}

func OnceValues[T any, E any](f func() (T, E)) func() (T, E) {
	return sync.OnceValues(f)
}
