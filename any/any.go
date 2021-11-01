package any

type Any[T any] T

func Equal[A Any[T], B Any[U], T, U any]() bool { return false }
