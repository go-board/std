package optional

func And[T, U any](o *Optional[T], u *Optional[U]) *Optional[U] {
	if o.IsSome() {
		return u
	}
	return None[U]()
}

func AndThen[T, U any](o *Optional[T], f func(T) *Optional[U]) *Optional[U] {
	if o.IsSome() {
		return f(o.Value())
	}
	return None[U]()
}

func Map[T, U any](o *Optional[T], f func(T) U) *Optional[U] {
	if o.IsSome() {
		return Some[U](f(o.Value()))
	}
	return None[U]()
}

func MapOr[T, U any, F func(T) U](o *Optional[T], f F, defValue U) U {
	if o.IsSome() {
		return f(o.Value())
	}
	return defValue
}

func MapOrElse[T, U any, F func(T) U, D func() U](o *Optional[T], f F, d D) U {
	if o.IsSome() {
		return f(o.Value())
	}
	return d()
}
