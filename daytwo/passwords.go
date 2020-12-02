package daytwo

import (
	"strconv"
	"strings"
)

// Policy is a password policy with a password that may or may not be valid
type Policy struct {
	Min      int
	Max      int
	Char     rune
	Password string
}

// NewPolicy creates a new policy based on a row of input
func NewPolicy(row string) (*Policy, error) {
	policy := Policy{}
	parts := strings.Split(row, " ")

	minMax := strings.Split(parts[0], "-")

	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		return nil, err
	}
	policy.Min = min

	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		return nil, err
	}
	policy.Max = max

	char := parts[1][0]
	policy.Char = rune(char)

	policy.Password = parts[2]

	return &policy, nil
}

// PasswordValidSledRental checks if the password is valid against the current policy
func (p *Policy) PasswordValidSledRental() bool {
	occurences := 0

	for _, c := range p.Password {
		if p.Char == c {
			occurences++
		}
	}

	valid := p.Min <= occurences && occurences <= p.Max

	return valid
}

// PasswordValidToboggan checks if the password is valid against the Toboggan Policy
func (p *Policy) PasswordValidToboggan() bool {
	pos1 := p.Min - 1
	pos2 := p.Max - 1

	match1 := rune(p.Password[pos1]) == p.Char
	match2 := rune(p.Password[pos2]) == p.Char

	return match1 != match2
}

// Policies is a slice of policy
type Policies []Policy

// NewPolicies creates a new Policy Slice
func NewPolicies(rows []string) (Policies, error) {
	policies := make([]Policy, len(rows))

	for i, r := range rows {
		p, err := NewPolicy(r)
		if err != nil {
			return nil, err
		}
		policies[i] = *p
	}

	return policies, nil
}

// NumValidSledRental counts the number of valid policies
func (p Policies) NumValidSledRental() int {
	numValid := 0
	for _, r := range p {
		if r.PasswordValidSledRental() {
			numValid++
		}
	}

	return numValid
}

// NumValidToboggan counts the number of valid policies
func (p Policies) NumValidToboggan() int {
	numValid := 0
	for _, r := range p {
		if r.PasswordValidToboggan() {
			numValid++
		}
	}

	return numValid
}
