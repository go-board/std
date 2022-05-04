package ptr

// Ref return reference of value
func Ref[T any](t T) *T { return &t }

// Default return default value of type
func Default[T any]() *T {
	var data T
	return Ref(data)
}
