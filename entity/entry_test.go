package entity

import (
	"testing"
)

// TestEmptyEntry_1: Should return expected default values
func TestEmptyEntry_1(t *testing.T) {
	got := EmptyEntry()

	if got.ID != -1 {
		t.Errorf("result = %v, wanted %v", got, -1)
	}
	if got.Score != -1 {
		t.Errorf("result = %v, wanted %v", got, -1)
	}
	if got.Concern != "" {
		t.Errorf("result = %v, wanted %v", got, "\"\"")
	}
	if got.Grateful != "" {
		t.Errorf("result = %v, wanted %v", got, "\"\"")
	}
	if got.Learn != "" {
		t.Errorf("result = %v, wanted %v", got, "\"\"")
	}
	if got.Milestone != "" {
		t.Errorf("result = %v, wanted %v", got, "\"\"")
	}
	if got.Entered != -1 {
		t.Errorf("result = %v, wanted %v", got, -1)
	}
}
