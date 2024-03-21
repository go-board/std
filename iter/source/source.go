package source

import (
	"github.com/go-board/std/constraints"
	"github.com/go-board/std/iter"
	"github.com/go-board/std/tuple"
)

// Bytes create an [iter.Seq] of byte from underlying string type.
func Bytes[E ~string](s E) iter.Seq[byte] {
	return func(yield func(byte) bool) {
		for _, x := range []byte(s) {
			if !yield(x) {
				return
			}
		}
	}
}

// Runes create an [iter.Seq] of rune from underlying string type.
func Runes[E ~string](s E) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, x := range s {
			if !yield(x) {
				return
			}
		}
	}
}

// Chan create an [iter.Seq] from readable channel.
func Chan[E any](c <-chan E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for x := range c {
			if !yield(x) {
				return
			}
		}
	}
}

// Map create an [iter.Seq] of k-v pair from map type.
//
// The pair order is unordered.
func Map[K comparable, V any, M ~map[K]V](m M) iter.Seq[tuple.Pair[K, V]] {
	return func(yield func(tuple.Pair[K, V]) bool) {
		for k, v := range m {
			if !yield(tuple.MakePair(k, v)) {
				return
			}
		}
	}
}

// Variadic creates an [iter.Seq] from variadic elements.
//
// Example:
//
//	source.Variadic(1,2,3) => seq: 1,2,3
func Variadic[E any](elems ...E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, elem := range elems {
			if !yield(elem) {
				break
			}
		}
	}
}

// Range1 creates an [iter.Seq] that yield number in range.
//
// The range start from 0, but not includes [stop].
//
// If `stop` less-equal to 0, yield nothing as [Empty].
//
// Example:
//
//	source.Range1(0) => seq:
//	source.Range1(2) => seq: 0,1
func Range1[N constraints.Integer](stop N) iter.Seq[N] {
	return Range3(0, stop, 1)
}

// Range2 creates an [iter.Seq] that yield number in range.
//
// The range includes [start], but not includes [stop].
//
// If `stop` less-equal than `start`, yield nothing as [Empty].
//
// Example:
//
//	source.Range2(0, 3) => seq: 0,1,2
//	source.Range2(2, 5) => seq: 2,3,4
func Range2[N constraints.Integer](start N, stop N) iter.Seq[N] {
	return Range3(start, stop, 1)
}

// Range3 creates an [iter.Seq] that yield number in range every [step].
//
// The range includes [start], but not includes [stop].
//
// If `stop` less-equal than `start` or `step` less-than 1, yield nothing as [Empty].
//
// Example:
//
//	source.Range3(0, 10, 3) => seq: 0,3,6,9
//	source.Range3(2, 5, 0) => seq:
func Range3[N constraints.Integer](start N, stop N, step N) iter.Seq[N] {
	return func(yield func(N) bool) {
		if start >= stop || step < 1 {
			return
		}
		for i := start; i < stop; i += step {
			if !yield(i) {
				break
			}
		}
	}
}

// Once creates an [iter.Seq] that yields an element exactly once.
func Once[E any](v E) iter.Seq[E] {
	return func(yield func(E) bool) { yield(v) }
}

// Repeat creates an [iter.Seq] that endlessly repeats a single element.
func Repeat[E any](v E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for yield(v) {
		}
	}
}

// RepeatTimes creates an [iter.Seq] that repeats a single element a given number of times.
//
// Example:
//
//	source.RepeatTimes(1, 3) => seq: 1,1,1
func RepeatTimes[E any, N constraints.Integer](v E, n N) iter.Seq[E] {
	return iter.Map(Range1(n), func(N) E { return v })
}

// Empty creates an [iter.Seq] that yields nothing.
func Empty[E any]() iter.Seq[E] {
	return func(yield func(E) bool) {}
}
