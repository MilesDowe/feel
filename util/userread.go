package util

import (
	"bufio"
	"os"
	"strings"
)

// GetUserConfirmation : Get a yes/no answer from the user
func GetUserConfirmation() bool {
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')

	response = strings.TrimSpace(response)
	response = strings.ToLower(response)

	if response == "y" || response == "yes" {
		return true
	}
	return false
}
