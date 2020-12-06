package daysix

import (
	"math"
	"math/bits"
	"strings"
)

func LetterToPower(l byte) float64 {
	t := l - byte('a')
	return float64(t)
}

func AnswersToInt(answers string) uint {
	input := []byte(answers)
	var res float64 = 0

	for _, v := range input {
		pow := LetterToPower(v)
		tmp := math.Pow(2, pow)
		res += tmp
	}

	return uint(res)
}

func GroupAnswersToIntPart1(answers []uint) uint {
	var total uint = 0

	for _, v := range answers {
		total |= v
	}

	return total
}

func GroupAnswersToIntPart2(answers []uint) uint {
	var total uint = answers[0]

	for _, v := range answers {
		total &= v
	}

	return total
}

func CountYesAnswers(answers uint) int {
	return bits.OnesCount(answers)
}

func CountAllAnswers(answers []int) int {
	total := 0

	for _, a := range answers {
		total += a
	}

	return total
}

func ParseInput(input string) [][]string {
	groups := strings.Split(input, "\n\n")

	groupsAnswers := make([][]string, len(groups))
	for i, v := range groups {
		split := strings.Split(v, "\n")
		var ga []string
		for _, s := range split {
			if s != "" {
				ga = append(ga, s)
			}
		}
		groupsAnswers[i] = ga
	}

	return groupsAnswers
}

func CalculateResultPart1(groups [][]string) int {
	totals := make([]int, len(groups))

	for i, group := range groups {
		individuals := make([]uint, len(group))
		for j, individual := range group {
			individuals[j] = AnswersToInt(individual)
		}

		groupAnswer := GroupAnswersToIntPart1(individuals)

		totals[i] = CountYesAnswers(groupAnswer)
	}

	return CountAllAnswers(totals)
}

func CalculateResultPart2(groups [][]string) int {
	totals := make([]int, len(groups))

	for i, group := range groups {
		individuals := make([]uint, len(group))
		for j, individual := range group {
			individuals[j] = AnswersToInt(individual)
		}

		groupAnswer := GroupAnswersToIntPart2(individuals)

		totals[i] = CountYesAnswers(groupAnswer)
	}

	return CountAllAnswers(totals)
}

func Part1(input string) int {
	parsed := ParseInput(input)
	result := CalculateResultPart1(parsed)
	return result
}

func Part2(input string) int {
	parsed := ParseInput(input)
	result := CalculateResultPart2(parsed)
	return result
}
