package collatzconjecture

import (
	"errors"
)

// CollatzConjecture see https://en.wikipedia.org/wiki/3x_%2B_1_problem
func CollatzConjecture(num int) (int, error) {
	var steps int

	if num < 0 {
		return 0, errors.New("negitive number found")
	}

	for {
		if num == 1 {
			break
		}
		if num%2 == 0 {
			num = num / 2
			steps++
		} else if num%2 != 0 {
			num = (num * 3) + 1
			steps++
		}
	}
	return steps, nil
}
