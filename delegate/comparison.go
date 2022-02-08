package delegate

type Comparison[T any] func(T, T) int

type Equal[T any] func(T, T) bool
