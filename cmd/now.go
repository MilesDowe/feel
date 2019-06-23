package cmd

import (
	"bufio"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/milesdowe/feel/util"
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

type entry struct {
	ID       int
	Score    string
	Concern  string
	Grateful string
	Learn    string
	Entered  int64
}

type date struct {
	Year  int
	Month time.Month
	Day   int
}

const addRecord = "INSERT INTO feel_recording (score, concern, grateful, learn, entered) VALUES (?, ?, ?, ?, ?)"
const getRecentRecord = "SELECT * FROM feel_recording where id = (select max(id) from feel_recording)"

// Min : Lowest a happy score can be
const Min = 1

// Max : Highest a happy score can be
const Max = 10

// MinStr : String representation of Min
var MinStr = strconv.Itoa(Min)

// MaxStr : String representation of Max
var MaxStr = strconv.Itoa(Max)

func getDateFromTime(unixTime int64) date {
	t := time.Unix(unixTime, 0)
	return date{t.Year(), t.Month(), t.Day()}
}

func getDateNow() date {
	t := time.Now()
	return date{t.Year(), t.Month(), t.Day()}
}

// prompts user for happiness details, returns results
func readUserInput() entry {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("How happy do you feel right now? Choose from %s (awful) to %s (great):\n", MaxStr, MinStr)
	score, _ := reader.ReadString('\n')

	fmt.Printf("Anything have you concerned? (<enter> to skip)\n")
	concern, _ := reader.ReadString('\n')

	fmt.Printf("Do you feel grateful for anything? (<enter> to skip)\n")
	grateful, _ := reader.ReadString('\n')

	fmt.Printf("Did you learn anything new today? (<enter> to skip)\n")
	learn, _ := reader.ReadString('\n')

	// default id and entry date to -1, will be provided upon insert
	return entry{-1, score, concern, grateful, learn, -1}
}

// saves happiness details to the database
func recordToDb(entry entry) {
	db := util.OpenDb()

	entry.Score = checkScoreInput(entry.Score)

	stmt, _ := db.Prepare(addRecord)
	defer stmt.Close()
	stmt.Exec(entry.Score, entry.Concern, entry.Grateful, entry.Learn, time.Now().Unix())
}

func checkForExistingEntry() entry {
	// get the latest record
	db := util.OpenDb()
	rows, _ := db.Query(getRecentRecord)

	defer rows.Close()

	var id, score int
	var concern, grateful, learn string
	var entered int64

	for rows.Next() {
		rows.Scan(&id, &score, &concern, &grateful, &learn, &entered)
	}

	// if it was entered today, provide the details
	recordTime := getDateFromTime(entered)
	nowTime := getDateNow()
	if cmp.Equal(nowTime, recordTime) {
		return entry{id, strconv.Itoa(score), concern, grateful, learn, entered}
	}
	return entry{-1, "", "", "", "", -1}
}

func overwriteEntry(entry entry) bool {
	fmt.Printf("An entry for today already exists:\n")
	fmt.Printf("---------------------------------\n")
	fmt.Printf("Score: %s\n", entry.Score)
	fmt.Printf("Concern:\n%s", entry.Concern)
	fmt.Printf("Grateful:\n%s", entry.Grateful)
	fmt.Printf("Learned:\n%s", entry.Learn)
	fmt.Printf("---------------------------------\n")
	fmt.Printf("Delete it and enter a new one? [Y/n]: ")

	return util.GetUserConfirmation()
}

// keep user's input score number within expected range
func checkScoreInput(score string) string {
	score = strings.TrimSpace(score)
	i, err := strconv.Atoi(score)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if i < Min {
		score = MinStr
	} else if i > Max {
		score = MaxStr
	}
	return score
}