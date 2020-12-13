package test

import (
	"adventofcode2020/daytwelve"
	"adventofcode2020/utils"
	"testing"
)

func TestDay12Part1Example1(t *testing.T) {
	input, err := utils.LoadLines("../inputs/day12example.txt")
	if err != nil {
		t.Error(err)
	}

	sut := daytwelve.NewShip(input)

	expected := 25

	sut.Run()

	actual := sut.DistanceFromStart()

	if actual != expected {
		t.Errorf("expected: %d got %d", expected, actual)
	}
}

func TestDay12Part1Exercise(t *testing.T) {
	input, err := utils.LoadLines("../inputs/day12.txt")
	if err != nil {
		t.Error(err)
	}

	sut := daytwelve.NewShip(input)

	expected := 1687
	sut.Run()

	actual := sut.DistanceFromStart()

	if actual != expected {
		t.Errorf("expected: %d got %d", expected, actual)
	}
}

func TestDay12Part2Example1(t *testing.T) {
	input, err := utils.LoadLines("../inputs/day12.txt")
	if err != nil {
		t.Error(err)
	}

	sut := daytwelve.NewShip(input)
	sut.Part = daytwelve.PART2

	expected := 20873

	sut.Run()

	actual := sut.DistanceFromStart()

	if actual != expected {
		t.Errorf("expected: %d got %d", expected, actual)
	}
}
