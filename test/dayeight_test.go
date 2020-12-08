package test

import (
	"adventofcode2020/dayeight"
	"adventofcode2020/utils"
	"testing"
)

func TestProcessesInstructionsCorrectly(t *testing.T) {
	program, err := utils.LoadLines("../inputs/day8example1.txt")
	if err != nil {
		t.Error(err)
	}

	expected := []dayeight.Instruction{
		{Op: dayeight.NOP, Arg: 0},
		{Op: dayeight.ACC, Arg: 1},
		{Op: dayeight.JMP, Arg: 4},
		{Op: dayeight.ACC, Arg: 1},
		{Op: dayeight.JMP, Arg: -4},
		{Op: dayeight.ACC, Arg: 3},
		{Op: dayeight.JMP, Arg: -3},
	}

	comp := dayeight.NewComputer(program)

	for i, eInstruction := range expected {
		cInstruction := comp.CurrInstruction()

		if cInstruction.Op != eInstruction.Op {
			t.Errorf("%d op expected %s got %s\n", i, eInstruction.Op, cInstruction.Op)
		}

		if cInstruction.Arg != eInstruction.Arg {
			t.Errorf("%d arg expected %d got %d\n", i, eInstruction.Arg, cInstruction.Arg)
		}

		comp.Step()
	}
}

func TestCanDetectInfiniteLoop(t *testing.T) {
	program, err := utils.LoadLines("../inputs/day8example1.txt")
	if err != nil {
		t.Error(err)
	}

	comp := dayeight.NewComputer(program)

	expected := 5

	comp.Run()

	actual := comp.Acc()

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay8Part1Excercise(t *testing.T) {
	program, err := utils.LoadLines("../inputs/day8.txt")
	if err != nil {
		t.Error(err)
	}

	comp := dayeight.NewComputer(program)

	expected := 1614

	comp.Run()

	actual := comp.Acc()

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay8Part2Example1(t *testing.T) {
	program, err := utils.LoadLines("../inputs/day8example2.txt")
	if err != nil {
		t.Error(err)
	}

	comp := dayeight.NewComputer(program)

	expected := 8

	comp.Run()

	actual := comp.Acc()

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestDay8Part2Excercise(t *testing.T) {
	program, err := utils.LoadLines("../inputs/day8.txt")
	if err != nil {
		t.Error(err)
	}

	comp := dayeight.NewComputer(program)

	expected := 1260

	comp.ErrorCorrectionTool()

	actual := comp.Acc()

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}
