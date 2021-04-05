package main

import "testing"

func TestMultiply(t *testing.T) {
	pr := multiply(4,5)
	if pr != 20 {
		t.Error("Expected 20 but got ", pr)
	}
}
