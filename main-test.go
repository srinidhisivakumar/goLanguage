package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 6
	if result != expected {
		t.Fatalf("Add(2, 3) will result %d, but we got %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 2)
	expected := 3
	if result != expected {
		t.Fatalf("Subtract(5, 2) will result %d, but we got %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	expected := 5
	if result != expected {
		t.Fatalf("Multiply(2, 3) will result %d, but we got %d", result, expected)
	}
}
