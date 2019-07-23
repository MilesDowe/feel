package cmd

import (
	"fmt"
	"github.com/milesdowe/feel/entity"
	"github.com/milesdowe/feel/util"
	"github.com/montanaflynn/stats"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

//
// Constants and structs
//
type entrySet []entity.Entry

type collectedData struct {
	scores       []float64
	concernCnt   int
	gratefulCnt  int
	learnCnt     int
	milestoneCnt int
}

//
// Cobra command creation details
//
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "View and export stats from your history of entries",
	Run: func(cmd *cobra.Command, args []string) {
		// get database records, depending on user input (all or date range)
		var entries entrySet
		if ago > 0 {
			entries = populateEntries(agoQuery(ago))
		} else {
			entries = populateEntries(rangeQuery(begin, end))
		}

		// if export option provided, print to stdout
		if export != "" {
			switch export {
			case "csv":
				printCsv(entries)
			default:
				fmt.Printf("Format %v unrecognized\n", export)
			}

		} else {
			// else, aggregate summary data and print
			data := getRelevantEntryData(entries)
			printStats(data)
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

//
// `stat` command
//
func printStats(data collectedData) {
	mean, _ := stats.Mean(data.scores)
	stddev, _ := stats.StdDevS(data.scores)

	printHeader("Your happiness at a glance")
	fmt.Printf("Mean    : %.2f\n", mean)
	fmt.Printf("Std.Dev.: %.2f\n\n", stddev)

	printHeader("Details summary")
	printProvided("Concerned", data.concernCnt, data.scores)
	printProvided("Grateful", data.gratefulCnt, data.scores)
	printProvided("Learned", data.learnCnt, data.scores)
	printProvided("Milestones", data.milestoneCnt, data.scores)
}

func printHeader(title string) {
	fmt.Printf("## %v\n\n", title)
}

func printProvided(category string, count int, scores []float64) {
	fmt.Printf("%v provided: %v (%.1f%%)\n", category, count, percent(count, len(scores)))
}

func percent(numer, denom int) float64 {
	return float64(numer) / float64(denom) * 100
}

func getRelevantEntryData(entries entrySet) collectedData {

	scores := make([]float64, len(entries))

	concernCnt := 0
	gratefulCnt := 0
	learnCnt := 0
	milestoneCnt := 0

	for i := 0; i < len(entries); i++ {
		// get an slice of scores for central tendency
		scores[i] = float64(entries[i].Score)

		// get counts of times extra details were provided
		concernCnt = incrementCount(entries[i].Concern, concernCnt)
		gratefulCnt = incrementCount(entries[i].Grateful, gratefulCnt)
		learnCnt = incrementCount(entries[i].Learn, learnCnt)
		milestoneCnt = incrementCount(entries[i].Milestone, milestoneCnt)

		// TODO: Do something with the text. Something simple, like most reported words, or complex,
		//       like categorizing inputs (e.g., seeing how many "concerns" are work-related)
	}

	return collectedData{scores, concernCnt, gratefulCnt, learnCnt, milestoneCnt}
}

func incrementCount(value string, cnt int) int {
	if strings.TrimSpace(value) != "" {
		cnt++
	}
	return cnt
}

// populateEntries : adds entries from database to the provided array.
func populateEntries(query string) []entity.Entry {
	result := []entity.Entry{}

	db := util.OpenDb()
	rows, _ := db.Query(query)

	defer rows.Close()

	var (
		id, score                           int
		concern, grateful, learn, milestone string
		entered                             int64
	)

	for rows.Next() {
		rows.Scan(&id, &score, &concern, &grateful, &learn, &milestone, &entered)
		result = append(result, entity.EntryWithAllFields(id, score, concern, grateful, learn, milestone, entered))
	}
	return result
}

// printCsv : for use in `export` flag, prints to stdout
func printCsv(entries entrySet) {
	fmt.Printf("\"id\",\"score\",\"concerned\",\"grateful\",\"learned\",\"milestone\",\"entered\"\n")
	for i := 0; i < len(entries); i++ {
		fmt.Printf(
			"\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\n",
			strconv.Itoa(entries[i].ID),
			strconv.Itoa(entries[i].Score),
			strings.ReplaceAll(entries[i].Concern, "\"", "\\\""),   // escape quotes since we
			strings.ReplaceAll(entries[i].Grateful, "\"", "\\\""),  // used them to encapsulate
			strings.ReplaceAll(entries[i].Learn, "\"", "\\\""),     // the text
			strings.ReplaceAll(entries[i].Milestone, "\"", "\\\""), // the text
			strconv.FormatInt(entries[i].Entered, 10))
	}
}

// agoQuery : Expands select-all query to limit the number of days ago.
func agoQuery(days int) string {
	daysStr := strconv.Itoa(days)
	return util.SelectAllQuery + "WHERE date(entered, 'unixepoch') >= date('now', 'start of day', '-" +
		daysStr + " day')"
}

// rangeQuery : Expands select-all query to include a date range.
// TODO: fix localization issue
func rangeQuery(begin, end string) string {
	result := util.SelectAllQuery

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
