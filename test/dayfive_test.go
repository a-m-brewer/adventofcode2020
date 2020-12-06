package test

import (
	"adventofcode2020/dayfive"
	"adventofcode2020/utils"
	"testing"
)

func TestDay5Part1Example1(t *testing.T) {
	input := "FBFBBFFRLR"
	expected := 357

	actual := dayfive.FindSeat(input)

	if actual != expected {
		t.Errorf("expected seat %d actual: %d", expected, actual)
	}
}

func TestDay5Part1Example2(t *testing.T) {
	input := "BFFFBBFRRR"
	expected := 567

	actual := dayfive.FindSeat(input)

	if actual != expected {
		t.Errorf("expected seat %d actual: %d", expected, actual)
	}
}

func TestDay5Part1Example3(t *testing.T) {
	input := "FFFBBBFRRR"
	expected := 119

	actual := dayfive.FindSeat(input)

	if actual != expected {
		t.Errorf("expected seat %d actual: %d", expected, actual)
	}
}

func TestDay5Part1Example4(t *testing.T) {
	input := "BBFFBBFRLL"
	expected := 820

	actual := dayfive.FindSeat(input)

	if actual != expected {
		t.Errorf("expected seat %d actual: %d", expected, actual)
	}
}

func TestDay5Part1Example6(t *testing.T) {
	expected := 820

	inputs := []string{
		"FBFBBFFRLR",
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}

	actual := dayfive.FindHighestSeatId(inputs)

	if actual != expected {
		t.Errorf("expected: %d got: %d", expected, actual)
	}
}

func TestDay5Part1Exercise(t *testing.T) {
	expected := 896

	inputs, err := utils.LoadLines("../inputs/day5.txt")
	if err != nil {
		t.Error(err)
	}

	actual := dayfive.FindHighestSeatId(inputs)

	if actual != expected {
		t.Errorf("expected: %d got: %d", expected, actual)
	}
}

func TestDay5Part2Exercise(t *testing.T) {
	expected := 0

	inputs, err := utils.LoadLines("../inputs/day5.txt")
	if err != nil {
		t.Error(err)
	}

	actual, err := dayfive.FindFreeSpace(inputs)
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("expected: %d got: %d", expected, actual)
	}
}
