// Package input provides functions related to user input.
package input

import (
	"fmt"
	"strings"
)

func isQuestion(msg string) bool {
	for {
		fmt.Printf(msg + " [y/n]: ")

		var retry string
		_, err := fmt.Scanln(&retry)
		if err != nil {
			continue
		}

		if len(retry) > 0 {
			return strings.ToLower(string(retry[0])) == "y"
		}
	}
}
