package daythirteen

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type BusSchedule struct {
	minLeavingTime uint64
	busIds         []uint64
	largestBusId   uint64
}

func NewBusSchedule(input []string) *BusSchedule {
	minLeavingTime, err := strconv.Atoi(strings.TrimSpace(input[0]))
	if err != nil {
		log.Fatal(err)
	}

	var busIds []uint64

	for _, bid := range strings.Split(input[1], ",") {
		sanitizedBID := strings.TrimSpace(bid)
		if sanitizedBID == "x" {
			sanitizedBID = "0"
		}

		busId, err := strconv.Atoi(sanitizedBID)
		if err != nil {
			log.Fatal(err)
		}

		busIds = append(busIds, uint64(busId))
	}

	largestBusId := busIds[0]
	for _, v := range busIds {
		if largestBusId < v {
			largestBusId = v
		}
	}

	fmt.Printf("Min Leaving Time: %d Largest BusId: %d Bus Ids: %v\n", minLeavingTime, largestBusId, busIds)

	return &BusSchedule{
		minLeavingTime: uint64(minLeavingTime),
		busIds:         busIds,
	}
}

func (b *BusSchedule) SetMinLeavingTime(time uint64) {
	b.minLeavingTime = time
}

func (b *BusSchedule) Part1() uint64 {
	currTime := b.minLeavingTime
	for {
		for _, busId := range b.busIds {
			if busId == 0 {
				continue
			}

			isBusLeavingTime := (currTime % busId) == 0

			if isBusLeavingTime {
				wait := currTime - b.minLeavingTime

				return wait * busId
			}
		}

		currTime++
	}
}

/*
The idea for this solution was not by me but I managed to figure out the code.

The idea of the question is that you loop through bus ids and find where all busIds leave 1 min after each other.
Including the canceled busses (x or 0).

e.g. the leaving schedule is

t0 bus 0 leaves, t1 bus 1 leaves, etc...

a bus leaves when t mod busid is a clean division i.e. 0

so finding the sequence is as simple as finding when t % busid == 0 is true for all buses.

This is fine and the algorithm solves it really quick for all the examples. However with a longer amount of bus ids this takes such a long time I could not get an answer. You have to find a way of skipping over time periods.

Well first I figured out that once we find the first bus leaving. We can skip time by the first buses id. This is still too slow on the real input.

What was mentioned on reddit a lot was
https://en.wikipedia.org/wiki/Chinese_remainder_theorem#Search_by_sieving

This is applied to the problem using Least Common Multiple LCM

e.g. for bus ids {67,7,x,59,61}

you find the first repition e.g 67 (like we did before) and start incrementing by that.

Then when you find the first appearence of bus id 1 (7) incrementing you then do LCM(67, 7)
This is the next step where they share a divisor in common i.e. when both t % busId == 0.

This means you can skip that amount of time knowing the sequence won't appear until at least then.

This keeps occuring for the rest of the bus ids except the last.

LCM(67,7,59)

The reason we don't need the last is once we found that we're done anyway and LCM(67,7,59,61) will find the second occurence.

With that we just have to move the last bit until we find out goal
*/
func (b *BusSchedule) Part2() uint64 {
	currTime := b.minLeavingTime

	var incrementAmount uint64 = 1
	foundBusIndex := -1

	for {
		foundSequence := true

		for i, bid := range b.busIds {
			if bid == 0 {
				continue
			}

			t := currTime + uint64(i)

			busLeaving := (t % bid) == 0

			if busLeaving {
				if foundBusIndex < i {
					foundBusIndex = i

					previousBusIds := removeZeros(b.busIds[:foundBusIndex+1])

					lcm := previousBusIds[0]
					if len(previousBusIds) > 1 {
						lcm = LCM(previousBusIds[0], previousBusIds[1], previousBusIds[2:]...)
					}

					incrementAmount = lcm
				}
			} else {
				foundSequence = false
				break
			}
		}

		if foundSequence {
			return currTime
		}

		currTime += incrementAmount
	}
}

func removeZeros(ints []uint64) []uint64 {
	var output []uint64
	for _, o := range ints {
		if o != 0 {
			output = append(output, o)
		}
	}

	return output
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b uint64, integers ...uint64) uint64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
