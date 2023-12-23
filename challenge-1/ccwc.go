package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var err error
	var input string
	var filename string
	var option string

	// Determine the option and filename
	if len(os.Args) == 1 {
		input, err = readFromStdin()
		handleError(err)
		filename = "stdin"
		option = "default"
	} else if len(os.Args) == 2 {
		arg := os.Args[1]
		if arg == "-c" || arg == "-l" || arg == "-w" || arg == "-m" {
			input, err = readFromStdin()
			handleError(err)
			filename = "stdin"
			option = arg
		} else {
			input, err = readFile(arg)
			handleError(err)
			filename = arg
			option = "default"
		}
	} else if len(os.Args) == 3 {
		option = os.Args[1]
		filename = os.Args[2]
		input, err = readFile(os.Args[2])
		handleError(err)
	} else {
		fmt.Println("Usage: ccwc [-c|-l|-w|-m] [filename]")
		os.Exit(1)
	}

	// Perform the requested operation
	switch option {
	case "-c":
		fmt.Printf("  %d %s\n", len(input), filename)
	case "-l":
		fmt.Printf("  %d %s\n", countLines(input), filename)
	case "-w":
		fmt.Printf("  %d %s\n", countWords(input), filename)
	case "-m":
		fmt.Printf("  %d %s\n", utf8.RuneCountInString(input), filename)
	default:
		fmt.Printf("  %d %d %d %s\n", countLines(input), countWords(input), len(input), filename)
	}
}

func readFromStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	var output strings.Builder
	for {
		input, err := reader.ReadString('\n')
		output.WriteString(input)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return output.String(), nil
}

func readFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func countLines(input string) int {
	return strings.Count(input, "\n")
}

func countWords(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	return wordCount
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
