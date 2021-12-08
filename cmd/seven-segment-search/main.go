package main

import (
	"fmt"
	"os"

	"github.com/wouterh/aoc2021/internal/input"
)

func parseEntries(sEntries []string) ([]*Entry, error) {
	result := make([]*Entry, len(sEntries))
	for i, v := range sEntries {
		entry, err := NewEntry(v)
		if err != nil {
			return nil, err
		}
		result[i] = entry
	}
	return result, nil
}

func CountUniques(entries []*Entry) int {
	count := 0
	for _, e := range entries {
		count = count + e.CountUniques()
	}
	return count
}

func SumOutputs(entries []*Entry) int {
	sum := 0
	for _, e := range entries {
		sum = sum + e.FindOutput()
	}
	return sum
}

func main() {
	sEntries, err := input.ReadStrings(os.Args[1])
	if err != nil {
		panic(err)
	}
	entries, err := parseEntries(sEntries)
	if err != nil {
		panic(err)
	}
	count := CountUniques(entries)
	fmt.Println(count)
	outputSum := SumOutputs(entries)
	fmt.Println(outputSum)
}
