package cmd

import (
	"strconv"
	"testing"
)

// TestCheckScoreInput_1 : If below Min is provided, change score string to Min
func TestCheckScoreInput_1(t *testing.T) {
	got := checkScoreInput(strconv.Itoa(Min - 1))

	if got != Min {
		t.Errorf("score = %v, wanted %v", got, Min)
	}
}

// TestCheckScoreInput_2 : If above Max is provided, change score string to Max
func TestCheckScoreInput_2(t *testing.T) {
	got := checkScoreInput(strconv.Itoa(Max + 1))

	if got != Max {
		t.Errorf("score = %v, wanted %v", got, Max)
	}
}

// TestCheckScoreInput_3 : Should handle whitespace
func TestCheckScoreInput_3(t *testing.T) {
	got := checkScoreInput(" 5")

	if got != 5 {
		t.Errorf("score = %v, wanted 5", got)
	}
}
