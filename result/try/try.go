package try

import (
	"errors"
	"fmt"
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

func Errorf(format string, args ...any) error {
	return &Error{
		cause: fmt.Errorf(format, args...),
		frame: frame(3),
	}
}

func frame(skip int) runtime.Frame {
	pcs := [1]uintptr{}
	runtime.Callers(skip, pcs[:])
	frame, _ := runtime.CallersFrames(pcs[:]).Next()
	return frame
}

func raise(e error) {
	// pcs := [1]uintptr{}
	// runtime.Callers(3, pcs[:])
	// frame, _ := runtime.CallersFrames(pcs[:]).Next()
	panic(&Error{cause: e, frame: frame(4)})
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

// ErrNilValue indicate both co-value and err are nil
var ErrNilValue = errors.New("err: nil value")

// HandleNillable process cornor case of error handling.
// Patch fucking stupid lang design.
//
//  1. if err is not nil, return val & err
//  2. if val is nil, return nil & ErrNilValue
//  3. return val & nil
func HandleNillable[T any](val *T, err error) (*T, error) {
	if err != nil {
		return val, err
	}
	if val == nil {
		return nil, ErrNilValue
	}
	return val, nil
}
