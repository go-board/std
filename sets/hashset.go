package sets

type HashSet[T comparable] struct{ store map[T]struct{} }

func NewHashSet[T comparable](slice ...T) *HashSet[T] {
	set := &HashSet[T]{store: make(map[T]struct{})}
	for _, t := range slice {
		set.store[t] = struct{}{}
	}
	return set
}

func (self *HashSet[T]) Add(targets ...T) {
	for _, t := range targets {
		self.store[t] = struct{}{}
	}
}

func (self *HashSet[T]) Remove(targets ...T) {
	for _, t := range targets {
		delete(self.store, t)
	}
}

func (self *HashSet[T]) Contains(target T) bool {
	_, ok := self.store[target]
	return ok
}

func (self *HashSet[T]) Iter(f func(T)) {
	for t := range self.store {
		f(t)
	}
}

func (self *HashSet[T]) Intersection(o *HashSet[T]) *HashSet[T] {
	return o
}
