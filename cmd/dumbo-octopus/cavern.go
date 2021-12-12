package main

import (
	"fmt"
	"strconv"
)

type Cavern [100]int

type Coord struct {
	X, Y int
}

func NewCavern(lines []string) (*Cavern, error) {
	cavern := &Cavern{}
	for i, line := range lines {
		for j, v := range line {
			iv, err := strconv.Atoi(string(v))
			if err != nil {
				return nil, err
			}
			cavern.Set(j, i, iv)
		}
	}
	return cavern, nil
}

func (c *Cavern) Set(x, y, v int) {
	(*c)[x+10*y] = v
}

func (c *Cavern) Get(x, y int) int {
	return (*c)[x+10*y]
}

func (c *Cavern) Increment(x, y int) int {
	val := c.Get(x, y)
	val++
	c.Set(x, y, val)
	return val
}

func (c *Cavern) IncrementAll() []Coord {
	toFlash := []Coord{}
	for i := 0; i < 100; i++ {
		(*c)[i]++
		if (*c)[i] > 9 {
			toFlash = append(toFlash, Coord{X: i % 10, Y: i / 10})
		}
	}
	return toFlash
}

func validCoords(c Coord) bool {
	x, y := c.X, c.Y
	return x >= 0 && x < 10 && y >= 0 && y < 10
}

func neighbourCoords(x, y int) (coords []Coord) {
	candidates := []Coord{
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
		{x - 1, y},
		{x + 1, y},
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}
	filtered := candidates[:0]
	for _, c := range candidates {
		if validCoords(c) {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func (c *Cavern) Step() int {
	nbFlashes := 0
	// First increment all energy levels
	toFlash := c.IncrementAll()
	// Now flash all
	for i := 0; i < len(toFlash); i++ {
		co := toFlash[i]
		nbFlashes++
		for _, n := range neighbourCoords(co.X, co.Y) {
			nVal := c.Increment(n.X, n.Y)
			if nVal == 10 {
				toFlash = append(toFlash, n)
			}
		}
	}
	// Now clear all flashed
	for _, co := range toFlash {
		c.Set(co.X, co.Y, 0)
	}
	// Return the number of flashes
	return nbFlashes
}

func (c *Cavern) Print() {
	for i, v := range *c {
		fmt.Printf("%2d", v)
		if i%10 == 9 {
			fmt.Println("")
		}
	}
	fmt.Println("")
}
