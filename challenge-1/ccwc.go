package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

const FILE_STDIN = "stdin"

const OPT_DEFAULT = "default"

const ARG_COUNT_CHAR = "-c"
const ARG_COUNT_LINE = "-l"
const ARG_COUNT_WORD = "-w"
const ARG_COUNT_RUNE = "-m"

var ALLOWED_ARGS = map[string]bool{
	ARG_COUNT_CHAR: true,
	ARG_COUNT_LINE: true,
	ARG_COUNT_WORD: true,
	ARG_COUNT_RUNE: true,
}

func main() {
	var err error
	var input string
	filename := FILE_STDIN
	option := OPT_DEFAULT

	args := os.Args
	argsLen := len(args)
	if argsLen > 3 {
		printHelpAndExit()
	}

	if argsLen == 3 {
		if _, ok := ALLOWED_ARGS[args[1]]; !ok {
			printHelpAndExit()
		}

		option = args[1]
		filename = args[2]
	}

	if argsLen == 2 {
		if _, ok := ALLOWED_ARGS[args[1]]; ok {
			option = args[1]
		} else {
			filename = args[1]
		}
	}

	if filename == FILE_STDIN {
		input, err = readFromStdin()
	} else {
		input, err = readFile(filename)
	}
	handleError(err)

	// Perform the requested operation
	switch option {
	case ARG_COUNT_CHAR:
		fmt.Printf("  %d %s\n", len(input), filename)
	case ARG_COUNT_LINE:
		fmt.Printf("  %d %s\n", countLines(input), filename)
	case ARG_COUNT_WORD:
		fmt.Printf("  %d %s\n", countWords(input), filename)
	case ARG_COUNT_RUNE:
		fmt.Printf("  %d %s\n", utf8.RuneCountInString(input), filename)
	default:
		fmt.Printf("  %d %d %d %s\n", countLines(input), countWords(input), len(input), filename)
	}
}

func printHelpAndExit() {
	fmt.Println(fmt.Printf("Usage: ccwc [%s|%s|%s|%s] [filename]", ARG_COUNT_CHAR, ARG_COUNT_LINE, ARG_COUNT_WORD, ARG_COUNT_RUNE))
	os.Exit(1)
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
