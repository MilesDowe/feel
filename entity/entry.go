package entity

import (
	"fmt"
)

// Entry : a user's `feel` entry
type Entry struct {
	ID        int
	Score     int
	Concern   string
	Grateful  string
	Learn     string
	Milestone string
	Entered   int64
}

// EmptyEntry : makes a new entry with default fields
func EmptyEntry() Entry {
	return Entry{
		ID:        -1,
		Score:     -1,
		Concern:   "",
		Grateful:  "",
		Learn:     "",
		Milestone: "",
		Entered:   -1,
	}
}

// EntryWithUserInput : makes a new entry with the user's input
func EntryWithUserInput(score int, concern, grateful, learn, milestone string) Entry {
	return Entry{
		ID:        -1,
		Score:     score,
		Concern:   concern,
		Grateful:  grateful,
		Learn:     learn,
		Milestone: milestone,
		Entered:   -1,
	}
}

// EntryWithAllFields : makes a new entry, setting all fields
func EntryWithAllFields(id, score int, concern, grateful, learn, milestone string, entered int64) Entry {
	return Entry{
		ID:        id,
		Score:     score,
		Concern:   concern,
		Grateful:  grateful,
		Learn:     learn,
		Milestone: milestone,
		Entered:   entered,
	}
}

// PrintEntry : Prints the Entry struct
func PrintEntry(entry Entry) {
	fmt.Printf("%+v\n", entry)
}
