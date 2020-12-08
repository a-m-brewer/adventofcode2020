package dayeight

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

const NOP = "nop"
const ACC = "acc"
const JMP = "jmp"

type Instruction struct {
	Op  string
	Arg int
}

type Computer struct {
	acc          int
	ip           int
	instructions []Instruction
	program      []string
	seen         []int
	changedIps   []int
}

func parseProgram(program []string) []Instruction {
	instructions := make([]Instruction, len(program))
	for i, p := range program {
		split := strings.Split(strings.TrimSpace(p), " ")
		op := strings.TrimSpace(split[0])
		arg, err := strconv.Atoi(strings.TrimSpace(split[1]))
		if err != nil {
			log.Fatal(err)
		}

		instructions[i] = Instruction{
			Op:  op,
			Arg: arg,
		}
	}

	return instructions
}

func NewComputer(program []string) *Computer {
	return &Computer{
		acc:          0,
		ip:           0,
		instructions: parseProgram(program),
		seen:         make([]int, 0),
		changedIps:   make([]int, 0),
		program:      program,
	}
}

func (c *Computer) Reset() {
	c.ip = 0
	c.acc = 0
	c.seen = make([]int, 0)
	c.instructions = parseProgram(c.program)
}

func (c *Computer) Acc() int {
	return c.acc
}

func (c *Computer) InstructionPointer() int {
	return c.ip
}

func (c *Computer) CurrInstruction() Instruction {
	return c.instructions[c.ip]
}

func (c *Computer) FirstIpOfTypes(types []string) (int, error) {
	for _, ip := range c.seen {
		skipIp := false
		for _, cip := range c.changedIps {
			if ip == cip {
				skipIp = true
			}
		}

		if skipIp {
			continue
		}

		instruction := c.instructions[ip]

		for _, opType := range types {
			if opType == instruction.Op {
				return ip, nil
			}
		}
	}

	return -1, errors.New("could not find instruction of types")
}

func (c *Computer) DetectedInfiniteLoop() bool {
	for _, ins := range c.seen {
		if c.ip == ins {
			return true
		}
	}

	return false
}

func (c *Computer) ShouldHalt() bool {
	return len(c.instructions) <= c.ip
}

func (c *Computer) Step() {
	currInstruction := c.CurrInstruction()

	c.seen = append(c.seen, c.ip)

	switch currInstruction.Op {
	case ACC:
		c.acc += currInstruction.Arg
		c.ip++
		break
	case JMP:
		c.ip += currInstruction.Arg
		break
	case NOP:
		c.ip++
		break
	default:
		log.Fatalf("unknown op: %s arg: %d\n", currInstruction.Op, currInstruction.Arg)
	}
}

func (c *Computer) Run() int {
	for {
		if c.ShouldHalt() {
			return 0
		}

		if c.DetectedInfiniteLoop() {
			return 1
		}

		c.Step()
	}
}

func (c *Computer) ErrorCorrectionTool() int {
	types := []string{JMP, NOP}
	for {
		statusCode := c.Run()
		if statusCode == 0 {
			break
		}

		ip, err := c.FirstIpOfTypes(types)
		if err != nil {
			log.Fatal(err)
		}

		newOp := ""
		if c.instructions[ip].Op == JMP {
			newOp = NOP
		} else if c.instructions[ip].Op == NOP {
			newOp = JMP
		} else {
			log.Fatalf("log to change is not NOP or JMP but: %s", c.instructions[ip].Op)
		}

		c.changedIps = append(c.changedIps, ip)

		c.Reset()

		c.instructions[ip].Op = newOp
	}

	return c.changedIps[len(c.changedIps)-1]
}
