package util

import (
	"testing"
)

// TestCheckInput_1 : should return true if 'y'
func TestCheckInput_1(t *testing.T) {
	got := CheckInput("y")

	if got != true {
		t.Errorf("result = %v, wanted %v", got, true)
	}
}

// TestCheckInput_2 : should return true despite case
func TestCheckInput_2(t *testing.T) {
	got := CheckInput("Y")

	if got != true {
		t.Errorf("result = %v, wanted %v", got, true)
	}
}

// TestCheckInput_3 : should return true if 'y'
func TestCheckInput_3(t *testing.T) {
	got := CheckInput("yes")

	if got != true {
		t.Errorf("result = %v, wanted %v", got, true)
	}
}

// TestCheckInput_4 : should return true despite case
func TestCheckInput_4(t *testing.T) {
	got := CheckInput("YES")

	if got != true {
		t.Errorf("result = %v, wanted %v", got, true)
	}
}
