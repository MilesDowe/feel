package cmd

import (
	"bufio"
	"fmt"
	"github.com/milesdowe/feel/util"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Save happiness score",
	Run: func(cmd *cobra.Command, args []string) {
		util.VerifyDbExists()

		score, concern, grateful, learn := readUserInput()
		recordToDb(score, concern, grateful, learn)
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)
}

const addRecord = "INSERT INTO feel_recording (score, concern, grateful, learn, entered) VALUES (?, ?, ?, ?, ?)"

// Min : Lowest a happy score can be
const Min = 1

// Max : Highest a happy score can be
const Max = 10

// MinStr : String representation of Min
var MinStr = strconv.Itoa(Min)

// MaxStr : String representation of Max
var MaxStr = strconv.Itoa(Max)

// prompts user for happiness details, returns results
func readUserInput() (string, string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Printf("How happy do you feel right now? Choose from %s (awful) to %s (great):", MaxStr, MinStr)
	score, _ := reader.ReadString('\n')

	fmt.Println("Anything have you concerned? (<enter> to skip)")
	concern, _ := reader.ReadString('\n')

	fmt.Println("Do you feel grateful for anything? (<enter> to skip)")
	grateful, _ := reader.ReadString('\n')

	fmt.Println("Did you learn anything new today? (<enter> to skip)")
	learn, _ := reader.ReadString('\n')
	fmt.Println()

	return score, concern, grateful, learn
}

// saves happiness details to the database
func recordToDb(score, concern, grateful, learn string) {
	db := util.OpenDb()

	score = checkScoreInput(score)

	stmt, _ := db.Prepare(addRecord)
	stmt.Exec(score, concern, grateful, learn, time.Now().Unix())
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
