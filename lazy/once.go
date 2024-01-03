//go:build !1.21

package lazy

import (
	"sync"
)

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

// OnceValues port from unreleased std module, [sync.OnceValues]
func OnceValues[E, T any](f func() (E, T)) func() (E, T) {
	var (
		once  sync.Once
		valid bool
		p     any
		e     E
		t     T
	)
	g := func() {
		defer func() {
			p = recover()
			if !valid {
				panic(p)
			}
		}()
		e, t = f()
		valid = true
	}
	return func() (E, T) {
		once.Do(g)
		if !valid {
			panic(p)
		}
		return e, t
	}
}
