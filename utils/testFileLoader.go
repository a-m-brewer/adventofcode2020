package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

// LoadIntsFromFile takes a path and loads the file expecting the file to contain one int per line
func LoadIntsFromFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		i, err := strconv.Atoi(txt)
		if err != nil {
			return nil, err
		}
		lines = append(lines, i)
	}

	err = scanner.Err()

	return lines, err
}

// LoadLines reads a file and returns it as an array where each line is an item in the array
func LoadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	return lines, err
}

// LoadString loads a file into a string
func LoadString(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	s := string(b)

	return s, nil
}
