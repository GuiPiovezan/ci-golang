package main

import "testing"

func TestSoma(t *testing.T) {
	result := Soma(10, 10)
	expected := 20

	if result != expected {
		t.Errorf("Soma(10, 10) = %d; want %d", result, expected)
	}
}