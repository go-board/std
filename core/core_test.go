package core

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	x := Slice[int]{}
	y := Slice[string]{}
	rx := reflect.TypeOf(x)
	ry := reflect.TypeOf(y)
	if rx.Name() != "Slice[int]" {
		t.Fail()
	}
	if ry.Name() != "Slice[string]" {
		t.Fail()
	}
	t.Logf("%+v, %s\n", rx, rx.Name())
	t.Logf("%+v, %s\n", ry, ry.Name())
}
