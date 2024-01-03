package lazy

// Cell is a lazy value cell.
//
// Which is initialized on the first access.
type Cell[T any] struct {
	compute func() T
	val     T
	ok      bool
}

func NewCell[T any](compute func() T) *Cell[T] {
	return &Cell[T]{compute: OnceValue(compute)}
}

// Get forces the evaluation of this lazy value and returns the computed result.
func (c *Cell[T]) Get() T {
	if !c.ok {
		c.ok = true
		c.val = c.compute()
	}
	return c.val
}

func Ternary[E any](cond bool, x func() E, y func() E) E {
	if cond {
		return x()
	}
	return y()
}
