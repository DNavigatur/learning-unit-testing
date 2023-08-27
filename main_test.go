package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(10, 6)
	if total != 16 {
		t.Errorf("Sum was incorrect, got: %d., want: %d.", total, 16)
	}
}
func TestSubtract(t *testing.T) {
	total := Subtract(10, 6)
	if total != 4 {
		t.Errorf("Sum was incorrect, got: %d., want: %d.", total, 4)
	}
}
func TestDoMath(t *testing.T) {
	total := DoMath(19, 6, Add)
	if total != 25 {
		t.Errorf("Sum was incorrect, got: %d., want: %d.", total, 4)
	}
}
