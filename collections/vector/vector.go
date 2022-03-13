package vector

type Vector[T any] struct{ vec []T }

func New[T any]() *Vector[T] {
	return &Vector[T]{}
}

func FromSlice[T any](slice []T) *Vector[T] {
	return &Vector[T]{vec: append([]T{}, slice...)}
}

func (self *Vector[T]) Append(target T) {
	self.vec = append(self.vec, target)
}

func (self *Vector[T]) Prepend(target T) {
	self.vec = append([]T{target}, self.vec...)
}

func (self *Vector[T]) Extend(vec *Vector[T]) {
	self.vec = append(self.vec, vec.vec...)
}
