package input

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
	if scanner.Err() != nil {
		return nil, err
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
	if scanner.Err() != nil {
		return nil, err
	}
	return result, nil
}

func ReadNumbersOnLine(filename string) ([]int, error) {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(string(bytes), ",")
	result := make([]int, len(parts))
	for i := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(parts[i]))
		if err != nil {
			return nil, err
		}
		result[i] = n
	}
	return result, nil
}
