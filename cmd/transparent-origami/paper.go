package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type XY struct {
	X int
	Y int
}

type Paper struct {
	dots map[XY]bool
	maxX int
	maxY int
}

func NewPaper(lines []string) (*Paper, int) {
	p := &Paper{
		dots: make(map[XY]bool),
	}
	var i int
	for i = 0; i < len(lines) && lines[i] != ""; i++ {
		p.AddDot(lines[i])
	}
	return p, i
}

func (p *Paper) AddDot(line string) {
	parts := strings.Split(line, ",")
	x, errx := strconv.Atoi(parts[0])
	y, erry := strconv.Atoi(parts[1])
	if errx != nil || erry != nil {
		panic("Failed to parse dot")
	}
	p.dots[XY{
		X: x,
		Y: y,
	}] = true
	if x > p.maxX {
		p.maxX = x
	}
	if y > p.maxY {
		p.maxY = y
	}
}

func (p *Paper) MoveDot(from, to XY) {
	if p.dots[from] {
		p.dots[to] = true
		delete(p.dots, from)
	}
}

func (p *Paper) Print() {
	for i := 0; i <= p.maxY; i++ {
		for j := 0; j <= p.maxX; j++ {
			if p.dots[XY{X: j, Y: i}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func (p *Paper) NumberOfDots() int {
	return len(p.dots)
}

var foldRegex = regexp.MustCompile(`fold along (x|y)=(.*)`)

func (p *Paper) Fold(instruction string) {
	matches := foldRegex.FindStringSubmatch(instruction)
	if len(matches) != 3 {
		panic("incorrect folding instruction")
	}
	direction := matches[1]
	where, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}

	switch direction {
	case "x":
		p.FoldVertical(where)
	case "y":
		p.FoldHorizontal(where)
	}
}

func (p *Paper) FoldVertical(x int) {
	for i := 2*x - p.maxX; i < x; i++ {
		mirX := 2*x - i
		for y := 0; y <= p.maxY; y++ {
			coords := XY{X: i, Y: y}
			mirCoords := XY{X: mirX, Y: y}
			p.MoveDot(mirCoords, coords)
		}
	}
	p.maxX = x - 1
}

func (p *Paper) FoldHorizontal(y int) {
	for j := 2*y - p.maxY; j < y; j++ {
		mirY := 2*y - j
		for x := 0; x <= p.maxX; x++ {
			coords := XY{X: x, Y: j}
			mirCoords := XY{X: x, Y: mirY}
			p.MoveDot(mirCoords, coords)
		}
	}
	p.maxY = y - 1
}
