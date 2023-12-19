package main

import "testing"

func TestDoMath(t *testing.T) {
	total := doMath(10, 15, add)
	if total != 25 {
		t.Errorf("Error in doMath. Expected %d, got %d", 25, total)
	}
}

func TestAdd(t *testing.T) {
	total := add(10, 12)
	if total != 22 {
		t.Errorf("Error in add. Expected %d, but got %d", 22, total)
	}
}

func TestSubtract(t *testing.T) {
	diff := subtract(15, 5)
	if diff != 10 {
		t.Errorf("Error in subtract. Expected %d, but got %d", 10, diff)
	}
}
