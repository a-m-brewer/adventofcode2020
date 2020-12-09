package daynine

import (
	"errors"
	"fmt"
	"sort"
)

type XMASDecrypter struct {
	input    []int
	preamble int
}

func NewXMASDecrypter(input []int, preamble int) *XMASDecrypter {
	return &XMASDecrypter{
		input:    input,
		preamble: preamble,
	}
}

func (x *XMASDecrypter) ContainsInvalidNumber() (bool, int) {
	for i := x.preamble; i < len(x.input); i++ {
		if !x.SumInPreamble(i) {
			return true, x.input[i]
		}
	}

	return false, -1
}

func (x *XMASDecrypter) SumInPreamble(currIndex int) bool {
	sum := x.input[currIndex]
	low := currIndex - x.preamble

	for i := low; i < currIndex; i++ {
		for j := low; j < currIndex; j++ {
			a := x.input[i]
			b := x.input[j]

			if a == b {
				continue
			}

			if a+b == sum {
				return true
			}
		}
	}

	return false
}

func (x *XMASDecrypter) FindEncryptionWeakness() (int, error) {
	foundInvalidNumber, number := x.ContainsInvalidNumber()
	if !foundInvalidNumber {
		return -1, errors.New("found no invalid key for weakness exploit")
	}

	subset := x.findInLine(number)
	if subset == nil {
		return -1, errors.New(fmt.Sprintf("found no subset that matches weakness num: %d", number))
	}

	min, max := minMax(subset)

	return min + max, nil
}

func (x *XMASDecrypter) findInLine(target int) []int {
	subsetSize := 2

	for subsetSize <= len(x.input) {

		for i := subsetSize - 1; i < len(x.input); i++ {
			point := i + 1
			subset := x.input[point-subsetSize : point]

			sum := sumSlice(subset)

			if sum == target {
				return subset
			}
		}

		subsetSize++
	}

	return nil
}

func sumSlice(slice []int) int {
	s := 0

	for _, i := range slice {
		s += i
	}

	return s
}

func minMax(slice []int) (int, int) {
	sort.Ints(slice)

	min := slice[0]
	max := slice[len(slice)-1]

	return min, max
}
