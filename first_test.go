package fitert_test

import (
	"slices"
	"testing"

	"github.com/iizudev/fitert"
)

func TestFirst(t *testing.T) {
	t.Run("non-empty iter.Seq", func(t *testing.T) {
		var (
			ev  = "Hello, World!"
			eok = true
			v   string
			ok  bool
		)
		v, ok = fitert.First(slices.Values([]string{ev}))
		if v != ev || ok != eok {
			t.Errorf(
				"expected: v = %v and ok = %v, got v = %v and ok = %v",
				ev, eok, v, ok,
			)
		}
	})
	t.Run("empty iter.Seq", func(t *testing.T) {
		var (
			e = false
			a bool
		)
		_, a = fitert.First[any](slices.Values([]any{}))
		if a != e {
			t.Errorf("expected: %v, got %v", e, a)
		}
	})
}
