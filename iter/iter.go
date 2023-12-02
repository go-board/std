package iter

// Seq is an iterator over sequences of individual values.
// When called as seq(yield), seq calls yield(v) for each value v in the sequence,
// stopping early if yield returns false.
//
// see: https://github.com/golang/go/issues/61897
type Seq[E any] func(yield func(E) bool)
