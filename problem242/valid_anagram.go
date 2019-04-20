package problem242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := [26]int{}
	for i := range s {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}

	for i := range m {
		if m[i] != 0 {
			return false
		}
	}

	return true
}
