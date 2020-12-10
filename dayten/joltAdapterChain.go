package dayten

import (
	"fmt"
	"sort"
)

type JoltAdapterChain struct {
	adapters     []int
	adapterDiffs map[int]int
}

func NewJoltAdapterChain(input []int) *JoltAdapterChain {
	adapters := append([]int{0}, input...)
	sort.Ints(adapters)
	deviceAdapter := adapters[len(adapters)-1] + 3
	adapters = append(adapters, deviceAdapter)

	return &JoltAdapterChain{
		adapters:     adapters,
		adapterDiffs: make(map[int]int),
	}
}

func (j *JoltAdapterChain) DeviceAdapter() int {
	return j.adapters[len(j.adapters)-1]
}

func (j *JoltAdapterChain) Connect() {
	for i := 1; i < len(j.adapters); i++ {
		a := j.adapters[i-1]
		b := j.adapters[i]

		diff := b - a
		j.UpJoltDiffCount(diff)
	}
}

func (j *JoltAdapterChain) NumAdaptersOfDiff(diff int) int {
	diff, ok := j.adapterDiffs[diff]
	if !ok {
		return 0
	}

	return diff
}

func (j *JoltAdapterChain) UpJoltDiffCount(diff int) {
	_, ok := j.adapterDiffs[diff]
	if !ok {
		j.adapterDiffs[diff] = 0
	}

	j.adapterDiffs[diff]++
}

func (j *JoltAdapterChain) Part1() int {
	j.Connect()
	one := j.NumAdaptersOfDiff(1)
	three := j.NumAdaptersOfDiff(3)

	return one * three
}

func (j *JoltAdapterChain) PossibleConnections(index int) []int {
	a := j.adapters[index]

	var connections []int
	for i := index + 1; i < len(j.adapters); i++ {
		b := j.adapters[i]
		diff := b - a

		if diff <= 3 {
			connections = append(connections, i)
		} else {
			break
		}
	}

	return connections
}

type Node struct {
	Index    int
	Children []Node
}

func (j *JoltAdapterChain) NewNode(index int) Node {
	children := j.PossibleConnections(index)

	childNodes := make([]Node, len(children))
	for i, v := range children {
		childNodes[i] = j.NewNode(v)
	}

	return Node{
		Index:    index,
		Children: childNodes,
	}
}

func (j *JoltAdapterChain) CountLeafNodes(node Node) int {
	if len(node.Children) == 0 {
		return 1
	}

	count := 0
	for _, c := range node.Children {
		count += j.CountLeafNodes(c)
	}

	return count
}

// Unique Permutations
func (j *JoltAdapterChain) Part2() int {
	tree := j.NewNode(0)
	leaves := j.CountLeafNodes(tree)
	return leaves
}

func GetOrDefault(key int, d int, m map[int]int) int {
	v, ok := m[key]
	if !ok {
		return d
	}
	return v
}

/*
Once again I missed out on math it seems had to do this as the above attempt does not work on the day's input

e.g. 0 1 4 5 6 7 10 11 12 15 16 19 22

[01] 00: 1 -1: 0 -2: 0 = 1
[04] 03: 0 02: 0 01: 1 = 1
[05] 04: 1 03: 0 02: 0 = 1
[06] 05: 1 04: 1 03: 0 = 2
[07] 06: 2 05: 1 04: 1 = 4
[10] 09: 0 08: 0 07: 4 = 4
[11] 10: 4 09: 0 08: 0 = 4
[12] 11: 4 10: 4 09: 0 = 8
[15] 14: 0 13: 0 12: 8 = 8
[16] 15: 8 14: 0 13: 0 = 8
[19] 18: 0 17: 0 16: 8 = 8
[22] 21: 0 20: 0 19: 8 = 8

*/
func (j *JoltAdapterChain) Part2Attempt2(debug bool) int {
	m := make(map[int]int)
	m[0] = 1
	for i := 1; i < len(j.adapters); i++ {
		a := j.adapters[i]
		one := GetOrDefault(a-1, 0, m)
		two := GetOrDefault(a-2, 0, m)
		three := GetOrDefault(a-3, 0, m)

		m[a] = one + two + three

		if debug {
			fmt.Printf("[%d] %d: %d %d: %d %d: %d = %d\n", a, a-1, one, a-2, two, a-3, three, m[a])
		}
	}

	res := m[j.DeviceAdapter()]

	return res
}
