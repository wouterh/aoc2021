package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/wouterh/aoc2021/internal/input"
)

func readLines(inputLines []string) ([]*Line, error) {
	lines := make([]*Line, len(inputLines))
	for i, l := range inputLines {
		line, err := NewLine(l)
		if err != nil {
			return nil, errors.New("couldn't parse input lines")
		}
		lines[i] = line
	}
	return lines, nil
}

func CountOverlapsHV(board *Board, lines []*Line) int {
	for _, line := range lines {
		if !line.IsDiagonal() {
			board.Draw(line)
		}
	}
	return board.CountOverlaps(2)
}

func AddDiagonalsAndCountOverlaps(board *Board, lines []*Line) int {
	for _, line := range lines {
		if line.IsDiagonal() {
			board.Draw(line)
		}
	}
	return board.CountOverlaps(2)
}

func main() {
	inputLines, err := input.ReadStrings(os.Args[1])
	if err != nil {
		panic("Couldn't read input")
	}
	lines, err := readLines(inputLines)
	if err != nil {
		panic(err)
	}
	board := NewBoard(1000)
	overlaps := CountOverlapsHV(board, lines)
	fmt.Println(overlaps)
	overlaps = AddDiagonalsAndCountOverlaps(board, lines)
	fmt.Println(overlaps)
}
