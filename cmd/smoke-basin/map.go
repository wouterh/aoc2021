package main

import (
	"strconv"

	"github.com/wouterh/aoc2021/internal/input"
)

type Map [][]int

func NewMap(filename string) (*Map, error) {
	lines, err := input.ReadStrings(filename)
	if err != nil {
		return nil, err
	}
	var m Map = make(Map, len(lines))
	for i, l := range lines {
		var row []int = make([]int, len(l))
		for j, r := range l {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}
			row[j] = n
		}
		m[i] = row

	}
	return &m, nil
}

func (m *Map) Dimensions() (int, int) {
	return len(*m), len((*m)[0])
}

func (m *Map) FindLowPoints() []int {
	var lowPoints []int
	rows, cols := m.Dimensions()
	for i, row := range *m {
		for j, v := range row {
			if j < cols-1 && v >= row[j+1] {
				continue
			} else if j > 0 && v >= row[j-1] {
				continue
			} else if i < rows-1 && v >= (*m)[i+1][j] {
				continue
			} else if i > 0 && v >= (*m)[i-1][j] {
				continue
			} else {
				lowPoints = append(lowPoints, v)
			}
		}
	}
	return lowPoints
}
