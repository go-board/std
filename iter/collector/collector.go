package collector

import (
	"github.com/go-board/std/cmp"
	"github.com/go-board/std/collections/ordered"
	"github.com/go-board/std/iter"
	"github.com/go-board/std/tuple"
)

func sliceSeq[E any, S ~[]E](slice S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, x := range slice {
			if !yield(x) {
				break
			}
		}
	}
}

// Collector collect all elements in [iter.Seq] to dest type.
type Collector[E any, C any] interface {
	CollectSeq(seq iter.Seq[E])
	Collect(x E)
	Finish() C
}

// Collect collects all elements in [iter.Seq] to collector.
func Collect[E any, C any](seq iter.Seq[E], collector Collector[E, C]) C {
	collector.CollectSeq(seq)
	return collector.Finish()
}

type collectorImpl[S, E, C any] struct {
	state      S
	collectSeq func(state S, s iter.Seq[E]) S
	collect    func(state S, x E) S
	finish     func(state S) C
}

func newCollectorImpl[S, E, C any](state S, collectSeq func(state S, s iter.Seq[E]) S, collect func(state S, x E) S, finish func(state S) C) Collector[E, C] {
	return &collectorImpl[S, E, C]{state: state, collectSeq: collectSeq, collect: collect, finish: finish}
}

func (c *collectorImpl[S, E, C]) CollectSeq(s iter.Seq[E]) { c.state = c.collectSeq(c.state, s) }
func (c *collectorImpl[S, E, C]) Collect(x E)              { c.state = c.collect(c.state, x) }
func (c *collectorImpl[S, E, C]) Finish() C                { return c.finish(c.state) }

// ToSlice collects all elements in [iter.Seq] to slice.
func ToSlice[E any]() Collector[E, []E] {
	return newCollectorImpl(
		make([]E, 0),
		func(state []E, s iter.Seq[E]) []E {
			iter.ForEach(s, func(e E) { state = append(state, e) })
			return state
		},
		func(state []E, x E) []E { state = append(state, x); return state },
		func(state []E) []E { return state },
	)
}

// ToMap collects all elements in [iter.Seq] to map.
func ToMap[E any, K comparable, V any](f func(E) (K, V)) Collector[E, map[K]V] {
	return newCollectorImpl(
		make(map[K]V),
		func(state map[K]V, s iter.Seq[E]) map[K]V {
			iter.ForEach(s, func(e E) { k, v := f(e); state[k] = v })
			return state
		},
		func(state map[K]V, x E) map[K]V { k, v := f(x); state[k] = v; return state },
		func(state map[K]V) map[K]V { return state },
	)
}

// ToOrderedMapFunc collects all elements in [iter.Seq] to ordered map.
func ToOrderedMapFunc[E, K, V any](cmp func(K, K) int, f func(E) (K, V)) Collector[E, *ordered.Map[K, V]] {
	return newCollectorImpl(
		ordered.NewMap[K, V](cmp),
		func(state *ordered.Map[K, V], s iter.Seq[E]) *ordered.Map[K, V] {
			state.InsertIter(iter.Map(s, func(e E) ordered.MapEntry[K, V] { return ordered.MapEntryOf(f(e)) }))
			return state
		},
		func(state *ordered.Map[K, V], x E) *ordered.Map[K, V] {
			k, v := f(x)
			state.Insert(k, v)
			return state
		},
		func(state *ordered.Map[K, V]) *ordered.Map[K, V] { return state },
	)
}

func ToOrderedMap[E any, K cmp.Ordered, V any](f func(E) (K, V)) Collector[E, *ordered.Map[K, V]] {
	return ToOrderedMapFunc(cmp.Compare[K], f)
}

// ToOrderedSetFunc collects all elements in [iter.Seq] to ordered set.
func ToOrderedSetFunc[E, V any](cmp func(V, V) int, f func(E) V) Collector[E, *ordered.Set[V]] {
	return newCollectorImpl(
		ordered.NewSet[V](cmp),
		func(state *ordered.Set[V], s iter.Seq[E]) *ordered.Set[V] {
			state.InsertIter(iter.Map(s, f))
			return state
		},
		func(state *ordered.Set[V], x E) *ordered.Set[V] { state.Insert(f(x)); return state },
		func(state *ordered.Set[V]) *ordered.Set[V] { return state },
	)
}

func ToOrderedSet[E any, V cmp.Ordered](f func(E) V) Collector[E, *ordered.Set[V]] {
	return ToOrderedSetFunc(cmp.Compare[V], f)
}

// GroupBy collects all elements in [iter.Seq] and group by key using given function.
func GroupBy[E any, K comparable](f func(E) K) Collector[E, iter.Seq[tuple.Pair[K, iter.Seq[E]]]] {
	return newCollectorImpl(
		make(map[K][]E),
		func(state map[K][]E, s iter.Seq[E]) map[K][]E {
			iter.ForEach(s, func(e E) { k := f(e); state[k] = append(state[k], e) })
			return state
		},
		func(state map[K][]E, x E) map[K][]E { k := f(x); state[k] = append(state[k], x); return state },
		func(state map[K][]E) iter.Seq[tuple.Pair[K, iter.Seq[E]]] {
			return func(yield func(tuple.Pair[K, iter.Seq[E]]) bool) {
				for k, rs := range state {
					if !yield(tuple.MakePair(k, sliceSeq(rs))) {
						break
					}
				}
			}
		},
	)
}

// Distinct remove duplicated elements in [iter.Seq] and
// returns a new [iter.Seq] which elements yield in visit order.
func Distinct[E comparable]() Collector[E, iter.Seq[E]] {
	return newCollectorImpl(
		tuple.MakePair(make(map[E]struct{}), make([]E, 0)),
		func(state tuple.Pair[map[E]struct{}, []E], s iter.Seq[E]) tuple.Pair[map[E]struct{}, []E] {
			iter.ForEach(s, func(e E) {
				if _, ok := state.First[e]; !ok {
					state.First[e] = struct{}{}
					state.Second = append(state.Second, e)
				}
			})
			return state
		},
		func(state tuple.Pair[map[E]struct{}, []E], x E) tuple.Pair[map[E]struct{}, []E] {
			if _, ok := state.First[x]; !ok {
				state.First[x] = struct{}{}
				state.Second = append(state.Second, x)
			}
			return state
		},
		func(state tuple.Pair[map[E]struct{}, []E]) iter.Seq[E] {
			return sliceSeq(state.Second)
		},
	)
}

// DistinctFunc remove duplicated elements in [iter.Seq] and
// returns a new [iter.Seq] which elements yield in visit order.
func DistinctFunc[E any](f func(E, E) int) Collector[E, iter.Seq[E]] {
	return newCollectorImpl(
		tuple.MakePair(ordered.NewSet(f), make([]E, 0)),
		func(state tuple.Pair[*ordered.Set[E], []E], s iter.Seq[E]) tuple.Pair[*ordered.Set[E], []E] {
			iter.ForEach(s, func(e E) {
				if !state.First.Contains(e) {
					state.First.Insert(e)
					state.Second = append(state.Second, e)
				}
			})
			return state
		},
		func(state tuple.Pair[*ordered.Set[E], []E], x E) tuple.Pair[*ordered.Set[E], []E] {
			if !state.First.Contains(x) {
				state.First.Insert(x)
				state.Second = append(state.Second, x)
			}
			return state
		},
		func(state tuple.Pair[*ordered.Set[E], []E]) iter.Seq[E] {
			return sliceSeq(state.Second)
		},
	)
}

// Partition creates two [iter.Seq], split by the given predicate function.
//
// The first [iter.Seq] contains elements that satisfies the predicate.
// The second [iter.Seq]
func Partition[E any](f func(E) bool) Collector[E, tuple.Pair[iter.Seq[E], iter.Seq[E]]] {
	return newCollectorImpl(
		tuple.MakePair(make([]E, 0), make([]E, 0)),
		func(state tuple.Pair[[]E, []E], s iter.Seq[E]) tuple.Pair[[]E, []E] {
			iter.ForEach(s, func(e E) {
				if f(e) {
					state.First = append(state.First, e)
				} else {
					state.Second = append(state.Second, e)
				}
			})
			return state
		},
		func(state tuple.Pair[[]E, []E], x E) tuple.Pair[[]E, []E] {
			if f(x) {
				state.First = append(state.First, x)
			} else {
				state.Second = append(state.Second, x)
			}
			return state
		},
		func(state tuple.Pair[[]E, []E]) tuple.Pair[iter.Seq[E], iter.Seq[E]] {
			return tuple.MakePair(sliceSeq(state.First), sliceSeq(state.Second))
		},
	)
}
