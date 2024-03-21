package fp

type Either[L, R any] struct {
	left   L
	right  R
	isLeft bool
}

// Left create an [Either] from left value.
func Left[L, R any](l L) Either[L, R] { return Either[L, R]{left: l, isLeft: true} }

// Right create an [Either] from right value.
func Right[L, R any](r R) Either[L, R] { return Either[L, R]{right: r, isLeft: false} }

// IsLeft test if left value.
func (e Either[L, R]) IsLeft() bool { return e.isLeft }

// IsRight test if right value.
func (e Either[L, R]) IsRight() bool { return !e.isLeft }

// InspectLeft call f if is left value.
func (e Either[L, R]) InspectLeft(f func(L)) {
	if e.IsLeft() {
		f(e.left)
	}
}

// InspectRight call f if is right value.
func (e Either[L, R]) InspectRight(f func(R)) {
	if e.IsRight() {
		f(e.right)
	}
}

// ToLeft convert right value to left value.
func (e Either[L, R]) ToLeft(f func(R) L) Either[L, R] {
	if e.IsLeft() {
		return Left[L, R](e.left)
	}
	return Left[L, R](f(e.right))
}

// ToRight convert left value to right value.
func (e Either[L, R]) ToRight(f func(L) R) Either[L, R] {
	if e.IsRight() {
		return Right[L](e.right)
	}
	return Right[L](f(e.left))
}

// Left unwrap to left value, if not, panic.
func (e Either[L, R]) Left() L {
	if e.IsLeft() {
		return e.left
	}
	panic("Left on Right value")
}

func (e Either[L, R]) LeftOr(v L) L {
	if e.IsLeft() {
		return e.left
	}
	return v
}

func (e Either[L, R]) LeftOrElse(f func() L) L {
	if e.IsLeft() {
		return e.left
	}
	return f()
}

// Right unwrap to right value, if not, panic.
func (e Either[L, R]) Right() R {
	if e.IsRight() {
		return e.right
	}
	panic("Right on Left value")
}

func (e Either[L, R]) RightOr(v R) R {
	if e.IsRight() {
		return e.right
	}
	return v
}

func (e Either[L, R]) RightOrElse(f func() R) R {
	if e.IsRight() {
		return e.right
	}
	return f()
}
