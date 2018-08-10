package problem87

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	n := len(s1)
	set := make([]int, 256)
	for i := 0; i < n; i++ {
		set[s1[i]]++
		set[s2[i]]--
	}

	for _, v := range set {
		if v != 0 {
			return false
		}
	}

	for i := 1; i < n; i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}

		if isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i]) {
			return true
		}
	}

	return false
}
