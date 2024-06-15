package problem567

func checkInclusion(s1 string, s2 string) bool {
	size1 := len(s1)
	size2 := len(s2)

	if size1 > size2 {
		return false
	}

	var n1 [26]int
	var n2 [26]int
	for i := 0; i < size1; i++ {
		n1[s1[i]-'a']++
		n2[s2[i]-'a']++
	}

	if n1 == n2 {
		return true
	}

	for i := size1; i < size2; i++ {
		n2[s2[i-size1]-'a']--
		n2[s2[i]-'a']++

		if n1 == n2 {
			return true
		}
	}

	return false
}
