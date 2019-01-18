package hamming

import (
	"errors"
)

// Distance compare distance between two dna strands of equal lengths
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("squences are different lengths")
	}

	var distance int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
