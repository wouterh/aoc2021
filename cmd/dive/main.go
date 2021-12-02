package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wouterh/aoc2021/internal/input"
)

func parseInstruction(instruction string) (string, int, error) {
	parts := strings.Split(instruction, " ")
	if len(parts) != 2 {
		return "", 0, errors.New("unable to parse instruction")
	}
	what := parts[0]
	amount, err := strconv.ParseInt(parts[1], 10, 0)
	if err != nil {
		return "", 0, err
	}
	return what, int(amount), nil
}

func findPosition(instructions []string) (x, d int, err error) {
	x, d = 0, 0
	for _, instruction := range instructions {
		what, amount, err := parseInstruction(instruction)
		if err != nil {
			return 0, 0, err
		}
		switch what {
		case "forward":
			x = x + amount
		case "down":
			d = d + amount
		case "up":
			d = d - amount
		default:
			return 0, 0, errors.New("unexpected instruction")
		}
	}
	return x, d, nil
}

func findPositionWithAim(instructions []string) (x, d int, err error) {
	x, d, aim := 0, 0, 0
	for _, instruction := range instructions {
		what, amount, err := parseInstruction(instruction)
		if err != nil {
			return 0, 0, err
		}
		switch what {
		case "forward":
			x = x + amount
			d = d + aim*amount
		case "down":
			aim = aim + amount
		case "up":
			aim = aim - amount
		default:
			return 0, 0, errors.New("unexpected instruction")
		}
	}
	return x, d, nil
}

func main() {
	instructions, err := input.ReadStrings(os.Args[1])
	if err != nil {
		panic("Could not fetch input")
	}
	x, d, err := findPosition(instructions)
	if err != nil {
		panic(err)
	}
	fmt.Println(x * d)
	x, d, err = findPositionWithAim(instructions)
	if err != nil {
		panic(err)
	}
	fmt.Println(x * d)
}
