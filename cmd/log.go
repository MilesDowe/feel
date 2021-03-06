package cmd

import (
	"fmt"
	"github.com/MilesDowe/feel/util"
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

	var (
		id, score                int
		concern, grateful, learn string
		entered                  int64
	)

	fmt.Println()
	for rows.Next() {
		rows.Scan(&id, &score, &concern, &grateful, &learn, &entered)

		fmt.Printf("Date: %v\n", time.Unix(entered, 0).String())
		fmt.Printf("Score: %v\n", strconv.Itoa(score))
		if given(concern) {
			fmt.Printf("Concerned:\n> %v\n", concern)
		}
		if given(grateful) {
			fmt.Printf("Grateful:\n> %v\n", grateful)
		}
		if given(learn) {
			fmt.Printf("Learned:\n> %v\n", learn)
		}
		fmt.Println()
	}
}

func given(in string) bool {
	result := strings.TrimSpace(in)
	return result != ""
}
