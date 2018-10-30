package problem140

import (
	"sort"
)

func wordBreak(s string, wordDict []string) []string {
	if len(wordDict) == 0 {
		return nil
	}

	dict := make(map[string]bool)
	lengths := make(map[int]bool)
	for _, word := range wordDict {
		dict[word] = true
		lengths[len(word)] = true
	}

	sizes := make([]int, 0, len(lengths))
	for length := range lengths {
		sizes = append(sizes, length)
	}
	sort.Ints(sizes)

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	relation := make(map[int][]int)
	for i := 0; i <= n; i++ {
		if !dp[i] {
			continue
		}

		for _, size := range sizes {
			if i+size <= n {
				dp[i+size] = dp[i+size] || dict[s[i:i+size]]

				if dict[s[i:i+size]] {
					relation[i] = append(relation[i], i+size)
				}
			}
		}
	}

	if !dp[n] {
		return nil
	}

	res := []string{}

	var getWord func(idx int, str string)
	getWord = func(idx int, str string) {
		if idx == n {
			res = append(res, str[1:])
			return
		}

		for _, nextIdx := range relation[idx] {
			getWord(nextIdx, str+" "+s[idx:nextIdx])
		}
	}

	getWord(0, "")

	return res
}

// 最佳
func wordBreak2(s string, wordDict []string) []string {
	if len(wordDict) <= 0 {
		return nil
	}

	dict := make(map[string]bool, len(wordDict))
	lengths := make(map[int]bool, len(wordDict))
	for _, word := range wordDict {
		dict[word] = true
		lengths[len(word)] = true
	}

	sizes := make([]int, 0, len(lengths))
	for length := range lengths {
		sizes = append(sizes, length)
	}
	sort.Ints(sizes)

	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		if dp[i] == 0 {
			continue
		}

		for _, size := range sizes {
			if i+size <= n && dict[s[i:i+size]] {
				dp[i+size] += dp[i]
			}
		}
	}

	if dp[n] == 0 {
		return nil
	}

	res := make([]string, 0, dp[n])

	var dfs func(i int, str string)
	dfs = func(i int, str string) {
		if i == n {
			res = append(res, str[1:])
			return
		}

		for _, size := range sizes {
			if i+size <= n && dict[s[i:i+size]] {
				dfs(i+size, str+" "+s[i:i+size])
			}
		}
	}
	dfs(0, "")

	return res
}

func wordBreak3(s string, wordDict []string) []string {
	if len(wordDict) <= 0 {
		return nil
	}

	dict := make(map[string]bool, len(wordDict))
	lengths := make(map[int]bool, len(wordDict))
	for _, word := range wordDict {
		dict[word] = true
		lengths[len(word)] = true
	}

	sizes := make([]int, 0, len(lengths))
	for length := range lengths {
		sizes = append(sizes, length)
	}
	sort.Ints(sizes)

	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	relation := make(map[int][]int, len(sizes))
	for i := 0; i < n; i++ {
		if dp[i] == 0 {
			continue
		}

		for _, size := range sizes {
			if i+size <= n && dict[s[i:i+size]] {
				dp[i+size] += dp[i]
				relation[i] = append(relation[i], i+size)
			}
		}
	}

	if dp[n] == 0 {
		return nil
	}

	res := make([]string, 0, dp[n])

	var dfs func(i int, str string)
	dfs = func(i int, str string) {
		if i == n {
			res = append(res, str[1:])
			return
		}

		for _, next := range relation[i] {
			dfs(next, str+" "+s[i:next])
		}
	}
	dfs(0, "")

	return res
}
