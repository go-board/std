package lazy

import (
	"sync"

	"github.com/go-board/std/optional"
)

// OnceCell is a cell which can be written to only once.
type OnceCell[T any] struct {
	once sync.Once
	val  *T
}

func NewOnceCell[T any]() OnceCell[T] {
	return OnceCell[T]{}
}

// Get return existing one if exists, otherwise panics
func (self *OnceCell[T]) Get() optional.Optional[T] {
	return optional.FromPtr(self.val)
}

// GetOrInit return existing one if existed, else create a new then return it
func (self *OnceCell[T]) GetOrInit(f func() T) T {
	self.once.Do(func() {
		val := f()
		self.val = &val
	})
	return *self.val
}

// OnceFunc port from unreleased std module, [sync.OnceFunc]
func OnceFunc(f func()) func() {
	var (
		once  sync.Once
		valid bool
		p     any
	)
	// Construct the inner closure just once to reduce costs on the fast path.
	g := func() {
		defer func() {
			p = recover()
			if !valid {
				// Re-panic immediately so on the first call the user gets a
				// complete stack trace into f.
				panic(p)
			}
		}()
		f()
		valid = true // Set only if f does not panic
	}
	return func() {
		once.Do(g)
		if !valid {
			panic(p)
		}
	}
}

// OnceValue port from unreleased std module, [sync.OnceValue]
func OnceValue[T any](f func() T) func() T {
	var (
		once   sync.Once
		valid  bool
		p      any
		result T
	)
	g := func() {
		defer func() {
			p = recover()
			if !valid {
				panic(p)
			}
		}()
		result = f()
		valid = true
	}
	return func() T {
		once.Do(g)
		if !valid {
			panic(p)
		}
		return result
	}
}

// OnceResult port from unreleased std module, [sync.OnceValues]
func OnceResult[T any](f func() (T, error)) func() (T, error) {
	var (
		once   sync.Once
		valid  bool
		p      any
		result T
		err    error
	)
	g := func() {
		defer func() {
			p = recover()
			if !valid {
				panic(p)
			}
		}()
		result, err = f()
		valid = true
	}
	return func() (T, error) {
		once.Do(g)
		if !valid {
			panic(p)
		}
		return result, err
	}
}
