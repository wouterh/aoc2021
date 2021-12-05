package main

import (
	"strconv"
	"strings"
)

type Square struct {
	value  int
	marked bool
}

type Board struct {
	squares [][]Square
}

func NewBoard(lines []string) (*Board, error) {
	board := &Board{}
	board.squares = make([][]Square, 5)
	for i, line := range lines {
		board.squares[i] = []Square{}
		parts := strings.Split(line, " ")
		for _, part := range parts {
			if part != "" {
				v, err := strconv.Atoi(part)
				if err != nil {
					return nil, err
				}
				board.squares[i] = append(board.squares[i], Square{value: v})
			}
		}
	}
	return board, nil
}

func (b *Board) Draw(drawn int) {
	for i, row := range b.squares {
		for j, square := range row {
			if square.value == drawn {
				b.squares[i][j].marked = true
				return
			}
		}
	}
}

func (b *Board) Wins() bool {
	for i := 0; i < 5; i++ {
		nbDrawnHorizontal := 0
		nbDrawnVertical := 0
		for j := 0; j < 5; j++ {
			if b.squares[i][j].marked {
				nbDrawnHorizontal++
			}
			if b.squares[j][i].marked {
				nbDrawnVertical++
			}
		}
		if nbDrawnHorizontal == 5 || nbDrawnVertical == 5 {
			return true
		}
	}
	return false
}

func (b *Board) SumUnmarked() int {
	var sum int = 0
	for _, row := range b.squares {
		for _, square := range row {
			if !square.marked {
				sum = sum + square.value
			}
		}
	}
	return sum
}
