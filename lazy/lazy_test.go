package lazy_test

import (
	"testing"

	"github.com/go-board/std/lazy"
)

func TestLazy(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		lazyValue := lazy.NewLazy(func() int { return 100 })
		if lazyValue.Get() != 100 {
			t.Errorf("lazy value should be 100, but got %d", lazyValue.Get())
		}
	})
	t.Run("With", func(t *testing.T) {
		lazyValue := lazy.NewLazy(func() int { return 100 })
		lazyValue.With(func(v int) {
			if v != 100 {
				t.Errorf("lazy value should be 100, but got %d", v)
			}
		})
	})
}

func TestOnceLazy(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		initial := 100
		pInitial := &initial
		lazyValue := lazy.NewOnceLazy(func() int {
			*pInitial += 100
			return *pInitial
		})
		if lazyValue.Get() != 200 {
			t.Errorf("lazy value should be 200, but got %d", lazyValue.Get())
		}
		if lazyValue.Get() != 200 {
			t.Errorf("lazy value should be 300, but got %d", lazyValue.Get())
		}
	})
	t.Run("With", func(t *testing.T) {
		initial := 100
		pInitial := &initial
		lazyValue := lazy.NewOnceLazy(func() int {
			*pInitial += 100
			return *pInitial
		})
		lazyValue.With(func(v int) {
			if v != 200 {
				t.Errorf("lazy value should be 200, but got %d", v)
			}
		})
		lazyValue.With(func(v int) {
			if v != 200 {
				t.Errorf("lazy value should be 300, but got %d", v)
			}
		})
	})
}
