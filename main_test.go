package main

import "testing"

func TestCalculateWith2 (t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestCalculateWith4 (t *testing.T) {
	if Calculate(4) != 5 {
		t.Error("Expected 4 + 2 to equal 6")
	}
}
