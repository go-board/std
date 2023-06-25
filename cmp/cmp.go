package cmp

import (
	"github.com/go-board/std/optional"
)

type Order int // -1, 0, 1

func (o Order) IsEq() bool { return o == 0 }
func (o Order) IsNe() bool { return o != 0 }
func (o Order) IsLt() bool { return o < 0 }
func (o Order) IsLe() bool { return o <= 0 }
func (o Order) IsGe() bool { return o >= 0 }
func (o Order) IsGt() bool { return o > 0 }
func (o Order) String() string {
	switch {
	case o < 0:
		return "Less"
	case o > 0:
		return "Greater"
	default:
		return "Equal"
	}
}

const (
	// OrderLess is an Ordering where a compared value is less than another
	OrderLess = Order(-1)

	// OrderEqual is an Ordering where a compared value is equal to another
	OrderEqual = Order(0)

	// OrderGreater is an Ordering where a compared value is greater than another
	OrderGreater = Order(1)
)

// PartialEq is a type that represents a partial equality comparison.
//
//	see: [PartialEq](https://en.wikipedia.org/wiki/Partial_equivalence_relation)
type PartialEq[A any] interface {
	Eq(A) bool
	Ne(A) bool
}

// Eq is a type that represents an equality comparison.
//
//	see: [Eq](https://en.wikipedia.org/wiki/Equivalence_relation)
type Eq[A any] interface {
	PartialEq[A]
}

// PartialOrd is a type that represents a partially ordered value.
//
//	see: [PartialOrd](https://en.wikipedia.org/wiki/Partially_ordered_set)
type PartialOrd[A any] interface {
	PartialCmp(A) optional.Optional[Order]
	Lt(A) bool
	Le(A) bool
	Gt(A) bool
	Ge(A) bool
}

// Ord is a type that represents an ordered value.
//
//	see: [Ord](https://en.wikipedia.org/wiki/Total_order)
type Ord[A any] interface {
	Eq[A]
	PartialOrd[A]
	Cmp(A) Order
}
