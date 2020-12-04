package dayfour

import (
	"regexp"
	"strings"
)

// PassportChecker is a struct for batch checking passports
type PassportChecker struct {
	validation map[string]*regexp.Regexp
}

// Passports is a list of parsed passports
type Passports []map[string]string

// NewPassports parses a new batch of passports
func NewPassports(input string) Passports {
	split := strings.Split(input, "\n\n")

	passports := make([]map[string]string, len(split))
	kvRe := regexp.MustCompile(" |\n")
	for i, p := range split {
		passports[i] = make(map[string]string)
		kvs := kvRe.Split(p, -1)
		for _, kv := range kvs {
			s := strings.Split(kv, ":")
			passports[i][s[0]] = s[1]
		}
	}

	return passports
}

// NewPassportChecker parses the batch file input and creates a passport checker
func NewPassportChecker() *PassportChecker {
	return &PassportChecker{
		validation: map[string]*regexp.Regexp{
			"byr": regexp.MustCompile("^(19[2-9][0-9]|200[0-2])$"),
			"iyr": regexp.MustCompile("^(201[0-9]|2020)$"),
			"eyr": regexp.MustCompile("^(202[0-9]|2030)$"),
			"hgt": regexp.MustCompile("^(1[5-8][0-9]cm|19[0-3]cm|59in|6[0-9]in|7[0-6]in)$"),
			"hcl": regexp.MustCompile("^(#[0-9a-f]{6})$"),
			"ecl": regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$"),
			"pid": regexp.MustCompile("^[0-9]{9}$"),
		},
	}
}

// NumValidPassports finds the number of valid passports in the batch
func (p *PassportChecker) NumValidPassports(passports Passports, validateFields bool) int {
	numValid := 0
	for _, passport := range passports {
		valid := true

		for key, validation := range p.validation {
			value, ok := passport[key]
			if !ok {
				valid = false
				break
			}

			valueValid := !validateFields || validation.MatchString(value)
			if !valueValid {
				valid = false
				break
			}
		}

		if valid {
			numValid++
		}
	}

	return numValid
}
