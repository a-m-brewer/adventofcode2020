package dayfive

import (
	"errors"
	"sort"
)

func FindSeats(inputs []string) []int {
	seats := make([]int, len(inputs))
	for i, v := range inputs {
		seats[i] = FindSeat(v)
	}

	return seats
}

func FindHighestSeatId(inputs []string) int {
	seats := FindSeats(inputs)
	highestId := 0
	for _, seatId := range seats {
		if highestId < seatId {
			highestId = seatId
		}
	}

	return highestId
}

func FindFreeSpace(inputs []string) (int, error) {
	seats := FindSeats(inputs)
	sort.Ints(seats)

	for i := 1; i < len(seats); i++ {
		diff := seats[i] - seats[i-1]
		if diff > 1 {
			return seats[i] - 1, nil
		}
	}

	return -1, errors.New("no free seats found")
}

func FindSeat(input string) int {
	inputArr := []rune(input)

	splitPoint := len(inputArr) - 3
	rowIdentifier := toLowHigh(inputArr[:splitPoint], rune('F'))
	colIdentifier := toLowHigh(inputArr[splitPoint:], rune('L'))

	row := findSeatRecursive(0, 127, rowIdentifier, 64)
	col := findSeatRecursive(0, 7, colIdentifier, 4)

	return row*8 + col
}

func findSeatRecursive(low, high int, input []bool, nextValue int) int {
	currInstruction := input[0]

	if len(input) == 1 {
		if currInstruction {
			return low
		}

		return high
	} else {
		if currInstruction {
			high -= nextValue
		} else {
			low += nextValue
		}
	}

	nextValue /= 2

	return findSeatRecursive(low, high, input[1:], nextValue)
}

func toLowHigh(input []rune, high rune) []bool {
	ba := make([]bool, len(input))
	for i, v := range input {
		ba[i] = v == high
	}

	return ba
}
