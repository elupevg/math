package math_test

import (
	"testing"

	"github.com/elupevg/math/v2"
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

func TestAdd_uint32(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		a, b, want uint32
	}{
		{2, 3, 5},
		{1, 0, 1},
		{4, 10, 14},
	}
	for _, tc := range testCases {
		got := math.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("math.Add(%d, %d): want %d, got %d", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAdd_float64(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		a, b, want float64
	}{
		{2.0, 3.0, 5.0},
		{-1.3, 1.3, 0.0},
		{-4.5, -10.2, -14.7},
	}
	for _, tc := range testCases {
		got := math.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("math.Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
