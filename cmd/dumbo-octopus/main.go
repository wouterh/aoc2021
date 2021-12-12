package main

import (
	"fmt"
	"os"

	"github.com/wouterh/aoc2021/internal/input"
)

func doHundredSteps(cavern *Cavern) int {
	count := 0
	for i := 0; i < 100; i++ {
		flashes := cavern.Step()
		count += flashes
	}
	return count
}

func findFirstAllFlash(cavern *Cavern) int {
	for i := 0; i < 1000; i++ {
		flashes := cavern.Step()
		if flashes == 100 {
			return i + 1
		}
	}
	return -1
}

func main() {
	lines, err := input.ReadStrings(os.Args[1])
	if err != nil {
		panic(err)
	}
	cavern, err := NewCavern(lines)
	copy := *cavern
	if err != nil {
		panic(err)
	}
	nb := doHundredSteps(cavern)
	fmt.Println(nb)
	sync := findFirstAllFlash(&copy)
	fmt.Println(sync)
}
