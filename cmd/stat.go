package cmd

import (
	"fmt"
	"github.com/milesdowe/feel/entity"
	"github.com/milesdowe/feel/util"
	"github.com/spf13/cobra"
	"math"
)

// Cobra command creation details
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "View and export stats from your history of entries",
	Run: func(cmd *cobra.Command, args []string) {
		// get the data
		var entries []entity.Entry
		if ago > 0 {
			entries = populateEntries(agoQuery)
		} else {
			entries = populateEntries(rangeQuery(begin, end))
		}

		printStats(entries)

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

// Flags
var (
	export string
	ago    int
	begin  string
	end    string
)

func init() {
	statCmd.Flags().StringVarP(&export, "export", "x", "", "Output stats to a file. Available formats are: csv")

	statCmd.Flags().IntVarP(&ago, "ago", "a", 0, "Get data for the last number of days provided.")

	statCmd.Flags().StringVarP(&begin, "begin", "b", "", `The date to begin data review, as YYYYMMDD.
Ignored if --ago flag is provided.`)
	statCmd.Flags().StringVarP(&end, "end", "e", "", `The date to stop data review, as YYYYMMDD.
Ignored if --ago flag is provided.`)

	rootCmd.AddCommand(statCmd)
}

// `stat` command

const (
	allQuery = `SELECT * from feel_recording `
	agoQuery = allQuery + `WHERE entered`
)

func printStats(entries []entity.Entry) {
	scores := make([]float64, len(entries))
	for i := 0; i < len(entries); i++ {
		scores[i] = float64(entries[i].Score)
	}

	fmt.Printf("Mean: %v\n", mean(scores))
	fmt.Printf("Std.Dev.: %v\n", stddev(scores))
}

// populateEntries : adds entries from database to the provided array.
func populateEntries(query string) []entity.Entry {
	result := []entity.Entry{}

	db := util.OpenDb()
	rows, _ := db.Query(query)

	defer rows.Close()

	var (
		id, score                int
		concern, grateful, learn string
		entered                  int64
	)

	for rows.Next() {
		rows.Scan(&id, &score, &concern, &grateful, &learn, &entered)
		result = append(result, entity.EntryWithAllFields(id, score, concern, grateful, learn, entered))
	}
	return result
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
			result = result + dbDate + ` <= '` + end + `'`
		}
	}
	return result
}

func mean(data []float64) float64 {
	var sum float64

	for i := 0; i < len(data); i++ {
		sum = sum + data[i]
	}
	return (sum / float64(len(data)))
}

/* Following function computes standard
 * deviation of a data point
 */
func stddev(data []float64) float64 {
	var dataMean, variance, temp float64

	squaredData := make([]float64, len(data))

	// Get the mean
	dataMean = mean(data)

	// Get distance from mean, then square each value
	for i := 0; i < len(data); i++ {
		temp = data[i] - dataMean
		squaredData[i] = temp * temp
	}

	// Get the variance
	variance = mean(squaredData)
	return float64(math.Sqrt(float64(variance)))
}
