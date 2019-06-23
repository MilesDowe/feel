package cmd

import (
	"github.com/spf13/cobra"
)

// Cobra command creation details
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "View and export stats from your history of entries",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
    var exportType string
    statCmd.Flags().StringVarP(&exportType, "export", "x", "csv", "Output stats to a file")

	rootCmd.AddCommand(statCmd)
}

// `stat` command

