package dayone

import "errors"

// CalculateExpenses finds the two entries that sum to 2020 and multiplies them
func CalculateExpenses(expenses []int) (int, error) {
	return bruteForce(expenses)
}

func bruteForce(expenses []int) (int, error) {
	for _, x := range expenses {
		for _, y := range expenses {
			if x+y == 2020 {
				return x * y, nil
			}
		}
	}

	return -1, errors.New("found no two numbers that add to 2020")
}

// CalculateExpenses2 finds the two entries that sum to 2020 and multiplies them
func CalculateExpenses2(expenses []int) (int, error) {
	return bruteForce2(expenses)
}

func bruteForce2(expenses []int) (int, error) {
	for _, x := range expenses {
		for _, y := range expenses {
			for _, z := range expenses {
				if x+y+z == 2020 {
					return x * y * z, nil
				}
			}
		}
	}

	return -1, errors.New("found no two numbers that add to 2020")
}
