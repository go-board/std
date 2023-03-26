package operator_test

import (
	"reflect"
	"testing"

	"github.com/go-board/std/operator"
)

func TestOperatorEq(t *testing.T) {
	reflect.ValueOf(t).MethodByName("").Call(nil)
	operator.Eq(1, 2)
}
