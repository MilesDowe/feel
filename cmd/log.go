package cmd

import (
	"fmt"
	"github.com/milesdowe/feel/util"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
)

// Cobra command creation details
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show log of entries",
	Run: func(cmd *cobra.Command, args []string) {
		printLog()
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}

// `log` command

const getAllRecords = `
SELECT id, score, concern, grateful, learn, entered
FROM feel_recording`

// PrintLog : outputs database records
func printLog() {
	db := util.OpenDb()
	rows, _ := db.Query(getAllRecords)
	defer rows.Close()

	var id int
	var score int
	var concern string
	var grateful string
	var learn string
	var entered int64

	fmt.Println()
	for rows.Next() {
		rows.Scan(&id, &score, &concern, &grateful, &learn, &entered)

		fmt.Println("Date: " + time.Unix(entered, 0).String())
		fmt.Println("Score: " + strconv.Itoa(score))
		fmt.Println("Concerned: " + format(concern))
		fmt.Println("Grateful: " + format(grateful))
		fmt.Println("Learned: " + format(learn))
		fmt.Println()
	}
}

func format(in string) string {
	result := strings.Trim(in, " \n")

	if result == "" {
		return "<skipped>"
	}
	return result
}
