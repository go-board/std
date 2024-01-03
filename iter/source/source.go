package source

import (
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
func RepeatTimes[E any, N core.Integer](v E, n N) iter.Seq[E] {
	return iter.Map(Gen(0, n), func(N) E { return v })
}

// Empty creates an [iter.Seq] that yields nothing.
func Empty[E any]() iter.Seq[E] {
	return func(yield func(E) bool) {}
}
