package problem139

import (
	"sort"
)

func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return false
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
	// dp[i]=true 表明 wordBreak(s[:i+1], wordDict) = true
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i <= n; i++ {
		if !dp[i] {
			continue
		}

		for _, size := range sizes {
			if i+size <= n {
				dp[i+size] = dp[i+size] || dict[s[i:i+size]]
			}
		}

	}

	return dp[n]
}
