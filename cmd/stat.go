package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Cobra command creation details
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "View and export stats from your history of entries",
	Run: func(cmd *cobra.Command, args []string) {

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

func init() {
	statCmd.Flags().StringVarP(&export, "export", "x", "", "Output stats to a file. Available formats are: csv")

	rootCmd.AddCommand(statCmd)
}

// `stat` command
