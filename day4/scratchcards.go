package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

// day 4 part 1 sets a task of parsing a set of lines. On one side is a set of winning numbers. On the other a set of "my" numbers
// the first match between "my" numbers and the winning one makes the card worth 1 point and each match after the first doubles the point value of that card.
// i.e. the score is basically two to the power of the number of matching numbers
func main() {
	inputText, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(inputText), "\n")

	score := 0

	for i := 0; i < len(lines); i++ {
		score += getLineScore(lines[i])
	}

	fmt.Println(score)
}

func getLineScore(line string) int {
	lineScore := 0
	// stripping the beginning of the string
	line = line[strings.Index(line, ":")+1:]
	// split the left and right hand side (sepaarted by |).
	parts := strings.Split(string(line), "|")
	re := regexp.MustCompile("\\d+")
	// 1st part is the winning numbers. 2nd part is "my" numbers. Find all numerical strings in each part
	winningNumbers := re.FindAllString(parts[0], -1)
	myNumbers := re.FindAllString(parts[1], -1)

	// now find the number of matches - i.e. the intersection between the two. go doesn't have an existing efficient implementation that I know of so use a simple O(n^2) search
	matches := 0
	for i := 0; i < len(myNumbers); i++ {
		if SliceContains(winningNumbers, myNumbers[i]) {
			matches++
		}
	}

	// score for a given "card" is basically 2 ^ the number of matched numbers
	if matches > 0 {
		lineScore += int(math.Pow(2, float64(matches-1)))

	}

	return lineScore
}

// SliceContains returns true if s contains v.
func SliceContains(stringValues []string, value string) bool {
	for _, strVal := range stringValues {
		if strVal == value {
			return true
		}
	}

	return false
}
