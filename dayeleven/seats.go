package dayeleven

import (
	"fmt"
	"log"
)

const (
	FLOOR         = rune('.')
	EMPTY_SEAT    = rune('L')
	OCCUPIED_SEAT = rune('#')
	PART1         = "PART1"
	PART2         = "PART2"
)

type Seating struct {
	width                  int
	height                 int
	seats                  [][]rune
	ruleset                string
	maxOccupationTolerance int
	Log                    func(format string, args ...interface{})
}

func l(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func NewSeating(input [][]rune, ruleset string) *Seating {
	mot := 0
	switch ruleset {
	case PART1:
		mot = 4
	case PART2:
		mot = 5
	default:
		log.Fatalf("Unknown Ruleset: %s", ruleset)
	}

	return &Seating{
		height:                 len(input),
		width:                  len(input[0]),
		seats:                  input,
		ruleset:                ruleset,
		maxOccupationTolerance: mot,
		Log:                    l,
	}
}

func (s *Seating) Run() {
	s.PrintSeating()
	for s.Step() {
		// s.Log("step")
		s.PrintSeating()
	}
}

func (s *Seating) Step() bool {
	nextSeatingArangement := s.copyCurrentSeating()

	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			occupied := s.OccupiedNeighourCount(x, y)

			if s.seats[y][x] == EMPTY_SEAT && occupied == 0 {
				nextSeatingArangement[y][x] = OCCUPIED_SEAT
			}

			if s.seats[y][x] == OCCUPIED_SEAT && occupied >= s.maxOccupationTolerance {
				nextSeatingArangement[y][x] = EMPTY_SEAT
			}
		}
	}

	if s.SeatingUnchanged(nextSeatingArangement) {
		return false
	}

	s.seats = nextSeatingArangement

	return true
}

func (s *Seating) SeatingUnchanged(nextSeatingArangement [][]rune) bool {
	for y, row := range s.seats {
		for x, current := range row {
			next := nextSeatingArangement[y][x]

			if next != current {
				return false
			}
		}
	}

	return true
}

func (s *Seating) OccupiedNeighourCount(x, y int) int {
	switch s.ruleset {
	case PART1:
		return s.occupiedNeighourCountPart1(x, y)
	case PART2:
		return s.occupiedNeighourCountPart2(x, y)
	default:
		log.Fatalf("Unknown Ruleset: %s", s.ruleset)
	}
	return -1
}

func (s *Seating) occupiedNeighourCountPart1(x, y int) int {
	numOccupiedNeighbours := 0

	for yn := y - 1; yn <= y+1; yn++ {
		for xn := x - 1; xn <= x+1; xn++ {
			if xn == x && yn == y {
				// current seat
				continue
			}

			if xn < 0 || s.width <= xn || yn < 0 || s.height <= yn {
				// out of bounds
				continue
			}

			if s.seats[yn][xn] == OCCUPIED_SEAT {
				numOccupiedNeighbours++
			}
		}
	}

	return numOccupiedNeighbours
}

func (s *Seating) occupiedNeighourCountPart2(x, y int) int {
	occupiedCount := 0

	for yDelta := -1; yDelta <= 1; yDelta++ {
		for xDelta := -1; xDelta <= 1; xDelta++ {
			if yDelta == 0 && xDelta == 0 {
				continue
			}

			if s.occupiedNeighbourInDirection(x, y, xDelta, yDelta) {
				occupiedCount++
			}
		}
	}

	return occupiedCount
}

func (s *Seating) occupiedNeighbourInDirection(startX, startY, deltaX, deltaY int) bool {
	currX := startX
	currY := startY

	for {
		currX += deltaX
		currY += deltaY

		if currX < 0 || s.width <= currX || currY < 0 || s.height <= currY {
			break
		}

		tile := s.seats[currY][currX]

		if tile == EMPTY_SEAT {
			return false
		}

		if tile == OCCUPIED_SEAT {
			return true
		}
	}

	return false
}

func (s *Seating) copyCurrentSeating() [][]rune {
	duplicate := make([][]rune, len(s.seats))
	for i := range s.seats {
		duplicate[i] = make([]rune, len(s.seats[i]))
		copy(duplicate[i], s.seats[i])
	}

	return duplicate
}

func (s *Seating) NumOccupied() int {
	count := 0

	for _, row := range s.seats {
		for _, seat := range row {
			if seat == OCCUPIED_SEAT {
				count++
			}
		}
	}

	return count
}

func (s *Seating) PrintSeating() {
	fmtStr := "\n"
	seats := make([]interface{}, s.width*s.height)
	i := 0
	for _, row := range s.seats {
		for _, seat := range row {
			fmtStr += "%s"
			seats[i] = string(seat)
			i++
		}
		fmtStr += "\n"
	}
	fmtStr += "\n"

	s.Log(fmtStr, seats...)
}
