package iter

import (
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/operator"
)

type Tuple[A, B any] struct {
	Left  A
	Right B
}

func Enumerate[E any](s Seq[E]) Seq[Tuple[int, E]] {
	i := -1
	return func(yield func(Tuple[int, E]) bool) {
		s(func(e E) bool {
			i++
			return yield(Tuple[int, E]{Left: i, Right: e})
		})
	}
}

// TryForEach call f on each element in [Seq],
// stopping at the first error and returning that error.
func TryForEach[E any](s Seq[E], f func(E) error) (err error) {
	s(func(x E) bool {
		err = f(x)
		return err == nil
	})
	return
}

// ForEach call f on each element in Seq.
func ForEach[E any](s Seq[E], f func(E)) {
	s(func(x E) bool {
		f(x)
		return true
	})
}

// Filter remove elements do not match predicate.
func Filter[E any](s Seq[E], f func(E) bool) Seq[E] {
	return func(yield func(E) bool) {
		s(func(x E) bool {
			if f(x) {
				return yield(x)
			}
			return true
		})
	}
}

func Map[E, T any](s Seq[E], f func(E) T) Seq[T] {
	return func(yield func(T) bool) {
		s(func(x E) bool { return yield(f(x)) })
	}
}

func TryFold[E, A any](s Seq[E], init A, f func(A, E) (A, error)) (res A, err error) {
	res = init
	s(func(x E) bool {
		res, err = f(res, x)
		return err == nil
	})
	return
}

func Fold[E, A any](s Seq[E], init A, f func(A, E) A) A {
	accum := init
	s(func(x E) bool {
		accum = f(accum, x)
		return true
	})
	return accum
}

func Reduce[E any](s Seq[E], f func(E, E) E) (result E, hasElem bool) {
	s(func(x E) bool {
		if !hasElem {
			hasElem = true
			result = x
		} else {
			result = f(result, x)
		}
		return true
	})
	return
}

func Find[E any](s Seq[E], f func(E) bool) (result E, ok bool) {
	s(func(x E) bool {
		if f(x) {
			ok = true
			result = x
			return false
		}
		return true
	})
	return
}

func Index[E any](s Seq[E], f func(E) bool) (int, bool) {
	var i int = -1
	s(func(x E) bool {
		i++
		return !f(x)
	})
	return i, i >= 0
}

func All[E any](s Seq[E], f func(E) bool) bool {
	ok := true
	s(func(x E) bool {
		ok = f(x)
		return ok
	})
	return ok
}

func Any[E any](s Seq[E], f func(E) bool) bool {
	var ok bool
	s(func(x E) bool {
		if f(x) {
			ok = true
			return false
		}
		return true
	})
	return ok
}

func Max[E cmp.Ordered](s Seq[E]) (m E, hasElem bool) {
	return MaxFunc(s, cmp.Compare[E])
}

func MaxFunc[E any](s Seq[E], f func(E, E) int) (m E, hasElem bool) {
	return Reduce(s, func(l, r E) E {
		if f(l, r) > 0 {
			return l
		}
		return r
	})
}

func Min[E cmp.Ordered](s Seq[E]) (m E, hasElem bool) {
	return MinFunc(s, cmp.Compare[E])
}

func MinFunc[E any](s Seq[E], f func(E, E) int) (m E, hasElem bool) {
	return Reduce(s, func(l, r E) E {
		if f(l, r) < 0 {
			return l
		}
		return r
	})
}

func Count[E comparable](s Seq[E], value E) int {
	return CountFunc(s, func(e E) bool { return value == e })
}

func CountFunc[E any](s Seq[E], f func(E) bool) int {
	n := 0
	s(func(x E) bool {
		if f(x) {
			n++
		}
		return true
	})
	return n
}

func Size[E any](s Seq[E]) int {
	return CountFunc(s, func(E) bool { return true })
}

func IsSorted[E cmp.Ordered](s Seq[E]) bool {
	return IsSortedFunc(s, cmp.Compare[E])
}

func IsSortedFunc[E any](s Seq[E], f func(E, E) int) bool {
	var prev *E
	ok := true
	s(func(x E) bool {
		if prev == nil {
			prev = &x
		} else {
			if f(*prev, x) > 0 {
				ok = false
				return false
			}
		}
		return true
	})
	return ok
}

func Take[E any](s Seq[E], n int) Seq[E] {
	return func(yield func(E) bool) {
		i := 0
		s(func(x E) bool {
			if i < n {
				i++
				return yield(x)
			}
			return false
		})
	}
}

func Skip[E any](s Seq[E], n int) Seq[E] {
	return func(yield func(E) bool) {
		i := 0
		s(func(x E) bool {
			if i >= n {
				return yield(x)
			}
			i++
			return true
		})
	}
}

func Distinct[E comparable](s Seq[E]) Seq[E] {
	m := make(map[E]struct{})
	return func(yield func(E) bool) {
		s(func(x E) bool {
			if _, ok := m[x]; !ok {
				m[x] = struct{}{}
				return yield(x)
			}
			return true
		})
	}
}

func Flatten[E any](s Seq[Seq[E]]) Seq[E] {
	return FlatMap(s, operator.Identify[Seq[E]])
}

func FlatMap[E, T any](s Seq[E], f func(E) Seq[T]) Seq[T] {
	return func(yield func(T) bool) {
		s(func(x E) bool {
			shouldCountine := true
			f(x)(func(e T) bool {
				if !yield(e) {
					shouldCountine = false
					return false
				}
				return true
			})
			return shouldCountine
		})
	}
}
