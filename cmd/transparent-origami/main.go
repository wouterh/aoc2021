package main

import (
	"fmt"
	"os"

	"github.com/wouterh/aoc2021/internal/input"
)

func main() {
	lines, err := input.ReadStrings(os.Args[1])
	if err != nil {
		panic(err)
	}
	paper, nbDots := NewPaper(lines)
	folds := lines[nbDots+1:]

	for i, fold := range folds {
		paper.Fold(fold)
		if i == 0 {
			fmt.Println("Answer part 1:", paper.NumberOfDots())
		}
	}
	// For part two we will have to look at the resulting paper
	paper.Print()
}
