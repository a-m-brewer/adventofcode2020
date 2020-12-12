package test

import (
	"adventofcode2020/dayeleven"
	"adventofcode2020/utils"
	"testing"
)

func TestDay11Part1Example1(t *testing.T) {
	input, err := utils.LoadStringGrid("../inputs/day11example.txt")
	if err != nil {
		t.Error(err)
	}

	sut := dayeleven.NewSeating(input, dayeleven.PART1)

	sut.Run()

	expected := 37
	actual := sut.NumOccupied()

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay11Part1Exercise(t *testing.T) {
	input, err := utils.LoadStringGrid("../inputs/day11.txt")
	if err != nil {
		t.Error(err)
	}

	sut := dayeleven.NewSeating(input, dayeleven.PART1)

	sut.Run()

	expected := 2263
	actual := sut.NumOccupied()

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay11Part2Example1(t *testing.T) {
	input, err := utils.LoadStringGrid("../inputs/day11example2.txt")
	if err != nil {
		t.Error(err)
	}

	sut := dayeleven.NewSeating(input, dayeleven.PART2)
	sut.Log = t.Logf

	sut.Run()

	expected := 26
	actual := sut.NumOccupied()

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay11Part2Exercise(t *testing.T) {
	input, err := utils.LoadStringGrid("../inputs/day11.txt")
	if err != nil {
		t.Error(err)
	}

	sut := dayeleven.NewSeating(input, dayeleven.PART2)
	sut.Log = t.Logf

	sut.Run()

	expected := 26
	actual := sut.NumOccupied()

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}
