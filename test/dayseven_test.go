package test

import (
	"adventofcode2020/dayseven"
	"adventofcode2020/utils"
	"testing"
)

func TestDay7Part1Example1(t *testing.T) {
	lines, err := utils.LoadLines("../inputs/day7example.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 4

	actual := dayseven.Part1(lines)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay7Part1Exercise(t *testing.T) {
	lines, err := utils.LoadLines("../inputs/day7.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 121

	actual := dayseven.Part1(lines)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay7Part2Example(t *testing.T) {
	lines, err := utils.LoadLines("../inputs/day7example2.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 126

	actual := dayseven.Part2(lines)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay7Part2Exercise(t *testing.T) {
	lines, err := utils.LoadLines("../inputs/day7.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 3805

	actual := dayseven.Part2(lines)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}
