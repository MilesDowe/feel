package cmd

import (
	"testing"
)

func TestConvertToUnix_1(t *testing.T) {
	got := convertToUnix("20190622")

	expected := "1561161600"

	if got != expected {
		t.Errorf("time = %s, wanted %s", got, expected)
	}
}
