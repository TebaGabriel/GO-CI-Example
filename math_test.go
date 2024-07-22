package main

import "testing"

func TestSoma(t *testing.T) {
	total := Sum(15, 15)

	if total != 30 {
		t.Errorf("Sum result is invalid. Expected value was %d and the final result is %d", 30, total)
	}
}
