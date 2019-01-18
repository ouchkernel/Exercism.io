package accumulate

func Accumulate(items []string, f func(string) string) []string {
	var results []string
	for _, v := range items {
		results = append(results, f(v))
	}
	return results
}
