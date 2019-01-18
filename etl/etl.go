package etl

import "strings"

func Transform(legacy map[int][]string) map[string]int {
	newSys := map[string]int{}

	for i, v := range legacy {
		for _, val := range v {
			newSys[strings.ToLower(val)] = i
		}
	}

	return newSys
}
