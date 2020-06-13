package problem392

func isSubsequence(s string, t string) bool {
	size := len(s)
	if size == 0 {
		return true
	}

	i := 0
	for j := range t {
		if s[i] == t[j] {
			i++
			if i == size {
				return true
			}
		}
	}

	return false
}
