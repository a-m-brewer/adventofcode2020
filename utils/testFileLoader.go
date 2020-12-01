package utils

import (
	"bufio"
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
