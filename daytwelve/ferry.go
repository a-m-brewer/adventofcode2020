package daytwelve

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

const (
	NORTH   = "N"
	EAST    = "E"
	SOUTH   = "S"
	WEST    = "W"
	LEFT    = "L"
	RIGHT   = "R"
	FORWARD = "F"
	PART1   = "PART1"
	PART2   = "PART2"
)

var directions []string = []string{NORTH, EAST, SOUTH, WEST}

func DirectionToDegrees(dir string) int {
	switch dir {
	case NORTH:
		return 0
	case EAST:
		return 90
	case SOUTH:
		return 180
	case WEST:
		return 270
	default:
		log.Fatalf("unknown direction: %s\n", dir)
	}

	return -1
}

func DegreesToDir(degrees int) string {
	directionDegrees := 360 + degrees
	directionIndex := int(math.Abs(float64((directionDegrees / 90) % 4)))
	return directions[directionIndex]
}

func DirectionToXY(dir string) (int, int) {
	switch dir {
	case NORTH:
		return 0, -1
	case EAST:
		return 1, 0
	case SOUTH:
		return 0, 1
	case WEST:
		return -1, 0
	default:
		log.Fatalf("unknown direction: %s\n", dir)
	}

	return -9999, -9999
}

type ShipInstruction struct {
	Op  string
	Arg int
}

func ParseInstruction(input string) ShipInstruction {
	op := string(input[0])
	arg, err := strconv.Atoi(input[1:])

	if err != nil {
		log.Fatal(err)
	}

	return ShipInstruction{
		Op:  op,
		Arg: arg,
	}
}

type Ship struct {
	x                  int
	y                  int
	waypointX          int
	waypointY          int
	currentDirection   string
	currentInstruction int
	instructions       []ShipInstruction
	Part               string
}

func NewShip(input []string) *Ship {
	instructions := make([]ShipInstruction, len(input))
	for i, v := range input {
		instructions[i] = ParseInstruction(v)
	}

	return &Ship{
		x:                  0,
		y:                  0,
		waypointX:          10,
		waypointY:          -1,
		currentDirection:   EAST,
		currentInstruction: 0,
		instructions:       instructions,
		Part:               PART1,
	}
}

func (s *Ship) Step() {
	currInstruction := s.instructions[s.currentInstruction]

	switch currInstruction.Op {
	case NORTH, EAST, SOUTH, WEST:
		xDelta, yDelta := DirectionToXY(currInstruction.Op)
		xDelta = (currInstruction.Arg * xDelta)
		yDelta = (currInstruction.Arg * yDelta)
		s.Move(xDelta, yDelta)
		break
	case LEFT, RIGHT:
		s.Rotate()
		break
	case FORWARD:
		s.Forward()
		break
	}

	fmt.Printf("op: %s arg: %d\nx: %d y: %d direction: %s\n\n", currInstruction.Op, currInstruction.Arg, s.x, s.y, s.currentDirection)

	s.currentInstruction++
}

func (s *Ship) Move(deltaX, deltaY int) {
	switch s.Part {
	case PART1:
		s.x += deltaX
		s.y += deltaY
		break
	case PART2:
		s.waypointX += deltaX
		s.waypointY += deltaY
		break
	default:
		log.Fatalf("unknown part: %s", s.Part)
	}
}

func (s *Ship) Rotate() {
	currentInstruction := s.instructions[s.currentInstruction]

	currDirInDegrees := DirectionToDegrees(s.currentDirection)
	changeDegrees := currentInstruction.Arg
	if currentInstruction.Op == LEFT {
		changeDegrees *= -1
	}

	nextDegrees := currDirInDegrees + changeDegrees
	nextDirection := DegreesToDir(nextDegrees)

	fmt.Printf("currDir: %d change: %d next: %d\n", currDirInDegrees, changeDegrees, nextDegrees)

	s.currentDirection = nextDirection

	fmt.Printf("%d\n", changeDegrees)
	if s.Part == PART2 {
		newX := s.waypointX
		newY := s.waypointY

		switch changeDegrees {
		case 90, -270:
			newX = s.waypointY * -1
			newY = s.waypointX
			break
		case 180, -180:
			newX *= -1
			newY *= -1
		case 270, -90:
			newX = s.waypointY
			newY = s.waypointX * -1
		}

		fmt.Printf("prevX: %d prevY: %d nextX: %d nextY: %d\n", s.waypointX, s.waypointY, newX, newY)

		s.waypointX = newX
		s.waypointY = newY
	}
}

func (s *Ship) Forward() {
	currentInstruction := s.instructions[s.currentInstruction]
	switch s.Part {
	case PART1:
		xDelta, yDelta := DirectionToXY(s.currentDirection)
		xDelta = (currentInstruction.Arg * xDelta)
		yDelta = (currentInstruction.Arg * yDelta)
		s.x += xDelta
		s.y += yDelta
		break
	case PART2:
		s.x += s.waypointX * currentInstruction.Arg
		s.y += s.waypointY * currentInstruction.Arg
		break
	default:
		log.Fatalf("unknown part: %s", s.Part)
	}
}

func (s *Ship) Run() {
	for i := 0; i < len(s.instructions); i++ {
		s.Step()
	}
}

func (s *Ship) DistanceFromStart() int {
	x := int(math.Abs(float64(s.x)))
	y := int(math.Abs(float64(s.y)))

	return x + y
}
