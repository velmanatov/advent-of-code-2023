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
// Part 2 changes this up so that each line you count the "winning numbers" - n. You then get an extra copy of the n next cards. This happens iteratively for all
// original cards and card won along the way. Count up the total number of cards at the end.
func main() {
	inputText, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(inputText), "\n")

	scorePart1 := 0
	cardNumWinningNumbersMap := make(map[int]int)

	for i := 0; i < len(lines); i++ {
		cardScore, numWinning := getLineScore(lines[i])
		scorePart1 += cardScore

		numCopiesOfThisCard := cardNumWinningNumbersMap[i] + 1

		// keep track of number of extra copies of following cards we have won (puzzle rules suggest we should never get index out of bounds at the end)
		for j := 1; j <= numWinning; j++ {
			cardNumWinningNumbersMap[i + j] = cardNumWinningNumbersMap[i + j] + numCopiesOfThisCard
		}
	}

	// we know we have at least one copy of all cards (so min score is number of lines in the file)
	scorePart2 := len(lines)
	// now add up all of extra copies we have won to get the answer for part 2 of the puzzle
	for i := 0; i < len(lines); i++ {
		scorePart2 += cardNumWinningNumbersMap[i];
	}

	fmt.Println(scorePart1)
	fmt.Println(scorePart2)
}

// Gets the line score for part 1 of the puzzle and the number of macthes for part 2
func getLineScore(line string) (int, int) {
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

	return lineScore, matches;
}

// SliceContains returns true if stringValues contains value.
func SliceContains(stringValues []string, value string) bool {
	for _, strVal := range stringValues {
		if strVal == value {
			return true
		}
	}

	return false
}
