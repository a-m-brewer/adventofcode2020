package test

import (
	"adventofcode2020/daythree"
	"adventofcode2020/utils"
	"testing"
)

func TestDay3Part1Example1(t *testing.T) {
	lines, err := utils.LoadLines("../inputs/day3example1.txt")
	if err != nil {
		t.Error(err.Error())
	}

	tmap := daythree.NewTobogganMap(lines)

	expected := 7

	actual := tmap.PlotCourse(daythree.Course{X: 3, Y: 1})

	if actual != expected {
		t.Errorf("Expected %d trees found %d", expected, actual)
	}
}

func TestDay3Part1Exersise(t *testing.T) {
	lines, err := utils.LoadLines("../inputs/day3.txt")
	if err != nil {
		t.Error(err.Error())
	}

	tmap := daythree.NewTobogganMap(lines)

	expected := 225

	actual := tmap.PlotCourse(daythree.Course{X: 3, Y: 1})

	if actual != expected {
		t.Errorf("Expected %d trees found %d", expected, actual)
	}
}

func TestDay3Part2Example1(t *testing.T) {
	lines, err := utils.LoadLines("../inputs/day3example1.txt")
	if err != nil {
		t.Error(err.Error())
	}

	tmap := daythree.NewTobogganMap(lines)

	expected := 336

	actual := tmap.ProbabilityOfSuddenArboreal(daythree.Courses{
		daythree.Course{X: 1, Y: 1},
		daythree.Course{X: 3, Y: 1},
		daythree.Course{X: 5, Y: 1},
		daythree.Course{X: 7, Y: 1},
		daythree.Course{X: 1, Y: 2},
	})

	if actual != expected {
		t.Errorf("Expected %d probability found %d", expected, actual)
	}
}

func TestDay3Part2Exercise(t *testing.T) {
	lines, err := utils.LoadLines("../inputs/day3.txt")
	if err != nil {
		t.Error(err.Error())
	}

	tmap := daythree.NewTobogganMap(lines)

	expected := 336

	actual := tmap.ProbabilityOfSuddenArboreal(daythree.Courses{
		daythree.Course{X: 1, Y: 1},
		daythree.Course{X: 3, Y: 1},
		daythree.Course{X: 5, Y: 1},
		daythree.Course{X: 7, Y: 1},
		daythree.Course{X: 1, Y: 2},
	})

	if actual != expected {
		t.Errorf("Expected %d probability found %d", expected, actual)
	}
}
