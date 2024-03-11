package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValidJSON(input string) bool {
	// Check if the trimmed input is exactly "{}"
	return strings.TrimSpace(input) == "{}"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	inputProvided := false

	for scanner.Scan() {
		inputProvided = true // Flag set to true if there's input
		input := scanner.Text()

		if isValidJSON(input) {
			fmt.Println("Valid JSON")
			os.Exit(0) // Exit with 0 for valid JSON
		} else {
			fmt.Println("Invalid JSON")
			os.Exit(1) // Exit with 1 for invalid JSON
		}
	}

	// Check if input was provided after the loop ends
	if !inputProvided {
		fmt.Println("Invalid JSON") // No input provided
		os.Exit(1)                  // Exit with 1 for invalid JSON due to no input
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1) // Exit with 1 if there is an error reading input
	}
}
