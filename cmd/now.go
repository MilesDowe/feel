package cmd

import (
	"bufio"
	"github.com/google/go-cmp/cmp"
	"github.com/milesdowe/feel/entity"
	"github.com/milesdowe/feel/util"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// Cobra command creation details
var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Save happiness score",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		// create object for prompting/guiding user
		prompter := util.PromptPrinter{
			Reader: reader,
			Min:    Min,
			Max:    Max,
		}

		// get most recent record, if today, prompt to overwrite
		if entry := checkForExistingEntry(); entry.ID != -1 {
			if prompter.OverwriteEntry(entry) {
				util.DeleteRecord(entry.ID)
			} else {
				return
			}
		}

		entry := readUserInput(prompter)
		recordToDb(entry)
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)
}

// `now` command

type date struct {
	Year  int
	Month time.Month
	Day   int
}

const addRecord = `
INSERT INTO feel_recording (score, concern, grateful, learn, milestone, entered)
VALUES (?, ?, ?, ?, ?, ?)`

const getRecentRecord = `
SELECT id, score, concern, grateful, learn, milestone, entered
FROM feel_recording
WHERE id = (SELECT max(id) FROM feel_recording)`

// Min : Lowest a happy score can be
const Min = 1

// Max : Highest a happy score can be
const Max = 10

func getDateFromUnixTime(unixTime int64) date {
	t := time.Unix(unixTime, 0)
	return date{t.Year(), t.Month(), t.Day()}
}

func getDateNow() date {
	t := time.Now()
	return date{t.Year(), t.Month(), t.Day()}
}

// prompts user for happiness details, returns results
func readUserInput(prompter util.PromptPrinter) entity.Entry {
	score := prompter.GetScore()

	concern := prompter.GetOptionalDetail("Anything have you concerned?")
	grateful := prompter.GetOptionalDetail("Do you feel grateful for anything?")
	learn := prompter.GetOptionalDetail("Did you learn anything new today?")
	milestone := prompter.GetOptionalDetail("Any noteable milestones?")

	// default id and entry date to -1, will be provided upon insert
	return entity.EntryWithUserInput(
		score,
		concern,
		grateful,
		learn,
		milestone,
	)
}

// saves happiness details to the database
func recordToDb(entry entity.Entry) {
	db := util.OpenDb()

	stmt, _ := db.Prepare(addRecord)
	defer stmt.Close()
	stmt.Exec(entry.Score, entry.Concern, entry.Grateful, entry.Learn, entry.Milestone, time.Now().Unix())
}

func checkForExistingEntry() entity.Entry {
	// get the latest record
	db := util.OpenDb()
	rows, _ := db.Query(getRecentRecord)

	defer rows.Close()

	var id, score int
	var concern, grateful, learn, milestone string
	var entered int64

	for rows.Next() {
		rows.Scan(&id, &score, &concern, &grateful, &learn, &milestone, &entered)
	}

	// if it was entered today, provide the details
	recordTime := getDateFromUnixTime(entered)
	nowTime := getDateNow()
	if cmp.Equal(nowTime, recordTime) {
		return entity.EntryWithAllFields(id, score, concern, grateful, learn, milestone, entered)
	}
	return entity.EmptyEntry()
}
