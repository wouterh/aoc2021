package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/wouterh/aoc2021/internal/input"
	"github.com/wouterh/aoc2021/internal/util"
)

var checkPoints map[rune]int = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completionPoints map[rune]int = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

// FindCompletion finds the (reversed) completion or returns the first invalid character
func findCompletion(line string) ([]rune, rune) {
	var expected []rune
	for _, r := range line {
		switch r {
		case '(':
			expected = append(expected, ')')
		case '[':
			expected = append(expected, ']')
		case '{':
			expected = append(expected, '}')
		case '<':
			expected = append(expected, '>')
		default:
			// other char -> check if it is expected
			if r == expected[len(expected)-1] {
				// OK! pop from the end expected
				expected = expected[:len(expected)-1]
			} else {
				// NOK!
				return nil, r
			}
		}
	}
	return expected, 0
}

func findWrongClosing(line string) rune {
	_, r := findCompletion(line)
	return r
}

func findCheckScores(lines []string) (scores []int) {
	for _, line := range lines {
		r := findWrongClosing(line)
		if r != 0 {
			scores = append(scores, checkPoints[r])
		}
	}
	return scores
}

func findCompletionScores(lines []string) (scores []int) {
	for _, line := range lines {
		reversedCompletion, _ := findCompletion(line)
		if reversedCompletion != nil {
			score := 0
			for i := len(reversedCompletion) - 1; i >= 0; i-- {
				score = score * 5
				score += completionPoints[reversedCompletion[i]]
			}
			scores = append(scores, score)
		}
	}
	return scores
}

func main() {
	lines, err := input.ReadStrings(os.Args[1])
	if err != nil {
		panic(err)
	}
	scores := findCheckScores(lines)
	fmt.Println(util.Sum(scores))
	scores = findCompletionScores(lines)
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
