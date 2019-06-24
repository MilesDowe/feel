package cmd

import (
	"fmt"
	"github.com/milesdowe/feel/entity"
	"github.com/milesdowe/feel/util"
	"github.com/spf13/cobra"
	"strconv"
)

// Cobra command creation details
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "View and export stats from your history of entries",
	Run: func(cmd *cobra.Command, args []string) {
		// get the data
		var entries []entity.Entry

		if ago > 0 {
			populateEntries(entries, agoQuery)
		} else {
			populateEntries(entries, rangeQuery(begin, end))
		}

		// if export option provided, instead contruct a file
		if export != "" {
			switch export {
			case "csv":
				// ...
				fmt.Println("Saved to feel.csv")
			}
		} else {
			// print the data
		}
	},
}

var export string

var ago int

var begin string

var end string

func init() {
	statCmd.Flags().StringVarP(&export, "export", "x", "", "Output stats to a file. Available formats are: csv")

	statCmd.Flags().IntVarP(&ago, "ago", "a", 0, "Get data for the last number of days provided.")

	statCmd.Flags().StringVarP(&begin, "begin", "b", "", `The date to begin data review.
Ignored if --ago flag is provided.`)
	statCmd.Flags().StringVarP(&end, "end", "e", "", `The date to stop data review.
Ignored if --ago flag is provided.`)

	rootCmd.AddCommand(statCmd)
}

// `stat` command

const (
	allQuery = `SELECT * from feel_recording `
	agoQuery = allQuery + `WHERE entered`
)

// populateEntries : adds entries from database to the provided array.
func populateEntries(entries []entity.Entry, query string) {
	db := util.OpenDb()
	rows, _ := db.Query(query)

	defer rows.Close()

	var id int
	var score int
	var concern string
	var grateful string
	var learn string
	var entered int64

	for rows.Next() {
		rows.Scan(&id, &score, &concern, &grateful, &learn, &entered)
		entries = append(entries, entity.EntryWithAllFields(id, strconv.Itoa(score), concern, grateful, learn, entered))
	}

	for _, entry := range entries {
		fmt.Println(entry.ID)
	}
}

// rangeQuery : Constructs a sql query for searching a date range. Gets all unless start and stop
// times are given.
func rangeQuery(begin, end string) string {
	result := allQuery

	hasBegin := begin != ""
	hasEnd := end != ""

    const dbDate string = `strftime('%Y%m%d', entered, 'unixepoch', 'start of day')`

	if hasBegin || hasEnd {
		result = result + `WHERE `
		if hasBegin {
			result = result + dbDate + ` >= '` + begin + `' `
			if hasEnd {
				result = result + `AND `
			}
		}
		if hasEnd {
			result = result + dbDate +  ` <= '` + end + `'`
		}
	}
	return result
}

