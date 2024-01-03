package source

import (
	"net"

	"github.com/go-board/std/core"
	"github.com/go-board/std/iter"
)

// Variadic creates an [iter.Seq] from variadic elements.
func Variadic[E any](elems ...E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, elem := range elems {
			if !yield(elem) {
				break
			}
		}
	}
}

// Gen creates an [iter.Seq] that yield number in range.
//
// If `to` less-equal than `from`, yield nothing as [Empty].
//
// Example:
//
//	source.Gen(2,5) => seq: 2,3,4
//	source.Gen(2,1) => seq:
func Gen[N core.Integer](from N, to N) iter.Seq[N] {
	return func(yield func(N) bool) {
		for i := from; i < to; i++ {
			if !yield(i) {
				break
			}
		}
	}
}

// Once creates an [iter.Seq] that yields an element exactly once.
func Once[E any](v E) iter.Seq[E] {
	return func(yield func(E) bool) {
		yield(v)
	}
}

// Repeat creates an [iter.Seq] that endlessly repeats a single element.
func Repeat[E any](v E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for {
			if !yield(v) {
				break
			}
		}
	}
}

// RepeatTimes creates an [iter.Seq] that repeats a single element a given number of times.
func RepeatTimes[E any, N core.Integer](v E, n N) iter.Seq[E] {
	return iter.Map(Gen(0, n), func(N) E { return v })
}

// Empty creates an [iter.Seq] that yields nothing.
func Empty[E any]() iter.Seq[E] {
	return func(yield func(E) bool) {}
}

// Chan creates an [iter.Seq] from channel.
func Chan[E any, C <-chan E](c C) iter.Seq[E] {
	return func(yield func(E) bool) {
		for x := range c {
			if !yield(x) {
				break
			}
		}
	}
}

func Tcp(l net.TCPListener) iter.Seq[iter.Tuple[net.Conn, error]] {
	return func(yield func(iter.Tuple[net.Conn, error]) bool) {
		for yield(iter.MakeTuple(l.Accept())) {
		}
	}
}

func Runes(s string) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, x := range s {
			if !yield(x) {
				break
			}
		}
	}
}

func Bytes(s string) iter.Seq[byte] {
	return func(yield func(byte) bool) {
		for _, x := range []byte(s) {
			if !yield(x) {
				break
			}
		}
	}
}
