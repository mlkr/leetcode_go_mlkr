package problem638

import "strconv"

func shoppingOffers(price []int, special [][]int, needs []int) int {
	filter := make([][]int, 0, len(special))
	for _, offer := range special {
		total := 0
		for i := 0; i < len(price); i++ {
			total += offer[i] * price[i]
		}

		if offer[len(offer)-1] < total {
			filter = append(filter, offer)
		}
	}

	mem := make(map[string]int)
	return dfs(price, filter, needs, mem)
}

func dfs(price []int, special [][]int, needs []int, mem map[string]int) int {
	lowPrice := 0
	key := toKey(needs)
	if lowPrice, ok := mem[key]; ok {
		return lowPrice
	}

	for i := 0; i < len(price); i++ {
		lowPrice += price[i] * needs[i]
	}

	for _, offer := range special {
		isValid := true
		for i := 0; i < len(needs); i++ {
			if offer[i] > needs[i] {
				isValid = false
				break
			}
		}

		if !isValid {
			continue
		}

		for i := 0; i < len(needs); i++ {
			needs[i] -= offer[i]
		}

		lowPrice = min(lowPrice, dfs(price, special, needs, mem)+offer[len(offer)-1])

		for i := 0; i < len(needs); i++ {
			needs[i] += offer[i]
		}

	}

	mem[key] = lowPrice

	return lowPrice
}

func toKey(needs []int) string {
	key := ""
	for i := range needs {
		key += strconv.Itoa(needs[i])
	}

	return key
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
