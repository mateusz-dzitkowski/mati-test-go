package main

import (
	"cmp"
	"fmt"
	"testing"
)

func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func TestMinBasic(t *testing.T) {
	ans := Min(2, -2)
	if ans != -2 {
		t.Errorf("Min(2, -2) = %d, want -2", ans)
	}
}

func TestMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b, want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d", tt.a, tt.b), func(t *testing.T) {
			ans := Min(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("Min(2, -2) = %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkMin(b *testing.B) {
	for range b.N {
		Min(1, 2)
	}
}
