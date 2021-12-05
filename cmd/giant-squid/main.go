package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wouterh/aoc2021/internal/input"
)

func readDraws(input string) ([]int, error) {
	parts := strings.Split(input, ",")
	results := make([]int, len(parts))
	for i, v := range parts {
		intv, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		results[i] = intv
	}
	return results, nil
}

func readBoards(input []string) ([]*Board, error) {
	results := []*Board{}
	for i := 0; i < len(input)-4; i = i + 6 {
		board, err := NewBoard(input[i : i+5])
		if err != nil {
			return nil, err
		}
		results = append(results, board)
	}
	return results, nil
}

func readInputs(lines []string) ([]int, []*Board, error) {
	draws, err := readDraws(lines[0])
	if err != nil {
		return nil, nil, err
	}
	boards, err := readBoards(lines[2:])
	if err != nil {
		return nil, nil, err
	}
	return draws, boards, nil
}

func findWinnersScore(lines []string) (int, error) {
	draws, boards, err := readInputs(lines)
	if err != nil {
		return 0, err
	}
	for _, draw := range draws {
		for _, board := range boards {
			board.Draw(draw)
			if board.Wins() {
				score := board.SumUnmarked() * draw
				return score, nil
			}
		}
	}
	return 0, errors.New("no winner found")
}

func findLosersScore(lines []string) (int, error) {
	draws, boards, err := readInputs(lines)
	if err != nil {
		return 0, err
	}
	won := make([]bool, len(boards))
	nbWon := 0
	for _, draw := range draws {
		for i, board := range boards {
			if won[i] {
				continue
			}
			board.Draw(draw)
			if board.Wins() {
				nbWon++
				won[i] = true
				if nbWon == len(boards) {
					score := board.SumUnmarked() * draw
					return score, nil
				}
			}
		}
	}
	return 0, errors.New("no loser found")
}

func main() {
	lines, err := input.ReadStrings(os.Args[1])
	if err != nil {
		panic("Could not read input")
	}
	score, err := findWinnersScore(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println(score)
	losersScore, err := findLosersScore(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println(losersScore)
}
