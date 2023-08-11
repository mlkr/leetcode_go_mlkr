package problem520

func detectCapitalUse(word string) bool {
	cnt := 0
	for _, v := range word {
		if v > 64 && v < 91 {
			cnt++
		}
	}

	if cnt == 0 || cnt == len(word) {
		return true
	}

	if cnt == 1 && (word[0] > 64 && word[0] < 91) {
		return true
	}

	return false
}
