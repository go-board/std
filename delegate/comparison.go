package delegate

type Ordering int

type Ord[T any] Function2[T, T, int]
type Equal[T any] Function2[T, T, bool]
type NotEqual[T any] Function2[T, T, bool]
type Gt[T any] Function2[T, T, bool]
type Lt[T any] Function2[T, T, bool]
type Ge[T any] Function2[T, T, bool]
type Le[T any] Function2[T, T, bool]
