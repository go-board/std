package slice

import (
	"github.com/go-board/std/optional"
	"github.com/go-board/std/sets"
)

func Map[T, U any](slice []T, f func(T) U) []U {
	rs := make([]U, 0, len(slice))
	for _, t := range slice {
		rs = append(rs, f(t))
	}
	return rs
}

func MapFilter[T, U any](slice []T, f func(T) *optional.Optional[U]) []U {
	rs := make([]U, 0, len(slice))
	for _, t := range slice {
		x := f(t)
		if x.IsSome() {
			rs = append(rs, x.Value())
		}
	}
	return rs
}

func Filter[T any](slice []T, f func(T) bool) []T {
	rs := make([]T, 0)
	for _, t := range slice {
		if f(t) {
			rs = append(rs, t)
		}
	}
	return rs
}

func Reduce[T, A any](slice []T, initial A, f func(T, A) A) A {
	a := initial
	for _, t := range slice {
		a = f(t, a)
	}
	return a
}

func Step[T any](slice []T, step int) []T {
	rs := make([]T, 0)
	i := 0
	for i < len(slice) {
		rs = append(rs, slice[i])
		i += step
	}
	return rs
}

func Unique[T comparable](slice []T) []T {
	s := sets.NewHashSet[T]()
	rs := make([]T, 0)
	for _, t := range slice {
		if !s.Contains(t) {
			s.Add(t)
			rs = append(rs, t)
		}
	}
	return rs
}

func First[T any](slice []T, predicate func(T) bool) (*T, bool) {
	for _, t := range slice {
		if predicate(t) {
			return &t, true
		}
	}
	return nil, false
}

func Last[T any](slice []T, predicate func(T) bool) (*T, bool) {
	for i := len(slice) - 1; i > 0; i-- {
		t := slice[i]
		if predicate(t) {
			return &t, true
		}
	}
	return nil, false
}
