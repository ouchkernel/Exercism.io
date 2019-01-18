package strand

import (
	"strings"
)

func ToRNA(dna string) string {
	var rna []string
	var value string

	for i := 0; i < len(dna); i++ {
		switch string(dna[i]) {
		case "G":
			value = "C"
		case "C":
			value = "G"
		case "T":
			value = "A"
		case "A":
			value = "U"
		}
		rna = append(rna, value)
	}
	return strings.Join(rna, "")
}
