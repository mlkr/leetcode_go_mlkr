package problem401

import "fmt"

func readBinaryWatch(num int) []string {
	res := make([]string, 0, 10)
	leds := make([]int, 10)

	var dfs func(idx, num int)
	dfs = func(idx, num int) {
		if num == 0 {
			h, m := get(leds[:4]), get(leds[4:])
			if h < 12 && m < 60 {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}

			return
		}

		for i := idx; i < 10-(num-1); i++ {
			leds[i] = 1
			dfs(i+1, num-1)
			leds[i] = 0
		}
	}

	dfs(0, num)
	return res
}

var bs = [6]int{1, 2, 4, 8, 16, 32}

func get(leds []int) int {
	sum := 0
	for k, v := range leds {
		if v == 1 {
			sum += bs[k]
		}
	}

	return sum
}
