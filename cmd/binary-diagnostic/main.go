package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/wouterh/aoc2021/internal/input"
)

var oneByte = ([]byte("1"))[0]
var zeroByte = ([]byte("0"))[0]

func calculatePowerConsumption(bitmasks []string) (int, error) {
	if len(bitmasks) == 0 {
		return 0, nil
	}
	nbBits := len(bitmasks[0])
	gamma := make([]byte, nbBits)
	epsilon := make([]byte, nbBits)
	for i := 0; i < nbBits; i++ {
		nbOnes := 0
		nbZeroes := 0
		for _, bitmask := range bitmasks {
			if bitmask[i] == oneByte {
				nbOnes++
			} else if bitmask[i] == zeroByte {
				nbZeroes++
			}
		}
		if nbOnes > nbZeroes {
			gamma[i] = oneByte
			epsilon[i] = zeroByte
		} else {
			gamma[i] = zeroByte
			epsilon[i] = oneByte
		}
	}
	gammaNumber, err := strconv.ParseUint(string(gamma), 2, 32)
	if err != nil {
		return 0, err
	}
	epsilonNumber, err := strconv.ParseUint(string(epsilon), 2, 32)
	if err != nil {
		return 0, err
	}
	return int(gammaNumber) * int(epsilonNumber), nil
}

func calculateRating(bitmasks []string, most bool) (int, error) {
	if len(bitmasks) == 0 {
		return 0, nil
	}
	nbBits := len(bitmasks[0])
	for i := 0; i < nbBits; i++ {
		withOne := []string{}
		withZero := []string{}
		nbOnes := 0
		nbZeroes := 0
		for _, bitmask := range bitmasks {
			if bitmask[i] == oneByte {
				nbOnes++
				withOne = append(withOne, bitmask)
			} else if bitmask[i] == zeroByte {
				nbZeroes++
				withZero = append(withZero, bitmask)
			}
		}
		if most {
			if nbOnes >= nbZeroes {
				bitmasks = withOne
			} else {
				bitmasks = withZero
			}
		} else {
			if nbZeroes <= nbOnes {
				bitmasks = withZero
			} else {
				bitmasks = withOne
			}
		}
		if len(bitmasks) == 1 {
			rating, err := strconv.ParseUint(bitmasks[0], 2, 32)
			if err != nil {
				return 0, err
			}
			return int(rating), nil
		}
	}
	return 0, errors.New("no rating found")
}

func calculateLifeSupportRating(bitmasks []string) (int, error) {
	oxygen, err := calculateRating(bitmasks, true)
	if err != nil {
		return 0, err
	}
	co2, err := calculateRating(bitmasks, false)
	if err != nil {
		return 0, err
	}
	return oxygen * co2, nil
}

func main() {
	bitmasks, err := input.ReadStrings(os.Args[1])
	if err != nil {
		panic("Could not fetch input")
	}
	power, err := calculatePowerConsumption(bitmasks)
	if err != nil {
		panic("Could not calculate power consumption")
	}
	fmt.Println(power)
	lifeSupport, err := calculateLifeSupportRating(bitmasks)
	if err != nil {
		panic("Could not calculate life support rating")
	}
	fmt.Println(lifeSupport)
}
