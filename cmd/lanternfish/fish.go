package main

import (
	"strconv"
	"strings"
)

type LanternFish int

func (l *LanternFish) Tick() bool {
	if *l > 0 {
		*l--
		return false
	} else {
		*l = 6
		return true
	}
}

// This was part of the naive solution :-)
// type School []*LanternFish

type School [9]int

func NewSchool(input string) (*School, error) {
	parts := strings.Split(input, ",")
	school := School{}
	for _, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		school[n]++
	}
	return &school, nil
}

func (s *School) Tick() {
	var temp int = (*s)[0]
	for i := 0; i < 8; i++ {
		(*s)[i] = (*s)[i+1]
	}
	(*s)[6] = (*s)[6] + temp
	(*s)[8] = temp
}

func (s *School) Count() int {
	count := 0
	for _, n := range *s {
		count = count + n
	}
	return count
}
