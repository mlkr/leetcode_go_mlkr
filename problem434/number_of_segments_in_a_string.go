package problem434

func countSegments(s string) int {
	count := 0
	isLetter := false
	for i:=0;i<len(s);i++{
		if s[i] != ' ' && !isLetter {
			count++
			isLetter = true
			continue
		}

		if s[i] == ' ' {
			isLetter = false
		}

	}


	return count
}