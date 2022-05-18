package delegate

// Deprecated: Use cmp.Ordering instead.
type Ordering int

// Deprecated: Use cmp.Ord instead.
type Ord[T any] Function2[T, T, int]

// Deprecated: Use cmp.Eq instead.
type Equal[T any] Function2[T, T, bool]

// Deprecated: Use cmp.Eq instead.
type NotEqual[T any] Function2[T, T, bool]

// Deprecated: Use cmp.PartialOrd instead.
type Gt[T any] Function2[T, T, bool]

// Deprecated: Use cmp.PartialOrd instead.
type Lt[T any] Function2[T, T, bool]

// Deprecated: Use cmp.PartialOrd instead.
type Ge[T any] Function2[T, T, bool]

// Deprecated: Use cmp.PartialOrd instead.
type Le[T any] Function2[T, T, bool]
