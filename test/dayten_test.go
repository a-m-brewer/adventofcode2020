package test

import (
	"adventofcode2020/dayten"
	"adventofcode2020/utils"
	"testing"
)

func TestCanGetDeviceJolt(t *testing.T) {
	adapters, err := utils.LoadIntsFromFile("../inputs/day10example.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 22

	sut := dayten.NewJoltAdapterChain(adapters)

	actual := sut.DeviceAdapter()

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestCanGetNumOfDiffs(t *testing.T) {
	adapters, err := utils.LoadIntsFromFile("../inputs/day10example2.txt")
	if err != nil {
		t.Error(err)
	}

	sut := dayten.NewJoltAdapterChain(adapters)

	sut.Connect()

	expectedNum1Diff := 22
	actualNum1Diff := sut.NumAdaptersOfDiff(1)

	if actualNum1Diff != expectedNum1Diff {
		t.Errorf("expected %d got %d for diff %d", expectedNum1Diff, actualNum1Diff, 1)
	}

	expectedNum3Diff := 10
	actualNum3Diff := sut.NumAdaptersOfDiff(3)

	if actualNum3Diff != expectedNum3Diff {
		t.Errorf("expected %d got %d for diff %d", expectedNum3Diff, actualNum3Diff, 3)
	}
}

func TestDay10Part1Exercise(t *testing.T) {
	adapters, err := utils.LoadIntsFromFile("../inputs/day10.txt")
	if err != nil {
		t.Error(err)
	}

	sut := dayten.NewJoltAdapterChain(adapters)

	expected := 2592

	actual := sut.Part1()

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay10Part2Example1(t *testing.T) {
	adapters, err := utils.LoadIntsFromFile("../inputs/day10example.txt")
	if err != nil {
		t.Error(err)
	}

	sut := dayten.NewJoltAdapterChain(adapters)

	expected := 0

	actual := sut.Part2Attempt2(true)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay10Part2Example2(t *testing.T) {
	adapters, err := utils.LoadIntsFromFile("../inputs/day10example2.txt")
	if err != nil {
		t.Error(err)
	}

	sut := dayten.NewJoltAdapterChain(adapters)

	expected := 19208

	actual := sut.Part2()

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay10Part2Exercise(t *testing.T) {
	adapters, err := utils.LoadIntsFromFile("../inputs/day10.txt")
	if err != nil {
		t.Error(err)
	}

	sut := dayten.NewJoltAdapterChain(adapters)

	expected := 198428693313536

	actual := sut.Part2Attempt2(false)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}
