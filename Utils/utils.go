package Utils

func SplitIntoLines(s string) []string {
	var results []string
	n := len(s)
	i := 0

	for i < n {
		if s[i] == '\n' {
			results = append(results, "\n")
		} else {
			j := i + 1
			for j < n && s[j] != '\n' {
				j++
			}
			if j < n && s[j] == '\n' {
				results = append(results, s[i:j+1])
				i = j + 1
			} else {
				results = append(results, s[i:])
				break
			}
		}
	}
	return results
}
