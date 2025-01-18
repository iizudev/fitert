package fitert_test

import (
	"slices"
	"testing"

	"github.com/iizudev/fitert"
)

func TestFiltered(t *testing.T) {
	var (
		i = []int{0, 40, 20, 10, 15, 33}
		e = []int{40, 20, 15, 33}
		a []int
	)
	a = slices.Collect(
		fitert.Filtered(
			slices.Values(i),
			func(t *int) bool { return *t > 10 },
		),
	)
	if !slices.Equal(a, e) {
		t.Errorf("expected: %v, got %v", e, a)
	}
}
