package test

import (
	"adventofcode2020/daythirteen"
	"adventofcode2020/utils"
	"testing"
)

func TestDay13Part1Example1(t *testing.T) {
	var expected uint64 = 295

	input, err := utils.LoadLines("../inputs/day13example.txt")
	if err != nil {
		t.Error(err)
	}

	sut := daythirteen.NewBusSchedule(input)

	actual := sut.Part1()

	if actual != expected {
		t.Errorf("expected: %d got %d", expected, actual)
	}
}

func TestDay13Part1Exercise(t *testing.T) {
	var expected uint64 = 3246

	input, err := utils.LoadLines("../inputs/day13.txt")
	if err != nil {
		t.Error(err)
	}

	sut := daythirteen.NewBusSchedule(input)

	actual := sut.Part1()

	if actual != expected {
		t.Errorf("expected: %d got %d", expected, actual)
	}
}

func TestDay13Part2Example1(t *testing.T) {
	var expected uint64 = 1068781

	input, err := utils.LoadLines("../inputs/day13example.txt")
	if err != nil {
		t.Error(err)
	}

	sut := daythirteen.NewBusSchedule(input)
	sut.SetMinLeavingTime(1068773)

	actual := sut.Part2()

	if actual != expected {
		t.Errorf("expected: %d got %d", expected, actual)
	}
}

func TestDay13Part2Example2(t *testing.T) {
	var expected uint64 = 3417

	input := []string{
		"0",
		"17,x,13,19",
	}

	sut := daythirteen.NewBusSchedule(input)
	sut.SetMinLeavingTime(3000)

	actual := sut.Part2()

	if actual != expected {
		t.Errorf("expected: %d got %d", expected, actual)
	}
}

func TestDay13Part2Example3(t *testing.T) {
	var expected uint64 = 754018

	input := []string{
		"0",
		"67,7,59,61",
	}

	sut := daythirteen.NewBusSchedule(input)
	sut.SetMinLeavingTime(750000)

	actual := sut.Part2()

	if actual != expected {
		t.Errorf("expected: %d got %d", expected, actual)
	}
}

func TestDay13Part2Example4(t *testing.T) {
	var expected uint64 = 779210

	input := []string{
		"0",
		"67,x,7,59,61",
	}

	sut := daythirteen.NewBusSchedule(input)
	sut.SetMinLeavingTime(750000)

	actual := sut.Part2()

	if actual != expected {
		t.Errorf("expected: %d got %d", expected, actual)
	}
}

func TestDay13Part2Example5(t *testing.T) {
	var expected uint64 = 1261476

	input := []string{
		"0",
		"67,7,x,59,61",
	}

	sut := daythirteen.NewBusSchedule(input)
	sut.SetMinLeavingTime(1200000)

	actual := sut.Part2()

	if actual != expected {
		t.Errorf("expected: %d got %d", expected, actual)
	}
}

func TestDay13Part2Example6(t *testing.T) {
	var expected uint64 = 1202161486

	input := []string{
		"0",
		"1789,37,47,1889",
	}

	sut := daythirteen.NewBusSchedule(input)
	sut.SetMinLeavingTime(1202000000)

	actual := sut.Part2()

	if actual != expected {
		t.Errorf("expected: %d got %d", expected, actual)
	}
}

func TestDay13Part2Exercise(t *testing.T) {
	var expected uint64 = 1010182346291467

	input, err := utils.LoadLines("../inputs/day13.txt")
	if err != nil {
		t.Error(err)
	}

	sut := daythirteen.NewBusSchedule(input)
	sut.SetMinLeavingTime(100000000000000)

	actual := sut.Part2()

	if actual != expected {
		t.Errorf("expected: %d got %d", expected, actual)
	}
}
