package iter

import (
	"github.com/go-board/std/cmp"
)

type Tuple[A, B any] struct {
	Left  A
	Right B
}

func MakeTuple[A, B any](a A, b B) Tuple[A, B] { return Tuple[A, B]{Left: a, Right: b} }

// Enumerate create a new seq which yield (index, item) pair.
//
// Example:
//
//	iter.Enumerate(seq(1,2,3)) => seq: (0, 1), (1, 2), (2, 3)
func Enumerate[E any](s Seq[E]) Seq[Tuple[int, E]] {
	i := -1
	return func(yield func(Tuple[int, E]) bool) {
		s(func(e E) bool { i++; return yield(MakeTuple(i, e)) })
	}
}

// Append create a new Seq which yield each element from origin Seq
// and additional elements.
//
// Example:
//
//	iter.Append(seq(1), 2, 3) => seq: 1,2,3
//	iter.Append(seq(1,2),) => seq: 1,2
func Append[E any](s Seq[E], elems ...E) Seq[E] {
	if len(elems) == 0 {
		return s
	}
	return func(yield func(E) bool) {
		breakEarly := false
		s(func(e E) bool {
			breakEarly = !yield(e)
			return !breakEarly
		})
		if !breakEarly {
			for _, e := range elems {
				if !yield(e) {
					break
				}
			}
		}
	}
}

// TryForEach call f on each element in [Seq],
// stopping at the first error and returning that error.
func TryForEach[E any](s Seq[E], f func(E) error) (err error) {
	s(func(x E) bool { err = f(x); return err == nil })
	return
}

// ForEach call f on each element in [Seq].
func ForEach[E any](s Seq[E], f func(E)) {
	s(func(x E) bool { f(x); return true })
}

// Filter creates an iterator which uses a closure to
// determine if an element should be yielded.
//
// Filter will not stop until iterate finished.
//
// Example:
//
//	iter.Filter(seq(1,2,3), func(x int) bool {return i%2==1}) => seq: 1, 3
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

// FilterMap creates an iterator that both filters and maps.
//
// Example:
//
//	slices.FilterMap(seq(1,2,3,4,5), func(x int) (string, bool) {
//	    if x%2 == 1 {
//		    return strconv.Itoa(x), true
//	    }
//	    return "", false
//	}) => seq: "1", "3", "5"
func FilterMap[E, T any](s Seq[E], f func(E) (T, bool)) Seq[T] {
	return func(yield func(T) bool) {
		s(func(e E) bool {
			t, ok := f(e)
			if ok {
				return yield(t)
			}
			return true
		})
	}
}

// FilterZero creates an iterator which remove zero variable from Seq.
//
// Example:
//
//	iter.FilterZero(seq(true, false, false, true)) => seq: true, true
//	iter.FilterZero(seq(1, 0.0, 2., .3)) 		   => seq: 1, 2.0, 0.3
func FilterZero[E comparable](s Seq[E]) Seq[E] {
	var zero E
	return Filter(s, func(e E) bool { return e != zero })
}

// First try to find first element in [Seq],
// if no elements in [Seq], return zero var and false.
//
// Example:
//
//	iter.First(seq(1,2,3)) => 1, true
//	iter.First(seq[int]()) => 0, false
func First[E any](s Seq[E], f func(E) bool) (result E, ok bool) {
	s(func(x E) bool {
		if f(x) {
			result, ok = x, true
			return false
		}
		return true
	})
	return
}

// Last try to find last element in [Seq],
// if no elements in [Seq], return zero var and false.
//
// Example:
//
//	iter.Last(seq(1,2,3)) => 3, true
//	iter.Last(seq[int]()) => 0, false
func Last[E any](s Seq[E], f func(E) bool) (m E, ok bool) {
	s(func(e E) bool {
		if f(e) {
			m, ok = e, true
		}
		return true
	})
	return
}

// Map call f on each element in [Seq], and map each element to another type.
//
// Example:
//
//	iter.Map(seq(1,2,3), strconv.Itoa) => seq: "1", "2", "3"
func Map[E, T any](s Seq[E], f func(E) T) Seq[T] {
	return func(yield func(T) bool) {
		s(func(x E) bool { return yield(f(x)) })
	}
}

// MapWhile call f on each element on [Seq], and map each element to another type.
//
// Stopping after an initial false.
//
// Example:
//
//	iter.MapWhile(seq("1","2","e","3"), func(x string) (int, bool) {
//		i, err := strconv.Atoi(x)
//		if err != nil {
//			return 0, false
//		}
//		return i, true
//	}) => seq: 1, 2
func MapWhile[E, T any](s Seq[E], f func(E) (T, bool)) Seq[T] {
	return func(yield func(T) bool) {
		s(func(e E) bool {
			t, ok := f(e)
			return ok && yield(t)
		})
	}
}

// TryFold applies a function as long as it returns
// successfully, producing a single, final value.
//
// It takes two arguments: an initial value, and a closure with
// two arguments: an 'accumulator', and an element. The closure either
// returns successfully, with the value that the accumulator should have
// for the next iteration, or it returns failure, with an error value that
// is propagated back to the caller immediately (short-circuiting).
//
// Example:
//
//	iter.TryFold(seq(1,2,3), 1, func(acc int, item int) (int, error) { return acc * item, nil }) => 6, nil
//	iter.TryFold(seq("1", "3", "e"), 1, func(acc int, item string) (int, error) {
//		x, err := strconv.Atoi(item)
//		if err != nil {return 0, err}
//		return acc * x, nil
//	}) => 0, err
func TryFold[E, A any](s Seq[E], init A, f func(A, E) (A, error)) (res A, err error) {
	res = init
	s(func(x E) bool { res, err = f(res, x); return err == nil })
	return
}

// Fold each element into an accumulator by applying an operation,
// returning the final result.
//
// Example:
//
//	iter.Fold(seq(1,2,3), 1, func(acc int, e int) int { return acc * e }) => 6
func Fold[E, A any](s Seq[E], init A, f func(A, E) A) A {
	accum := init
	s(func(x E) bool { accum = f(accum, x); return true })
	return accum
}

// Scan returns a [Seq] containing successive accumulation values generated by applying [f]
// on each element and current accumulator value that starts with [init] value.
//
// Example:
//
//	iter.Scan(seq(1,2,3,4,5),1,func(state int, elem int) int { return state * elem }) => seq: 1,2,6,24,120
func Scan[E, A any](s Seq[E], init A, f func(state A, elem E) A) Seq[A] {
	accum := init
	return func(yield func(A) bool) {
		s(func(e E) bool {
			accum = f(accum, e)
			return yield(accum)
		})
	}
}

// Reduce reduces the elements to a single one, by repeatedly applying a reducing
// operation.
//
// Example:
//
//	iter.Reduce(seq(1,2,3), func(x, y int) int { return x * y }) => 6, true
//	iter.Reduce(seq[int](), func(x, y int) int { return x * y }) => 0, false
func Reduce[E any](s Seq[E], f func(E, E) E) (result E, hasElem bool) {
	s(func(x E) bool {
		if !hasElem {
			result, hasElem = x, true
		} else {
			result = f(result, x)
		}
		return true
	})
	return
}

// Find searches for an element of an iterator that satisfies a predicate.
func Find[E any](s Seq[E], f func(E) bool) (E, bool) { return First(s, f) }

// FindMap applies function to the elements of iterator and returns
// the first non-none result.
func FindMap[E, T any](s Seq[E], f func(E) (T, bool)) (result T, ok bool) {
	s(func(e E) bool {
		if x, o := f(e); o {
			result, ok = x, o
			return false
		}
		return true
	})
	return
}

// Contains returns true if element is found in the [Seq].
func Contains[E comparable](s Seq[E], x E) bool {
	return ContainsFunc(s, func(e E) bool { return x == e })
}

// Contains returns true if element is satisfies the given predicate in the [Seq].
func ContainsFunc[E any](s Seq[E], f func(e E) bool) bool {
	return Index(s, f) >= 0
}

// Index searches index of element which satisfying the given predicate.
func Index[E any](s Seq[E], f func(E) bool) int {
	i := -1
	s(func(x E) bool { defer func() { i++ }(); return !f(x) })
	return i
}

// All tests if every element of the iterator matches a predicate.
//
// Example:
//
//	iter.All(seq(1,2,3), func(x int) bool { return i > 0 }) => true
//	iter.Any(seq(1,2,3), func(x int) bool { return i > 2 }) => false
func All[E any](s Seq[E], f func(E) bool) bool {
	ok := true
	s(func(x E) bool { ok = f(x); return ok })
	return ok
}

// Any tests if any element of the iterator matches a predicate.
//
// Example:
//
//	iter.Any(seq(1,2,3), func(x int) bool { return i % 2 == 0 }) => true
//	iter.Any(seq(1,2,3), func(x int) bool { return i < 0 })      => false
func Any[E any](s Seq[E], f func(E) bool) (ok bool) {
	s(func(x E) bool {
		if f(x) {
			ok = true
			return false
		}
		return true
	})
	return ok
}

// Max returns the maximum element of an iterator.
//
// Example:
//
//	iter.Max(seq(1,2,3)) => 3, true
//	iter.Max(seq[int]()) => 0, false
func Max[E cmp.Ordered](s Seq[E]) (E, bool) {
	return MaxFunc(s, cmp.Compare[E])
}

// MaxFunc returns the element that gives the maximum value with respect to the
// specified comparison function.
//
// Example:
//
//	iter.MaxFunc(seq(1,2,3), cmp.Compare[int]) => 3, true
//	iter.MaxFunc(seq[int](), cmp.Compare[int]) => 0, false
func MaxFunc[E any](s Seq[E], f func(E, E) int) (E, bool) {
	return Reduce(s, func(l, r E) E {
		if f(l, r) > 0 {
			return l
		}
		return r
	})
}

// MaxByKey returns the element that gives the maximum value from the
// specified function.
//
// Example:
//
//	iter.MaxByKey(seq(user{name: "a"}, user{"name": "c"}), func(x user) string { return x.name }) => user{name: "b"}, true
//	iter.MaxByKey(seq(), func(x user) string { return x.name }) => nil, false
func MaxByKey[E any, K cmp.Ordered](s Seq[E], f func(E) K) (E, bool) {
	return MaxFunc(s, func(x E, y E) int { return cmp.Compare(f(x), f(y)) })
}

// Min returns the minimum element of an iterator.
//
// Example:
//
//	iter.Min(seq(1,2,3)) // 1, true
//	iter.Min(seq[int]()) // 0, false
func Min[E cmp.Ordered](s Seq[E]) (E, bool) {
	return MinFunc(s, cmp.Compare[E])
}

// MinFunc returns the element that gives the minimum value with respect to the
// specified comparison function.
//
// Example:
//
//	iter.MinFunc(seq(1,2,3), cmp.Compare[int]) // 1, true
//	iter.MinFunc(seq[int](), cmp.Compare[int]) // 0, false
func MinFunc[E any](s Seq[E], f func(E, E) int) (E, bool) {
	return Reduce(s, func(l, r E) E {
		if f(l, r) < 0 {
			return l
		}
		return r
	})
}

// MinByKey returns the element that gives the minimum value from the
// specified function.
//
// Example:
//
//	iter.MinByKey(seq(user{name: "a"}, user{"name": "c"}), func(x user) string { return x.name }) => user{name: "a"}, true
//	iter.MinByKey(seq(), func(x user) string { return x.name }) => nil, false
func MinByKey[E any, K cmp.Ordered](s Seq[E], f func(E) K) (E, bool) {
	return MinFunc(s, func(x E, y E) int { return cmp.Compare(f(x), f(y)) })
}

// Count consumes the iterator, counting the number of elements
// equal to the given one and returning it.
//
// Example:
//
//	iter.Count(seq(1,2,3,4,5,1,2,3), 2) // 2
//	iter.Count(seq(1,2,3,4,5,1,2,3), 5) // 1
func Count[E comparable](s Seq[E], value E) int {
	return CountFunc(s, func(e E) bool { return value == e })
}

// CountFunc consumes the iterator, counting the number of elements
// match the predicate function and returning it.
//
// Example:
//
//	iter.CountFunc(seq(1,2,3), func(x int) bool { return x % 2 == 0 }) // 1
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

// Size consumes the iterator, counting the number of iterations and returning it.
//
// Example:
//
//	iter.Size(seq(1,2,3)) // 3
func Size[E any](s Seq[E]) int {
	return CountFunc(s, func(E) bool { return true })
}

// IsSorted checks if the elements of this iterator are sorted.
//
// Example:
//
//	iter.IsSorted(seq(1, 2, 3)) // true
//	iter.IsSorted(seq(2, 1, 3)) // false
func IsSorted[E cmp.Ordered](s Seq[E]) bool {
	return IsSortedFunc(s, cmp.Compare[E])
}

// IsSortedFunc checks if the elements of this iterator are sorted using the given comparator function.
//
// Example:
//
//	iter.IsSortedFunc(seq(1, 2, 3), cmp.Compare[int]) // true
//	iter.IsSortedFunc(seq(2, 1, 3), cmp.Compare[int]) // false
func IsSortedFunc[E any](s Seq[E], f func(E, E) int) bool {
	var prev *E
	ok := true
	s(func(x E) bool {
		defer func() { prev = &x }()
		if prev != nil && f(*prev, x) > 0 {
			ok = false
			return false
		}
		return true
	})
	return ok
}

// StepBy creates an iterator starting at the same point, but stepping by
// the given amount at each iteration.
//
// Example:
//
//	iter.StepBy(seq(1,2,3,4,5), 2) // seq: 1,3,5
func StepBy[E any](s Seq[E], n int) Seq[E] {
	var i int
	return func(yield func(E) bool) {
		s(func(e E) bool {
			if i%n == 0 && !yield(e) {
				return false
			}
			i++
			return true
		})
	}
}

// Take creates an iterator that yields the first `n` elements, or fewer
// if the underlying iterator ends sooner.
//
// Example:
//
//	iter.Take(seq(1,2,3), 2) // seq: 1,2
//	iter.Take(seq(1,2,3), 5) // seq: 1,2,3
func Take[E any](s Seq[E], n int) Seq[E] {
	return func(yield func(E) bool) {
		i := -1
		s(func(x E) bool {
			i++
			return i < n && yield(x)
		})
	}
}

// TakeWhile creates an iterator that yields elements based on a predicate.
//
// Stopping after an initial `false`.
func TakeWhile[E any](s Seq[E], f func(E) bool) Seq[E] {
	return func(yield func(E) bool) {
		s(func(x E) bool { return f(x) && yield(x) })
	}
}

// Skip creates an iterator that skips the first `n` elements.
//
// Example:
//
//	iter.Skip(seq(1,2,3), 1) // seq: 2,3
//	iter.Skip(seq(1,2), 3)   // seq: none
func Skip[E any](s Seq[E], n int) Seq[E] {
	return func(yield func(E) bool) {
		i := -1
		s(func(x E) bool {
			i++
			if i >= n {
				return yield(x)
			}
			return true
		})
	}
}

// SkipWhile creates an iterator that [`skip`]s elements based on a predicate.
//
// Stop skip after an initial `false`.
func SkipWhile[E any](s Seq[E], f func(E) bool) Seq[E] {
	return func(yield func(E) bool) {
		var ok bool
		s(func(x E) bool {
			if !ok && !f(x) {
				ok = true
				return yield(x)
			}
			if ok {
				return yield(x)
			}
			return true
		})
	}
}

// Dedup removes consecutive repeated elements in the [Seq].
//
// Example:
//
//	iter.Dedup(seq(1,1,2,2,3,3)) => seq: 1,2,3
//	iter.Dedup(seq(1,2,2,3,2,2)) => seq: 1,2,3,2
func Dedup[E comparable](s Seq[E]) Seq[E] {
	return DedupFunc(s, func(x E, y E) bool { return x == y })
}

// DedupFunc removes all but the first of consecutive elements in
// the Seq satisfying a given equality relation.
//
// Example:
//
//	iter.DedupFunc(seq(1,1,2,2,3,3), func(x int, y int) bool { return x + y < 5 }) => seq: 1,3
//	iter.DedupFunc(seq(1,2,3,2,1), func(x int, y int) bool { return x != y}) => seq: 1
func DedupFunc[E any](s Seq[E], f func(E, E) bool) Seq[E] {
	return func(yield func(E) bool) {
		var prev E
		var consumed bool
		s(func(e E) bool {
			defer func() { prev = e }()
			if !consumed {
				consumed = true
				return yield(e)
			}
			if !f(prev, e) {
				return yield(e)
			}
			return true
		})
	}
}

func Flatten[E any](s Seq[Seq[E]]) Seq[E] {
	return FlatMap(s, func(x Seq[E]) Seq[E] { return x })
}

// FlatMap call f on each element in Seq which create a new type Seq by each element.
//
// Example:
//
//	iter.FlatMap(seq(1,2,3), func(x int) Seq[int] { return rangeTo(x)}) // seq: 1,1,2,1,2,3
func FlatMap[E, T any](s Seq[E], f func(E) Seq[T]) Seq[T] {
	return func(yield func(T) bool) {
		s(func(x E) bool {
			shouldContinue := true
			f(x)(func(e T) bool {
				if !yield(e) {
					shouldContinue = false
					return false
				}
				return true
			})
			return shouldContinue
		})
	}
}

// Chain takes two iterators and creates a new iterator over both in sequence.
//
// Example:
//
//	iter.Chain(seq(1,2,3), seq(4,5,6)) => seq: 1,2,3,4,5,6
func Chain[E any](x Seq[E], y Seq[E]) Seq[E] {
	return func(yield func(E) bool) {
		shouldBreak := false
		x(func(x E) bool {
			y := yield(x)
			if !y {
				shouldBreak = true
			}
			return y
		})
		if !shouldBreak {
			y(func(x E) bool {
				return yield(x)
			})
		}
	}
}

// CollectFunc call func on each element to collect it.
//
// Stopping iterate when collect returns false.
//
// collect return false means collector maybe full,
// or no need to collect more elements.
//
// Example:
//
//	res := make(chan int, 2)
//	iter.CollectFunc(seq(1,2,3), func (x int) bool {
//		select {
//		case res <- x:
//			return true
//		default:
//			return false
//		}
//	})
func CollectFunc[E any](s Seq[E], collect func(E) bool) {
	s(func(x E) bool { return collect(x) })
}

// Intersperse creates a new iterator which places a copy of `separator`
// between adjacent items of the original iterator.
//
// Example:
//
//	iter.Intersperse(seq(1,2,3), 0) // seq: 1,0,2,0,3
func Intersperse[E any](s Seq[E], sep E) Seq[E] {
	return func(yield func(E) bool) {
		first := true
		s(func(x E) bool {
			if first {
				first = false
				return yield(x)
			} else {
				return yield(sep) && yield(x)
			}
		})
	}
}

// Inspect call f on each element in [Seq].
//
// Like ForEach, Inspect will iterate all element in Seq without stopping.
// Unlike ForEach, Inspect will not consume Seq.
//
// Example:
//
//	iter.Inspect(seq(1,2,3), func(x int) { println(x) }) // seq: 1,2,3
func Inspect[E any](s Seq[E], f func(E)) Seq[E] {
	return func(yield func(E) bool) {
		s(func(e E) bool {
			f(e)
			return yield(e)
		})
	}
}
