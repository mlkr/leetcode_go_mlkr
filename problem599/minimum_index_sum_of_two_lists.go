package problem599

func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int, len(list1))
	for i, s := range list1 {
		m[s] = i
	}

	min := len(list1) + len(list2)
	res := []string{}
	for i, s := range list2 {
		if j, ok := m[s]; ok {
			if i+j == min {
				res = append(res, s)
			}

			if i+j < min {
				min = i + j
				res = []string{s}
			}
		}
	}

	return res
}
