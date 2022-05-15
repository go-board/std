package cmp

import (
	"fmt"

	"github.com/go-board/std/cond"
	"github.com/go-board/std/optional"
)

// private is a guard for Ordering safety.
// when Ordering using this, it can't be implemented out this package.
type private struct{}

// Ordering is a type that represents the ordering of two values.
type Ordering interface {
	ordering(private)
	IsEq() bool
	IsNe() bool
	IsLt() bool
	IsLe() bool
	IsGt() bool
	IsGe() bool
	fmt.Stringer
}

const (
	// Less is an Ordering where a compared value is less than another
	Less = order(-1)

	// Equal is an Ordering where a compared value is equal to another
	Equal = order(0)

	// Greater is an Ordering where a compared value is greater than another
	Greater = order(1)
)

type order int // -1, 0, 1

func (o order) ordering(private) {}
func (o order) IsEq() bool       { return o == 0 }
func (o order) IsNe() bool       { return o != 0 }
func (o order) IsLt() bool       { return o < 0 }
func (o order) IsLe() bool       { return o <= 0 }
func (o order) IsGe() bool       { return o >= 0 }
func (o order) IsGt() bool       { return o > 0 }
func (o order) String() string {
	return cond.Ternary(o == 0, "Equal", cond.Ternary(o < 0, "Less", "Greater"))
}

// PartialEq is a type that represents a partial equality comparison.
//  see: [PartialEq](https://en.wikipedia.org/wiki/Partial_equivalence_relation)
type PartialEq[A any] interface {
	Eq(A) bool
	Ne(A) bool
}

// Eq is a type that represents an equality comparison.
//  see: [Eq](https://en.wikipedia.org/wiki/Equivalence_relation)
type Eq[A any] interface {
	PartialEq[A]
}

// PartialOrd is a type that represents a partially ordered value.
//  see: [PartialOrd](https://en.wikipedia.org/wiki/Partially_ordered_set)
type PartialOrd[A any] interface {
	PartialCmp(A) optional.Optional[Ordering]
	Lt(A) bool
	Le(A) bool
	Gt(A) bool
	Ge(A) bool
}

// Ord is a type that represents an ordered value.
//  see: [Ord](https://en.wikipedia.org/wiki/Total_order)
type Ord[A any] interface {
	Eq[A]
	PartialOrd[A]
	Cmp(A) Ordering
}

// OrdFunc is a function that returns an Ordering.
type OrdFunc[A any] func(lhs A, rhs A) Ordering

// EqFunc is a function that returns a bool.
type EqFunc[A any] func(lhs A, rhs A) bool
