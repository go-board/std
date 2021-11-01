package sets

type HashSet[T comparable] struct{ store map[T]struct{} }

func NewHashSet[T comparable](slice ...T) *HashSet[T] {
	set := &HashSet[T]{store: make(map[T]struct{})}
	for _, t := range slice {
		set.store[t] = struct{}{}
	}
	return set
}

func (s *HashSet[T]) Add(targets ...T) {
	for _, t := range targets {
		s.store[t] = struct{}{}
	}
}

func (s *HashSet[T]) Remove(targets ...T) {
	for _, t := range targets {
		delete(s.store, t)
	}
}

func (s *HashSet[T]) Contains(target T) bool {
	_, ok := s.store[target]
	return ok
}

func (s *HashSet[T]) Iter(f func(T)) {
	for t := range s.store {
		f(t)
	}
}

func (s *HashSet[T]) Intersection(o *HashSet[T]) *HashSet[T] {
	return o
}
