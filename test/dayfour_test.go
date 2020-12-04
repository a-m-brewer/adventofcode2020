package test

import (
	"adventofcode2020/dayfour"
	"adventofcode2020/utils"
	"testing"
)

func TestDay4Part1Example1(t *testing.T) {
	batch, err := utils.LoadString("../inputs/day4example.txt")
	if err != nil {
		t.Error(err.Error())
	}

	passports := dayfour.NewPassports(batch)
	pc := dayfour.NewPassportChecker()

	expected := 2

	actual := pc.NumValidPassports(passports, false)

	if actual != expected {
		t.Errorf("expected %d valid passports got %d", expected, actual)
	}
}

func TestDay4Part1Exercise(t *testing.T) {
	batch, err := utils.LoadString("../inputs/day4.txt")
	if err != nil {
		t.Error(err.Error())
	}

	passports := dayfour.NewPassports(batch)
	pc := dayfour.NewPassportChecker()

	expected := 237

	actual := pc.NumValidPassports(passports, false)

	if actual != expected {
		t.Errorf("expected %d valid passports got %d", expected, actual)
	}
}

func TestDay4Part2ExampleInvalid(t *testing.T) {
	batch, err := utils.LoadString("../inputs/day4exampleinvalid.txt")
	if err != nil {
		t.Error(err.Error())
	}

	passports := dayfour.NewPassports(batch)
	pc := dayfour.NewPassportChecker()

	expected := 0

	actual := pc.NumValidPassports(passports, true)

	if actual != expected {
		t.Errorf("expected %d valid passports got %d", expected, actual)
	}
}
func TestDay4Part2ExampleValid(t *testing.T) {
	batch, err := utils.LoadString("../inputs/day4valid.txt")
	if err != nil {
		t.Error(err.Error())
	}

	passports := dayfour.NewPassports(batch)
	pc := dayfour.NewPassportChecker()

	expected := 4

	actual := pc.NumValidPassports(passports, false)

	if actual != expected {
		t.Errorf("expected %d valid passports got %d", expected, actual)
	}
}

func TestDay4Part2Exercise(t *testing.T) {
	batch, err := utils.LoadString("../inputs/day4.txt")
	if err != nil {
		t.Error(err.Error())
	}

	passports := dayfour.NewPassports(batch)
	pc := dayfour.NewPassportChecker()

	expected := 172

	actual := pc.NumValidPassports(passports, true)

	if actual != expected {
		t.Errorf("expected %d valid passports got %d", expected, actual)
	}
}
