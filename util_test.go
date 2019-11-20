package main

import "testing"

func TestAbs(t *testing.T) {
	a := 0.0
	b := 1.0

	pass := Abs(a-b) == Abs(b-a)

	if !pass {
		t.Error("Abs test failed")
	}
}

func TestFloatEqual(t *testing.T) {
	a := 0.0
	b := 1.0

	pass := !FloatEqual(a, b)

	if !pass {
		t.Errorf("FloatEqual: %f should not equal %f", a, b)
	}

	a = 2.000001
	b = 2.000000
	pass = FloatEqual(a, b)

	if !pass {
		t.Errorf("FloatEqual: %f should equal %f", a, b)
	}
}
