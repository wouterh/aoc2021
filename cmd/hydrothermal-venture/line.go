package main

import (
	"strconv"
	"strings"
)

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (l Line) IsVertical() bool {
	return l.x1 == l.x2
}

func (l Line) IsHorizontal() bool {
	return l.y1 == l.y2
}

func (l Line) IsDiagonal() bool {
	return !l.IsHorizontal() && !l.IsVertical()
}

func (l Line) IsDiagonalUp() bool {
	return l.IsDiagonal() && l.YOfMinX() < l.YOfMaxX()
}

func (l Line) IsDiagonalDown() bool {
	return l.IsDiagonal() && l.YOfMinX() > l.YOfMaxX()
}

func (l Line) MinX() int {
	if l.x1 < l.x2 {
		return l.x1
	} else {
		return l.x2
	}
}

func (l Line) MaxX() int {
	if l.x1 > l.x2 {
		return l.x1
	} else {
		return l.x2
	}
}

func (l Line) MinY() int {
	if l.y1 < l.y2 {
		return l.y1
	} else {
		return l.y2
	}
}

func (l Line) MaxY() int {
	if l.y1 > l.y2 {
		return l.y1
	} else {
		return l.y2
	}
}

func (l Line) YOfMinX() int {
	if l.x1 < l.x2 {
		return l.y1
	} else {
		return l.y2
	}
}

func (l Line) YOfMaxX() int {
	if l.x1 > l.x2 {
		return l.y1
	} else {
		return l.y2
	}
}

func NewLine(input string) (*Line, error) {
	parts := strings.Split(input, " -> ")
	firstCo := strings.Split(parts[0], ",")
	secondCo := strings.Split(parts[1], ",")
	x1, err := strconv.Atoi(firstCo[0])
	if err != nil {
		return nil, err
	}
	y1, err := strconv.Atoi(firstCo[1])
	if err != nil {
		return nil, err
	}
	x2, err := strconv.Atoi(secondCo[0])
	if err != nil {
		return nil, err
	}
	y2, err := strconv.Atoi(secondCo[1])
	if err != nil {
		return nil, err
	}
	return &Line{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}, nil
}

type Board struct {
	points [][]int
}

func NewBoard(size int) *Board {
	board := &Board{}
	board.points = make([][]int, size)
	for i := 0; i < size; i++ {
		board.points[i] = make([]int, size)
	}
	return board
}

func (b *Board) Draw(l *Line) {
	if l.IsVertical() {
		for y := l.MinY(); y <= l.MaxY(); y++ {
			b.points[l.x1][y]++
		}
	} else if l.IsHorizontal() {
		for x := l.MinX(); x <= l.MaxX(); x++ {
			b.points[x][l.y1]++
		}
	} else if l.IsDiagonal() {
		var ydelta int
		if l.IsDiagonalUp() {
			ydelta = 1
		} else {
			ydelta = -1
		}
		for x, y := l.MinX(), l.YOfMinX(); x <= l.MaxX(); x, y = x+1, y+ydelta {
			b.points[x][y]++
		}
	}
}

func (b *Board) CountOverlaps(threshold int) int {
	count := 0
	for _, row := range b.points {
		for _, p := range row {
			if p >= threshold {
				count++
			}
		}
	}
	return count
}
