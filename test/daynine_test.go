package test

import (
	"adventofcode2020/daynine"
	"adventofcode2020/utils"
	"testing"
)

func TestDay9Part1Example1(t *testing.T) {
	input, err := utils.LoadIntsFromFile("../inputs/day9example1.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 127

	sut := daynine.NewXMASDecrypter(input, 5)

	found, actual := sut.ContainsInvalidNumber()

	if !found {
		t.Error("did not find invalid number in sequence")
	}

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay9Part1Exercise(t *testing.T) {
	input, err := utils.LoadIntsFromFile("../inputs/day9.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 1504371145

	sut := daynine.NewXMASDecrypter(input, 25)

	found, actual := sut.ContainsInvalidNumber()

	if !found {
		t.Error("did not find invalid number in sequence")
	}

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay9Part2Example1(t *testing.T) {
	input, err := utils.LoadIntsFromFile("../inputs/day9example2.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 62

	sut := daynine.NewXMASDecrypter(input, 5)

	actual, err := sut.FindEncryptionWeakness()
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay9Part2Exercise(t *testing.T) {
	input, err := utils.LoadIntsFromFile("../inputs/day9.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 62

	sut := daynine.NewXMASDecrypter(input, 25)

	actual, err := sut.FindEncryptionWeakness()
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}
