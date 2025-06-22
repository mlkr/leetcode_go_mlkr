package problem638

import "strconv"

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(needs)

	filter := make([][]int, 0, len(special))
	for _, s := range special {
		sum := 0
		for i := 0; i < n; i++ {
			sum += price[i] * s[i]
		}

		if sum > s[n] {
			filter = append(filter, s)
		}
	}

	m := make(map[string]int)
	return dfs(price, filter, needs, m)
}

func dfs(price []int, special [][]int, needs []int, m map[string]int) int {
	key := toKey(needs)
	if _, ok := m[key]; ok {
		return m[key]
	}

	n := len(needs)

	lowPrice := 0
	for i := 0; i < n; i++ {
		lowPrice += needs[i] * price[i]
	}

	for _, s := range special {
		failed := false
		for i := 0; i < n; i++ {
			if needs[i] < s[i] {
				failed = true
				break
			}
		}

		if failed {
			continue
		}

		for i := 0; i < n; i++ {
			needs[i] -= s[i]
		}

		lowPrice = min(lowPrice, dfs(price, special, needs, m)+s[n])

		for i := 0; i < n; i++ {
			needs[i] += s[i]
		}
	}

	m[key] = lowPrice

	return lowPrice
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func toKey(needs []int) string {
	key := ""
	for _, need := range needs {
		key += strconv.Itoa(need)
	}

	return key
}
