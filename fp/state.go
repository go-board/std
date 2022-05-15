package fp

// UseState is a convenience function for creating a stateful function.
func UseState[T any](initial T) (func() T, func(T)) {
	state := initial
	getter := func() T { return state }
	setter := func(newState T) { state = newState }
	return getter, setter
}

// UseFuncState is a convenience function for creating a stateful function.
func UseFuncState[T any](initial T) (func() T, func(func(T) T)) {
	state := initial
	getter := func() T { return state }
	setter := func(f func(prevState T) T) { state = f(state) }
	return getter, setter
}
