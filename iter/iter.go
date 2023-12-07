package iter

// Seq is an iterator over sequences of individual values.
// When called as seq(yield), seq calls yield(v) for each value v in the sequence,
// stopping early if yield returns false.
//
// see: https://github.com/golang/go/issues/61897
type Seq[E any] func(yield func(E) bool)

func (s Seq[E]) ForEach(f func(E))                                  { ForEach(s, f) }
func (s Seq[E]) TryForEach(f func(E) error) error                   { return TryForEach(s, f) }
func (s Seq[E]) Map(f func(E) E) Seq[E]                             { return Map(s, f) }
func (s Seq[E]) MapWhile(f func(E) (E, bool)) Seq[E]                { return MapWhile(s, f) }
func (s Seq[E]) TryFold(init E, f func(E, E) (E, error)) (E, error) { return TryFold(s, init, f) }
func (s Seq[E]) Fold(init E, f func(E, E) E) E                      { return Fold(s, init, f) }
func (s Seq[E]) Reduce(f func(E, E) E) (E, bool)                    { return Reduce(s, f) }
func (s Seq[E]) Filter(f func(E) bool) Seq[E]                       { return Filter(s, f) }
func (s Seq[E]) FilterMap(f func(E) (E, bool)) Seq[E]               { return FilterMap(s, f) }
func (s Seq[E]) MaxFunc(f func(E, E) int) (E, bool)                 { return MaxFunc(s, f) }
func (s Seq[E]) MinFunc(f func(E, E) int) (E, bool)                 { return MinFunc(s, f) }
func (s Seq[E]) Find(f func(E) bool) (E, bool)                      { return Find(s, f) }
func (s Seq[E]) FindMap(f func(E) (E, bool)) (E, bool)              { return FindMap(s, f) }
func (s Seq[E]) Index(f func(E) bool) (int, bool)                   { return Index(s, f) }
func (s Seq[E]) All(f func(E) bool) bool                            { return All(s, f) }
func (s Seq[E]) Any(f func(E) bool) bool                            { return Any(s, f) }
func (s Seq[E]) CountFunc(f func(E) bool) int                       { return CountFunc(s, f) }
func (s Seq[E]) Size() int                                          { return Size(s) }
func (s Seq[E]) IsSortedFunc(f func(E, E) int) bool                 { return IsSortedFunc(s, f) }
func (s Seq[E]) StepBy(n int) Seq[E]                                { return StepBy(s, n) }
func (s Seq[E]) Take(n int) Seq[E]                                  { return Take(s, n) }
func (s Seq[E]) TakeWhile(f func(E) bool) Seq[E]                    { return TakeWhile(s, f) }
func (s Seq[E]) Skip(n int) Seq[E]                                  { return Skip(s, n) }
func (s Seq[E]) SkipWhile(f func(E) bool) Seq[E]                    { return SkipWhile(s, f) }
func (s Seq[E]) Chain(y Seq[E]) Seq[E]                              { return Chain(s, y) }
func (s Seq[E]) CollectFunc(collect func(E) bool)                   { CollectFunc(s, collect) }
func (s Seq[E]) Intersperse(sep E) Seq[E]                           { return Intersperse(s, sep) }
func (s Seq[E]) First(f func(E) bool) (E, bool)                     { return First(s, f) }
func (s Seq[E]) Last(f func(E) bool) (E, bool)                      { return Last(s, f) }
func (s Seq[E]) Inspect(f func(E)) Seq[E]                           { return Inspect(s, f) }
func (s Seq[E]) DedupFunc(f func(E, E) bool) Seq[E]                 { return DedupFunc(s, f) }
