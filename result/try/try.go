package try

import (
	"runtime"
	"strconv"
)

// Error represents an error that occurred during a Try operation.
type Error struct {
	cause error
	frame runtime.Frame
}

func (e *Error) Error() string {
	return e.frame.File + ":" + strconv.Itoa(e.frame.Line) + ": " + e.cause.Error()
}

// Frame returns the frame of the caller of the function that called Try.
func (e *Error) Frame() runtime.Frame { return e.frame }

func (e *Error) Unwrap() error { return e.cause }

func raise(e error) {
	pcs := [1]uintptr{}
	runtime.Callers(3, pcs[:])
	frame, _ := runtime.CallersFrames(pcs[:]).Next()
	panic(&Error{cause: e, frame: frame})
}

func Try(err error) {
	if err != nil {
		raise(err)
	}
}

func Try1[A any](a A, err error) A {
	if err != nil {
		raise(err)
	}
	return a
}

func Catch(errRef *error) {
	if err := recover(); err != nil {
		switch err := err.(type) {
		case *Error:
			*errRef = err.cause
		case error:
			*errRef = err
		default:
			panic(err)
		}
	}
}

func CatchFunc(fn func(e *Error)) {
	if err := recover(); err != nil {
		if e, ok := err.(*Error); ok {
			fn(e)
		} else {
			panic(err)
		}
	}
}
