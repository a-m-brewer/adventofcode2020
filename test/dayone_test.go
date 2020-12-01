package test

import (
	"adventofcode2020/dayone"
	"adventofcode2020/utils"
	"testing"
)

func TestDay1Part1Example1(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	expected := 514579

	actual, err := dayone.CalculateExpenses(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	if actual != expected {
		t.Errorf("expected: %d got: %d", expected, actual)
	}
}

func TestDay1Part1Exercise1(t *testing.T) {
	input, err := utils.LoadIntsFromFile("../inputs/day1.txt")

	expected := 980499

	if err != nil {
		t.Error(err.Error())
	}

	actual, err := dayone.CalculateExpenses(input)
	if err != nil {
		t.Error(err.Error())
	}

	if actual != expected {
		t.Errorf("expected: %d got: %d", expected, actual)
	}
}

func TestDay1Part2Example1(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	expected := 241861950

	actual, err := dayone.CalculateExpenses2(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	if actual != expected {
		t.Errorf("expected: %d got: %d", expected, actual)
	}
}

func TestDay1Part2Exercise1(t *testing.T) {
	input, err := utils.LoadIntsFromFile("../inputs/day1.txt")

	expected := 200637446

	if err != nil {
		t.Error(err.Error())
	}

	actual, err := dayone.CalculateExpenses2(input)
	if err != nil {
		t.Error(err.Error())
	}

	if actual != expected {
		t.Errorf("expected: %d got: %d", expected, actual)
	}
}
