package romannumerals

import (
	"errors"
)

func ToRomanNumeral(s int) (string, error) {
	if s <= 0 || s > 3000 {
		return "", errors.New("out of range")
	}

	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	M := []string{"", "M", "MM", "MMM"}
	answer := M[s/1000] + C[(s%1000)/100] + X[(s%100)/10] + I[s%10]

	return answer, nil
}
