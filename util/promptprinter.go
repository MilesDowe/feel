package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"github.com/milesdowe/feel/entity"
)

/* A utility struct for guiding and reading user responses, specifically for the `now` command.
 * This is because the `now` command is currently the most interactive command so far. */

// PromptPrinter : a helper struct to read user input and convert to expect values
type PromptPrinter struct {
    Reader *bufio.Reader
    Min int
    Max int
}

// OverwriteEntry : ask user to rewrite values if entry already exists
func (pp *PromptPrinter) OverwriteEntry(entry entity.Entry) bool {
	fmt.Printf("----------------------------------\n")
	fmt.Printf("An entry for today already exists:\n")
	fmt.Printf("----------------------------------\n")
	fmt.Printf("Score:\n> %v\n\n", entry.Score)
	fmt.Printf("Concern:\n> %v\n\n", entry.Concern)
	fmt.Printf("Grateful:\n> %v\n\n", entry.Grateful)
	fmt.Printf("Learned:\n> %v\n\n", entry.Learn)
	fmt.Printf("Milestone:\n> %v\n", entry.Milestone)
	fmt.Printf("---------------------------------\n")
	fmt.Printf("Delete it and enter a new one? [Y/n]: ")

	return pp.getUserConfirmation()
}

// GetOptionalDetail : a method to prompt for skippable string inputs
func (pp *PromptPrinter) GetOptionalDetail(prompt string) string {
	skipNotice := " (<enter> to skip)\n> "

	fmt.Printf(prompt + skipNotice)
	response, _ := pp.Reader.ReadString('\n')
	fmt.Printf("\n")
	return strings.TrimSpace(response)
}

// GetScore : a method to prompt for the required happiness score
func (pp *PromptPrinter) GetScore() int {
    min := strconv.Itoa(pp.Min)
    max := strconv.Itoa(pp.Max)

	fmt.Printf("How happy do you feel right now? Choose from %s (awful) to %s (great):\n> ", min, max)
	score, _ := pp.Reader.ReadString('\n')
	fmt.Printf("\n")

	// check provided score is in range
	scoreNum := pp.checkScoreInput(score)

	return scoreNum
}

// get a yes/no answer from the user
func (pp *PromptPrinter) getUserConfirmation() bool {
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')

	response = strings.TrimSpace(response)
	response = strings.ToLower(response)

	if response == "y" || response == "yes" {
		return true
	}
	return false
}

// keep user's input score number within expected range
func (pp *PromptPrinter) checkScoreInput(score string) int {
	score = strings.TrimSpace(score)
	result, err := strconv.Atoi(score)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if result < pp.Min {
		result = pp.Min
	} else if result > pp.Max {
		result = pp.Max
	}
	return result
}
