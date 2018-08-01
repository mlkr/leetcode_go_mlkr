package problem58

func lengthOfLastWord(s string) int {
	sLen := len(s)
	if sLen == 0 {
		return 0
	}

	res := 0
	for i := sLen - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res != 0 {
				return res
			}
			continue
		}

		res++
	}

	return res
}
