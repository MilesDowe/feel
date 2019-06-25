package cmd

import (
	"fmt"
	"github.com/milesdowe/feel/entity"
	"github.com/milesdowe/feel/util"
	"github.com/spf13/cobra"
	"math"
	"strings"
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

		// if export option provided, contruct a file
		if export != "" {
            // TODO: think about if we should just print the format and have the user pipe it
            //       to a file if they want.
			switch export {
			case "csv":
				// ...
				fmt.Println("Saved to feel.csv")
			default:
				fmt.Printf("Format %v unrecognized\n", export)
			}

		} else {
			// else, print the data
			printStats(entries)
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

	filledConcern := 0
	filledGrateful := 0
	filledLearn := 0

	for i := 0; i < len(entries); i++ {
		// get an slice of scores for central tendency calculation
		scores[i] = float64(entries[i].Score)

		// get counts of times extra details were provided
		//   may be useful to see how commonly a user tends to report
		//   a particular factor
		if strings.TrimSpace(entries[i].Concern) != "" {
			filledConcern++
		}
		if strings.TrimSpace(entries[i].Grateful) != "" {
			filledGrateful++
		}
		if strings.TrimSpace(entries[i].Learn) != "" {
			filledLearn++
		}

		// TODO: Do something with the text. Something simple, like most reported words, or complex,
		//       like categorizing inputs (e.g., seeing how many "concerns" are work-related)
	}

	fmt.Printf("Your happiness at a glance:\n")
	fmt.Printf("---------------------------\n")
	fmt.Printf("Mean: %.2f\n", mean(scores))
	fmt.Printf("Std.Dev.: %.2f\n\n", stddev(scores))

	fmt.Printf("Details provided:\n")
	fmt.Printf("---------------------------\n")
	fmt.Printf("\"Concerned\" provided: %v (%.1f%%)\n",
		filledConcern, percent(filledConcern, len(scores)))
	fmt.Printf("\"Grateful\" provided: %v (%.1f%%)\n",
		filledGrateful, percent(filledGrateful, len(scores)))
	fmt.Printf("\"Learned\" provided: %v (%.1f%%)\n",
		filledLearn, percent(filledLearn, len(scores)))
}

func percent(numer, denom int) float64 {
	return float64(numer) / float64(denom) * 100
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
// TODO: fix localization issue
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
