package dayseven

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func sanitizeBagName(bag string) string {
	out := strings.TrimSpace(bag)
	removeBagsRe := regexp.MustCompile("[ ]*bag[s]*")
	out = removeBagsRe.ReplaceAllString(out, "")
	return out
}

func parseChild(child string) (string, int) {
	re := regexp.MustCompile("([0-9]+) ([a-z ]*) bag[s]*")
	match := re.FindStringSubmatch(child)

	mc := len(match)
	if mc < 3 {
		log.Printf("no match --%s %d", child, mc)
		return "", -1
	}

	num, err := strconv.Atoi(match[1])
	if err != nil {
		log.Fatal(err.Error())
	}
	c := match[2]

	return c, num
}

func sanitizeChildBags(bags string) map[string]int {
	outBags := make(map[string]int)
	children := strings.Split(bags, ",")
	for _, child := range children {
		colour, num := parseChild(child)
		if colour == "" && num == -1 {
			continue
		}
		outBags[colour] = num
	}

	return outBags
}

func parseRule(rule string) (string, map[string]int) {
	split := strings.Split(rule, "contain")
	parent := sanitizeBagName(split[0])
	children := sanitizeChildBags(split[1])
	return parent, children
}

func parseRules(rules []string) map[string]map[string]int {
	rMap := make(map[string]map[string]int)
	for _, r := range rules {
		p, c := parseRule(r)
		rMap[p] = c
	}

	return rMap
}

func contains(bags []string, bag string) bool {
	if bags == nil {
		return false
	}

	for _, b := range bags {
		if b == bag {
			return true
		}
	}

	return false
}

type Node struct {
	HasParent bool
	Colour    string
	Num       int
	Children  []Node
}

func NewTree(colour string, num int, hasParent bool, allBags map[string]map[string]int) Node {
	children, ok := allBags[colour]
	if !ok {
		log.Fatalf("bag %s not found in all bags\n", colour)
	}

	childNodes := make([]Node, len(children))
	i := 0
	for childColour, childNum := range children {
		childNodes[i] = NewTree(childColour, childNum, true, allBags)
		i++
	}

	return Node{
		Colour:    colour,
		Num:       num,
		HasParent: hasParent,
		Children:  childNodes,
	}
}

func ContainsBag(n Node, bag string) bool {
	if n.HasParent && n.Colour == bag {
		return true
	}

	for _, child := range n.Children {
		if ContainsBag(child, bag) {
			return true
		}
	}

	return false
}

func countBags(bag string, bags map[string]map[string]int) int {
	count := 0

	for parent := range bags {
		t := NewTree(parent, 0, false, bags)

		if ContainsBag(t, bag) {
			count++
		}
	}

	return count
}

func numBags(node Node) int {
	count := node.Num

	multiplier := node.Num
	if multiplier == 0 {
		multiplier = 1
	}

	for _, c := range node.Children {
		m := multiplier * numBags(c)
		count += m
	}

	return count
}

func numBagsInBag(bag string, bags map[string]map[string]int) int {
	t := NewTree(bag, 0, false, bags)
	c := numBags(t)
	return c
}

func Part1(rules []string) int {
	rulesMap := parseRules(rules)
	cb := countBags("shiny gold", rulesMap)
	return cb
}

func Part2(rules []string) int {
	rulesMap := parseRules(rules)
	cb := numBagsInBag("shiny gold", rulesMap)
	return cb
}
