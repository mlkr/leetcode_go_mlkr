package problem389

func findTheDifference(s string, t string) byte {
	m := make(map[rune]int)
	for _, b := range s {
		m[b]++
	}

	var res byte
	for _, b := range t {
		if m[b] == 0 {
			res = byte(b)
			break
		}
		m[b]--
	}

	return res
}