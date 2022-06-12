package clone

type Cloneable[T any] interface{ Clone() T }

func Clone[T Cloneable[T]](o T) T { return o.Clone() }

type Cloner[T any] func(T) T
