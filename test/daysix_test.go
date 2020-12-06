package test

import (
	"adventofcode2020/daysix"
	"adventofcode2020/utils"
	"testing"
)

func TestConvertLetterToPower_a(t *testing.T) {
	var expected float64 = 0

	actual := daysix.LetterToPower('a')

	if actual != expected {
		t.Errorf("expected %f got %f", expected, actual)
	}
}

func TestCanConvertQuestionnaireToNumber(t *testing.T) {
	var expected uint = 1
	input := "a"

	actual := daysix.AnswersToInt(input)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestCanConvertQuestionnaireToNumber2(t *testing.T) {
	var expected uint = 67108863
	input := "abcdefghijklmnopqrstuvwxyz"

	actual := daysix.AnswersToInt(input)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestCanConvertQuestionnaireToNumber3(t *testing.T) {
	var expected uint = 0
	input := ""

	actual := daysix.AnswersToInt(input)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestCanConvertGroupAnswersToUInt(t *testing.T) {
	var expected uint = 3

	input := []uint{1, 2, 3}

	actual := daysix.GroupAnswersToIntPart1(input)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestCanCountYesAnswers(t *testing.T) {
	expected := 2

	var input uint = 6

	actual := daysix.CountYesAnswers(input)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay6Part1Example1(t *testing.T) {
	input, err := utils.LoadString("../inputs/day6example.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 11

	actual := daysix.Part1(input)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay6Part1Exercise(t *testing.T) {
	input, err := utils.LoadString("../inputs/day6.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 6782

	actual := daysix.Part1(input)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay6Part2Example1(t *testing.T) {
	input, err := utils.LoadString("../inputs/day6example.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 6

	actual := daysix.Part2(input)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay6Part2Exercise(t *testing.T) {
	input, err := utils.LoadString("../inputs/day6.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 3596

	actual := daysix.Part2(input)

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}
