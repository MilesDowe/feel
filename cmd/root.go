package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

var rootCmd = &cobra.Command{
	Use:   "feel",
	Short: "A command line happiness tracker",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute : runs the program
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
