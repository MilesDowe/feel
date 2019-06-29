package util

import (
    "testing"
)

func TestMean(t *testing.T) {
    got := Mean([]float64{1, 2, 3, 4, 5})

    expected := 3.0

    if got != expected {
        t.Errorf("Mean = %v, expected = %v", got, expected)
    }
}

func TestStdDev(t *testing.T) {
    got := StdDev([]float64{1, 2, 3, 4, 5})

    expected := 1.4142135623730951

    if got != expected {
        t.Errorf("Standard deviation = %v, expected = %v", got, expected)
    }
}
