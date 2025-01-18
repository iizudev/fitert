package fitert_test

import (
	"slices"
	"testing"

	"github.com/iizudev/fitert"
)

func TestMapped(t *testing.T) {
	var (
		i = []int{0, 1, 5, 10, 20, 50}
		e = []int{0, -1, -5, -10, -20, -50}
		a []int
	)
	a = slices.Collect(fitert.Mapped(slices.Values(i), func(t int) int { return 0 - t }))
	if !slices.Equal(a, e) {
		t.Errorf("expected %v, got %v", e, a)
	}
}
