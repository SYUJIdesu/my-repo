package main

import "testing"

func TestEvenOrOdd(T *testing.T)  {
	result := EvenOrOdd(10)
	if result != "even" {
		t.Errorf("expected: even, actual: %s", result)
	}
}