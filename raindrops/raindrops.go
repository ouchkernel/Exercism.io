package raindrops

import "strconv"

// Convert converts a number to raindrop sounds
func Convert(n int) string {
	var speak string

	// check for factors 3,5,7
	if n%3 == 0 {
		speak += "Pling"
	}
	if n%5 == 0 {
		speak += "Plang"
	}
	if n%7 == 0 {
		speak += "Plong"
	}

	if speak == "" {
		return strconv.Itoa(n)
	}

	return speak
}
