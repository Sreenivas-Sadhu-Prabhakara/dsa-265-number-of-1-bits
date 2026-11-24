package main

import (
	"math/bits"
	"testing"
)

func TestOnesCount(t *testing.T) {
	cases := []struct {
		in   uint32
		want int
	}{
		{0b1011, 3},
		{0b10000000, 1},
		{0xFFFFFFFF, 32},
		{0, 0},
		{1, 1},
	}
	for _, c := range cases {
		if got := OnesCount(c.in); got != c.want {
			t.Errorf("OnesCount(%b) = %d, want %d", c.in, got, c.want)
		}
	}
}

// Property: agree with the stdlib across a sweep of values.
func TestAgainstStdlib(t *testing.T) {
	for v := uint32(0); v < 10000; v++ {
		if OnesCount(v) != bits.OnesCount32(v) {
			t.Fatalf("mismatch at %d", v)
		}
	}
}
