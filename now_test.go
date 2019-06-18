package main

import (
	"strconv"
	"testing"
)

// TestCheckScoreInput_1 : If below Min is provided, change score string to Min
func TestCheckScoreInput_1(t *testing.T) {
	got := checkScoreInput(strconv.Itoa(Min - 1))

	if got != MinStr {
		t.Errorf("score = %s, wanted %s", got, MinStr)
	}
}

// TestCheckScoreInput_2 : If above Max is provided, change score string to Max
func TestCheckScoreInput_2(t *testing.T) {
	got := checkScoreInput(strconv.Itoa(Max + 1))

	if got != MaxStr {
		t.Errorf("score = %s, wanted %s", got, MaxStr)
	}
}

// TestCheckScoreInput_3 : Should handle whitespace
func TestCheckScoreInput_3(t *testing.T) {
	got := checkScoreInput(" 5")

	if got != "5" {
		t.Errorf("score = %s, wanted 5", got)
	}
}
