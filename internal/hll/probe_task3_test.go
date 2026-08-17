package hll

import (
	"math"
	"testing"
)

func TestProbeLargeRangeCorrectionUsesNaturalLog(t *testing.T) {
	h := &HLL{p: 4, m: 16, registers: make([]uint8, 16)}
	for i := range h.registers {
		h.registers[i] = 56
	}
	m := float64(h.m)
	raw := h.alpha() * m * m / (m * math.Pow(2, -56))
	twoPow64 := math.Ldexp(1, wordBits)
	want := -twoPow64 * math.Log(1-raw/twoPow64)
	got := h.Estimate()
	if math.Abs(got-want)/want > 1e-12 {
		t.Fatalf("large-range correction = %.0f, want %.0f", got, want)
	}
}
