package entity

// Entry : a user's `feel` entry
type Entry struct {
	ID       int
	Score    string
	Concern  string
	Grateful string
	Learn    string
	Entered  int64
}

// EmptyEntry : makes a new entry with default fields
func EmptyEntry() Entry {
	return Entry{
		ID:       -1,
		Score:    "",
		Concern:  "",
		Grateful: "",
		Learn:    "",
		Entered:  -1,
	}
}

// EntryWithUserInput : makes a new entry with the user's input
func EntryWithUserInput(score, concern, grateful, learn string) Entry {
	return Entry{
		ID:       -1,
		Score:    score,
		Concern:  concern,
		Grateful: grateful,
		Learn:    learn,
		Entered:  -1,
	}
}

// EntryWithAllFields : makes a new entry, setting all fields
func EntryWithAllFields(id int, score, concern, grateful, learn string, entered int64) Entry {
	return Entry{
		ID:       id,
		Score:    score,
		Concern:  concern,
		Grateful: grateful,
		Learn:    learn,
		Entered:  entered,
	}
}
