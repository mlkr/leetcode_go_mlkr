package problem205

func isIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)

	size := len(s)
	if size != len(t) {
		return false
	}

	for i := 0; i < size; i++ {
		tb, ok1 := s2t[s[i]]
		sb, ok2 := t2s[t[i]]

		if ok1 {
			if tb != t[i] {
				return false
			}
		}

		if ok2 {
			if sb != s[i] {
				return false
			}
		}

		s2t[s[i]] = t[i]
		t2s[t[i]] = s[i]
	}

	return true
}

// 解法二(最佳)
func isIsomorphic2(s string, t string) bool {
	size := len(s)
	if size != len(t) {
		return false
	}

	m1 := make([]int, 256)
	m2 := make([]int, 256)

	for i := 0; i < size; i++ {
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}

		m1[int(s[i])] = i + 1
		m2[int(t[i])] = i + 1
	}

	return true
}
