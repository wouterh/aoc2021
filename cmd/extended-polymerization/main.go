package main

import (
	"fmt"
	"os"

	"github.com/wouterh/aoc2021/internal/input"
)

func partOne(lines []string) int {
	p := NewPolymerization(lines)
	for i := 0; i < 10; i++ {
		p.Step()
	}
	min, max := p.MinimaxCounts()
	return max - min
}

func partTwo(lines []string) int {
	p := NewPolymerization(lines)
	for i := 0; i < 40; i++ {
		p.Step()
	}
	min, max := p.MinimaxCounts()
	return max - min
}

func main() {
	lines, err := input.ReadStrings(os.Args[1])
	if err != nil {
		panic(err)
	}
	answerOne := partOne(lines)
	fmt.Println("Answer 1:", answerOne)

	answerTwo := partTwo(lines)
	fmt.Println("Answer 2:", answerTwo)
}
