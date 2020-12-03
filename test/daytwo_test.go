package test

import (
	"adventofcode2020/daytwo"
	"adventofcode2020/utils"
	"testing"
)

func TestCanParsePolicy(t *testing.T) {
	expected := daytwo.Policy{
		Min:      8,
		Max:      9,
		Char:     'x',
		Password: "xxxxxxxrk",
	}

	actual, err := daytwo.NewPolicy("8-9 x: xxxxxxxrk")
	if err != nil {
		t.Error(err.Error())
	}

	minMatch := expected.Min == actual.Min

	if !minMatch {
		t.Errorf("min expected: %d got: %d", expected.Min, actual.Min)
	}

	maxMatch := expected.Max == actual.Max

	if !maxMatch {
		t.Errorf("max expected: %d got: %d", expected.Max, actual.Max)
	}

	charMatch := expected.Char == actual.Char

	if !charMatch {
		t.Errorf("char expected: %U got: %U", expected.Char, actual.Char)
	}

	passwordMatch := expected.Password == actual.Password

	if !passwordMatch {
		t.Errorf("password expected: %s got: %s", expected.Password, actual.Password)
	}
}

func TestDay2Part1Example1(t *testing.T) {
	expected := true

	policy, err := daytwo.NewPolicy("1-3 a: abcde")

	if err != nil {
		t.Error(err.Error())
	}

	actual := policy.PasswordValidSledRental()
	if actual != expected {
		t.Errorf("expected %t got %t", expected, actual)
	}
}

func TestDay2Part1Example2(t *testing.T) {
	expected := false

	policy, err := daytwo.NewPolicy("1-3 b: cdefg")

	if err != nil {
		t.Error(err.Error())
	}

	actual := policy.PasswordValidSledRental()
	if actual != expected {
		t.Errorf("expected %t got %t", expected, actual)
	}
}

func TestDay2Part1Example3(t *testing.T) {
	expected := true

	policy, err := daytwo.NewPolicy("2-9 c: ccccccccc")

	if err != nil {
		t.Error(err.Error())
	}

	actual := policy.PasswordValidSledRental()
	if actual != expected {
		t.Errorf("expected %t got %t", expected, actual)
	}
}

func TestDay2Part1Example4(t *testing.T) {
	expected := 2

	policies, err := daytwo.NewPolicies([]string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	})

	if err != nil {
		t.Error(err.Error())
	}

	actual := policies.NumValidSledRental()
	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay2Part1Exercise(t *testing.T) {
	expected := 645

	policyStrings, err := utils.LoadLines("../inputs/day2.txt")
	if err != nil {
		t.Error(err.Error())
	}

	policies, err := daytwo.NewPolicies(policyStrings)
	if err != nil {
		t.Error(err.Error())
	}

	actual := policies.NumValidSledRental()

	if actual != expected {
		t.Errorf("expected: %d got: %d", expected, actual)
	}
}

func TestDay2Part2Example1(t *testing.T) {
	expected := true

	policy, err := daytwo.NewPolicy("1-3 a: abcde")

	if err != nil {
		t.Error(err.Error())
	}

	actual := policy.PasswordValidToboggan()
	if actual != expected {
		t.Errorf("expected %t got %t", expected, actual)
	}
}

func TestDay2Part2Example2(t *testing.T) {
	expected := false

	policy, err := daytwo.NewPolicy("1-3 b: cdefg")

	if err != nil {
		t.Error(err.Error())
	}

	actual := policy.PasswordValidToboggan()
	if actual != expected {
		t.Errorf("expected %t got %t", expected, actual)
	}
}

func TestDay2Part2Example3(t *testing.T) {
	expected := false

	policy, err := daytwo.NewPolicy("2-9 c: ccccccccc")

	if err != nil {
		t.Error(err.Error())
	}

	actual := policy.PasswordValidToboggan()
	if actual != expected {
		t.Errorf("expected %t got %t", expected, actual)
	}
}

func TestDay2Part2Example4(t *testing.T) {
	expected := 1

	policies, err := daytwo.NewPolicies([]string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	})

	if err != nil {
		t.Error(err.Error())
	}

	actual := policies.NumValidToboggan()
	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay2Part2Exercise(t *testing.T) {
	expected := 737

	policyStrings, err := utils.LoadLines("../inputs/day2.txt")
	if err != nil {
		t.Error(err.Error())
	}

	policies, err := daytwo.NewPolicies(policyStrings)
	if err != nil {
		t.Error(err.Error())
	}

	actual := policies.NumValidToboggan()

	if actual != expected {
		t.Errorf("expected: %d got: %d", expected, actual)
	}
}
