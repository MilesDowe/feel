package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

//
// Cobra command creation details
//
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "See a chronological list of a single entry item (e.g., \"Concerned\")",
	Long: `
    See a chronological list of a single entry item (e.g., \"Concerned\").

    The 'list' command doesn't necessarily provide anything that cannot be retrieved
    from the 'stat' command's export function. But it provides the user the ability
    to specifically review a focused, sequential list of things they were concerned
    about, for example.
    `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("word")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

//
// `list` command
//
type itemList struct {
    datetime int64
    itemValue string
}

func selectItemQuery(item string) string {
    mainQuery := "SELECT entered, %s from feel_recording"

    switch item {
    case "c":
        return strings.ReplaceAll(mainQuery, "%s", "concern")
    case "g":
        return strings.ReplaceAll(mainQuery, "%s", "grateful")
    case "l":
        return strings.ReplaceAll(mainQuery, "%s", "learn")
    case "m":
        return strings.ReplaceAll(mainQuery, "%s", "milestone")
    default:
        return ""
    }
}
