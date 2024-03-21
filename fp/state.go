package fp

type state[T any] struct{ value T }

func (s *state[T]) get() T              { return s.value }
func (s *state[T]) set(v T)             { s.value = v }
func (s *state[T]) setFunc(f func(T) T) { s.value = f(s.value) }

// UseState is a convenience function for creating a stateful function.
func UseState[T any](initial T) (func() T, func(T)) {
	s := &state[T]{value: initial}
	return s.get, s.set
}

// UseFuncState is a convenience function for creating a stateful function.
func UseFuncState[T any](initial T) (func() T, func(func(T) T)) {
	s := &state[T]{value: initial}
	return s.get, s.setFunc
}
