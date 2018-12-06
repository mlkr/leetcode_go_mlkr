package problem187

func findRepeatedDnaSequences(s string) []string {
	size := len(s)
	if size < 10 {
		return nil
	}

	res := make([]string, 0, size-9)
	m := make(map[string]int, size-9)
	for i := 0; i+9 < size; i++ {
		str := s[i : i+10]
		m[str]++

		if m[str] == 2 {
			res = append(res, str)
		}
	}

	return res
}
