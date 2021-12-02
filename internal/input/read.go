package input

import (
	"bufio"
	"os"
	"strconv"
)

func ReadNumbers(filename string) ([]int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := []int64{}
	for scanner.Scan() {
		n, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, nil

}

func ReadStrings(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, nil
}
