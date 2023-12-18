package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	filePath := "file.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		first, last := getFirstAndLastDigits(line)
		combined, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
		sum += combined
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading file:", err)
	}

	fmt.Println("Total Sum:", sum)
}

func getFirstAndLastDigits(s string) (int, int) {
	firstDigit := -1
	lastDigit := -1

	wordToDigit := map[string]int{
		"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4,
		"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	for _, word := range strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	}) {
		if val, exists := wordToDigit[word]; exists {
			if firstDigit == -1 {
				firstDigit = val
			}
			lastDigit = val
		} else if len(word) > 0 {
			for _, r := range word {
				if unicode.IsDigit(r) {
					digit, _ := strconv.Atoi(string(r))
					if firstDigit == -1 {
						firstDigit = digit
					}
					lastDigit = digit
				}
			}
		}
	}

	if firstDigit == -1 {
		firstDigit = 0
	}
	if lastDigit == -1 {
		lastDigit = firstDigit
	}

	return firstDigit, lastDigit
}
