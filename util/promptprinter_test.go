package util

import (
    "os"
	"bufio"
	"strconv"
	"testing"
)

// TestCheckScoreInput_1 : If below Min is provided, change score string to Min
func TestCheckScoreInput_1(t *testing.T) {
    uut := PromptPrinter{bufio.NewReader(os.Stdin), 1, 10}

	got := uut.checkScoreInput(strconv.Itoa(1 - 1))

	if got != 1 {
		t.Errorf("score = %v, wanted %v", got, 1)
	}
}

// TestCheckScoreInput_2 : If above Max is provided, change score string to Max
func TestCheckScoreInput_2(t *testing.T) {
    uut := PromptPrinter{bufio.NewReader(os.Stdin), 1, 10}

	got := uut.checkScoreInput(strconv.Itoa(10 + 1))

	if got != 10 {
		t.Errorf("score = %v, wanted %v", got, 10)
	}
}

// TestCheckScoreInput_3 : Should handle whitespace
func TestCheckScoreInput_3(t *testing.T) {
    uut := PromptPrinter{bufio.NewReader(os.Stdin), 1, 10}

	got := uut.checkScoreInput(" 5")

	if got != 5 {
		t.Errorf("score = %v, wanted 5", got)
	}
}
