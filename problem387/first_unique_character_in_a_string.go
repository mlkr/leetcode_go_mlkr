package problem387

func firstUniqChar(s string) int {
	size := len(s)
	if size == 0 {
		return -1
	}

	idxs := make([]int, 26)
	for i := 0; i < size; i++ {
		key := s[i] - 'a'
		if idxs[key] == 0 {
			idxs[key] = i + 1
		} else {
			idxs[key] = -1
		}
	}

	res := size + 1
	for _, idx := range idxs {
		if idx > 0 && idx < res {
			res = idx
		}
	}

	if res == size+1 {
		return -1
	}

	return res - 1
}
