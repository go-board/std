package lazy

// LazyCell is a lazy value.
//
// Which is initialized on the first access.
type LazyCell[T any] struct {
	compute func() T
	cell    OnceCell[T]
}

func NewLazyCell[T any](compute func() T) *LazyCell[T] {
	return &LazyCell[T]{compute: compute, cell: OnceCell[T]{}}
}

// Get forces the evaluation of this lazy value and returns the computed result.
func (self *LazyCell[T]) Get() T { return self.cell.GetOrInit(self.compute) }

// Value is lazy evaluated data
type Value[T any] interface{ Eval() T }

type ptrValue[T any] struct{ ptr *T }

func (p ptrValue[T]) Eval() T { return *p.ptr }

func PtrOf[T any](v *T) Value[T] { return ptrValue[T]{ptr: v} }

type fnValue[T any] struct{ inner func() T }

func (f fnValue[T]) Eval() T { return f.inner() }

func FuncOf[T any](f func() T) Value[T] { return fnValue[T]{inner: f} }

type value[T any] struct{ inner T }

func (v value[T]) Eval() T { return v.inner }

func ValueOf[T any](v T) Value[T] { return value[T]{inner: v} }

func Ternary[T any](ok bool, lhs Value[T], rhs Value[T]) T {
	if ok {
		return lhs.Eval()
	}
	return rhs.Eval()
}

func TernaryValue[T any](ok Value[bool], lhs Value[T], rhs Value[T]) T {
	if ok.Eval() {
		return lhs.Eval()
	}
	return rhs.Eval()
}
