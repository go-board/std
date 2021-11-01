package vector

type Vector[T any] struct{ vec []T }

func New[T any]() *Vector[T] {
	return &Vector[T]{}
}

func FromSlice[T any](slice []T) *Vector[T] {
	return &Vector[T]{vec: append([]T{}, slice...)}
}

func (v *Vector[T]) Append(target T) {
	v.vec = append(v.vec, target)
}

func (v *Vector[T]) Prepend(target T) {
	v.vec = append([]T{target}, v.vec...)
}

func (v *Vector[T]) Extend(vec *Vector[T]) {
	v.vec = append(v.vec, vec.vec...)
}
