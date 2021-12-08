package main

import (
	"errors"
	"strings"
)

var DigitToSegmentCount map[int]int = map[int]int{
	0: 6,
	1: 2,
	2: 5,
	3: 5,
	4: 4,
	5: 5,
	6: 6,
	7: 3,
	8: 7,
	9: 6,
}

var signalToPrime map[rune]int = map[rune]int{
	'a': 2,
	'b': 3,
	'c': 5,
	'd': 7,
	'e': 11,
	'f': 13,
	'g': 17,
}

type Pattern string

func Contains(p, o string) bool {
	for _, r := range o {
		if !strings.ContainsRune(p, r) {
			return false
		}
	}
	return true
}

func Fingerprint(p string) int {
	fingerprint := 1
	for _, r := range p {
		fingerprint = fingerprint * signalToPrime[r]
	}
	return fingerprint
}

type Entry struct {
	signalPatterns []string
	outputs        []string
}

func NewEntry(input string) (*Entry, error) {
	parts := strings.Split(input, " | ")
	if len(parts) != 2 {
		return nil, errors.New("couldn't parse input")
	}
	signals, outputs := strings.Split(parts[0], " "), strings.Split(parts[1], " ")
	if len(signals) != 10 || len(outputs) != 4 {
		return nil, errors.New("couldn't parse input")
	}
	return &Entry{
		signalPatterns: signals,
		outputs:        outputs,
	}, nil
}

func (e *Entry) CountUniques() int {
	count := 0
	for _, o := range e.outputs {
		l := len(o)
		if l == DigitToSegmentCount[1] ||
			l == DigitToSegmentCount[4] ||
			l == DigitToSegmentCount[7] ||
			l == DigitToSegmentCount[8] {
			count++
		}
	}
	return count
}

func (e *Entry) FindOutput() int {
	mapping := e.Decode()
	// Reverse the mappping
	reversedMapping := map[int]int{}
	for k, v := range mapping {
		reversedMapping[v] = k
	}
	result := 1000*reversedMapping[Fingerprint(e.outputs[0])] +
		100*reversedMapping[Fingerprint(e.outputs[1])] +
		10*reversedMapping[Fingerprint(e.outputs[2])] +
		reversedMapping[Fingerprint(e.outputs[3])]

	return result
}

func (e *Entry) Decode() map[int]int {
	var zero, one, two, three, four, five, six, seven, eight, nine string
	fiveSegments, sixSegments := []string{}, []string{}
	// Find the unique ones + group rest by length
	for _, pattern := range e.signalPatterns {
		switch len(pattern) {
		case DigitToSegmentCount[1]:
			one = pattern
		case DigitToSegmentCount[4]:
			four = pattern
		case DigitToSegmentCount[7]:
			seven = pattern
		case DigitToSegmentCount[8]:
			eight = pattern
		case 5:
			fiveSegments = append(fiveSegments, pattern)
		case 6:
			sixSegments = append(sixSegments, pattern)
		}
	}
	// Find 9 as the one with six segments that contains 4
	for _, pattern := range sixSegments {
		if Contains(pattern, four) {
			nine = pattern
			break
		}
	}
	// Find 0 as the one with six segments that contains 7 and is not 9
	for _, pattern := range sixSegments {
		if pattern != nine && Contains(pattern, seven) {
			zero = pattern
			break
		}
	}
	// The remaining one of six segments is 6
	for _, pattern := range sixSegments {
		if pattern != nine && pattern != zero {
			six = pattern
			break
		}
	}
	// Find the upper right segment as the one in 0 that is not in 6
	var upperRight rune
	for _, r := range zero {
		if !strings.ContainsRune(six, r) {
			upperRight = r
			break
		}
	}
	// Find 3 as the one with five segments that contains 7
	for _, pattern := range fiveSegments {
		if Contains(pattern, seven) {
			three = pattern
			break
		}
	}
	// Find 2 as the one with five segments that contains upperRight and is not 3
	for _, pattern := range fiveSegments {
		if pattern != three && strings.ContainsRune(pattern, upperRight) {
			two = pattern
			break
		}
	}
	// Find 5 as the remaining one with five segments
	for _, pattern := range fiveSegments {
		if pattern != two && pattern != three {
			five = pattern
			break
		}
	}
	// We should have all of them
	if zero == "" || one == "" || two == "" ||
		three == "" || four == "" || five == "" ||
		six == "" || seven == "" || eight == "" ||
		nine == "" {
		panic("Couldn't decode entry")
	}
	result := map[int]int{
		0: Fingerprint(zero),
		1: Fingerprint(one),
		2: Fingerprint(two),
		3: Fingerprint(three),
		4: Fingerprint(four),
		5: Fingerprint(five),
		6: Fingerprint(six),
		7: Fingerprint(seven),
		8: Fingerprint(eight),
		9: Fingerprint(nine),
	}
	return result
}
