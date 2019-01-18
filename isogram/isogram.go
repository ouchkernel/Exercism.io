package isogram

import (
	"strings"
)

func IsIsogram(s string) bool {

	str := strings.Replace(s, " ", "", -1)
	str = strings.Replace(str, "-", "", -1)
	for _, v := range str {
		count := strings.Count(strings.ToLower(str), string(v))
		if count > 1 {
			return false
		}
	}
	return true
}
