package cmp

import (
	"github.com/go-board/std/optional"
)

type Order int // -1, 0, 1

func MakeOrder(o int) Order {
	switch {
	case o < 0:
		return OrderLess
	case o > 0:
		return OrderGreater
	default:
		return OrderEqual
	}
}

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

type Comparator[A any] interface {
	Cmp(lhs A, rhs A) int
	Lt(lhs A, rhs A) bool
	Le(lhs A, rhs A) bool
	Gt(lhs A, rhs A) bool
	Ge(lhs A, rhs A) bool
	Eq(lhs A, rhs A) bool
	Ne(lhs A, rhs A) bool
}

type ComparatorFunc[A any] func(A, A) int

func (f ComparatorFunc[A]) Cmp(lhs A, rhs A) int {
	return f(lhs, rhs)
}

func (f ComparatorFunc[A]) Lt(lhs A, rhs A) bool {
	return f.Cmp(lhs, rhs) < 0
}

func (f ComparatorFunc[A]) Le(lhs A, rhs A) bool {
	return f.Cmp(lhs, rhs) <= 0
}

func (f ComparatorFunc[A]) Gt(lhs A, rhs A) bool {
	return f.Cmp(lhs, rhs) > 0
}

func (f ComparatorFunc[A]) Ge(lhs A, rhs A) bool {
	return f.Cmp(lhs, rhs) >= 0
}

func (f ComparatorFunc[A]) Eq(lhs A, rhs A) bool {
	return f.Cmp(lhs, rhs) == 0
}

func (f ComparatorFunc[A]) Ne(lhs A, rhs A) bool {
	return f.Cmp(lhs, rhs) != 0
}

func MakeComparator[A Ordered]() Comparator[A] {
	return ComparatorFunc[A](Compare[A])
}

func MakeComparatorFunc[A any](f func(A, A) int) Comparator[A] {
	return ComparatorFunc[A](f)
}
