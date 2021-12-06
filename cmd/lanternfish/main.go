package main

import (
	"fmt"
	"os"

	"github.com/wouterh/aoc2021/internal/input"
)

func countAfterDays(school *School, number int) int {
	for i := 0; i < number; i++ {
		school.Tick()
	}
	return school.Count()
}

func main() {
	inputs, err := input.ReadStrings(os.Args[1])
	if err != nil {
		panic("Couldn't read input")
	}
	school, err := NewSchool(inputs[0])
	if err != nil {
		panic(err)
	}
	count := countAfterDays(school, 80)
	fmt.Println(count)
	count = countAfterDays(school, 256-80)
	fmt.Println(count)
}
