package cmd

import (
	"bufio"
	"fmt"
	"github.com/MilesDowe/feel/entity"
	"github.com/MilesDowe/feel/util"
	"github.com/google/go-cmp/cmp"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Cobra command creation details
var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Save happiness score",
	Run: func(cmd *cobra.Command, args []string) {
		// get most recent record, if today, prompt to overwrite
		if entry := checkForExistingEntry(); entry.ID != -1 {
			if overwriteEntry(entry) {
				util.DeleteRecord(entry.ID)
			} else {
				return
			}
		}

		// prompt user for happy score and save it
		entry := readUserInput()
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

// MinStr : String representation of Min
var MinStr = strconv.Itoa(Min)

// MaxStr : String representation of Max
var MaxStr = strconv.Itoa(Max)

func getDateFromUnixTime(unixTime int64) date {
	t := time.Unix(unixTime, 0)
	return date{t.Year(), t.Month(), t.Day()}
}

func getDateNow() date {
	t := time.Now()
	return date{t.Year(), t.Month(), t.Day()}
}

// prompts user for happiness details, returns results
func readUserInput() entity.Entry {
	reader := bufio.NewReader(os.Stdin)

	skipNotice := " (<enter> to skip)\n> "

	fmt.Printf("How happy do you feel right now? Choose from %s (awful) to %s (great):\n> ", MinStr, MaxStr)
	score, _ := reader.ReadString('\n')

	fmt.Printf("Anything have you concerned?" + skipNotice)
	concern, _ := reader.ReadString('\n')

	fmt.Printf("Do you feel grateful for anything?" + skipNotice)
	grateful, _ := reader.ReadString('\n')

	fmt.Printf("Did you learn anything new today?" + skipNotice)
	learn, _ := reader.ReadString('\n')

	fmt.Printf("Any noteable milestones?" + skipNotice)
	milestone, _ := reader.ReadString('\n')

	// check provided score is in range
	scoreNum := checkScoreInput(score)

	// default id and entry date to -1, will be provided upon insert
	return entity.EntryWithUserInput(
		scoreNum,
		strings.TrimSpace(concern),
		strings.TrimSpace(grateful),
		strings.TrimSpace(learn),
		strings.TrimSpace(milestone),
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

func overwriteEntry(entry entity.Entry) bool {
	fmt.Printf("An entry for today already exists:\n")
	fmt.Printf("---------------------------------\n")
	fmt.Printf("Score:\n> %v\n", entry.Score)
	fmt.Printf("Concern:\n> %v\n", entry.Concern)
	fmt.Printf("Grateful:\n> %v\n", entry.Grateful)
	fmt.Printf("Learned:\n> %v\n", entry.Learn)
	fmt.Printf("Milestone:\n> %v\n", entry.Milestone)
	fmt.Printf("---------------------------------\n")
	fmt.Printf("Delete it and enter a new one? [Y/n]: ")

	return util.GetUserConfirmation()
}

// keep user's input score number within expected range
func checkScoreInput(score string) int {
	score = strings.TrimSpace(score)
	result, err := strconv.Atoi(score)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if result < Min {
		result = Min
	} else if result > Max {
		result = Max
	}
	return result
}
