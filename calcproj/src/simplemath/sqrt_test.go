package simplemath

import "testing"

func TestSqrt1(t *testing.T) {
	var v int
	v = Sqrt(16)
	v = v - 1
	if v != 3 {
		t.Errorf("Sqrt(16) failed. Got %v, expected 4.", v)
	}
}
