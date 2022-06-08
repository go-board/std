package maps_test

import (
	"strings"
	"testing"

	"github.com/frankban/quicktest"
	"github.com/go-board/std/maps"
)

func TestDefaultHashMap_Has(t *testing.T) {
	m := maps.NewDefaultHashMap(func(k string) string { return k })
	m.Set("a", "a")
	a := quicktest.New(t)
	a.Assert(m.Has("a"), quicktest.IsTrue)
	a.Assert(m.Has("b"), quicktest.IsFalse)
}

func TestDefaultHashMap_Get(t *testing.T) {
	m := maps.NewDefaultHashMap(func(k string) string { return k })
	a := quicktest.New(t)
	a.Assert(m.Get("a"), quicktest.Equals, "a")
}

func TestDefaultHashMap_Set(t *testing.T) {
	m := maps.NewDefaultHashMap(func(k string) string { return k })
	m.Set("a", "b")
	a := quicktest.New(t)
	a.Assert(m.Get("a"), quicktest.Equals, "b")
	a.Assert(m.Get("b"), quicktest.Equals, "b")
}

func TestDefaultHashMap_Range(t *testing.T) {
	testCases := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	m := maps.NewDefaultHashMap(func(k string) string { return k })
	for key, value := range testCases {
		m.Set(key, value)
	}
	a := quicktest.New(t)
	m.Range(func(key, val string) {
		a.Assert(val, quicktest.Equals, strings.Repeat(key, 2))
	})
}
