package math_test

import (
	"testing"

	"github.com/elupevg/math"
)

func TestAdd_int(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		a, b, want int
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{-4, -10, -14},
	}
	for _, tc := range testCases {
		got := math.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("math.Add(%d, %d): want %d, got %d", tc.a, tc.b, tc.want, got)
		}
	}
}
