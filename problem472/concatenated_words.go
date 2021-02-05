package problem472

import "sort"

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	res := []string{}
	hasWords := make(map[string]struct{})
	for _, word := range words {
		if canForm(word, hasWords) {
			res = append(res, word)
		}

		hasWords[word] = struct{}{}
	}

	return res
}

func canForm(word string, hasWords map[string]struct{}) bool {
	if len(hasWords) == 0 {
		return false
	}

	dp := make([]bool, len(word)+1)
	dp[0] = true

	for i := 1; i <= len(word); i++ {
		for j := i - 1; j >= 0; j-- {
			if !dp[j] {
				continue
			}

			if _, ok := hasWords[word[j:i]]; ok {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(word)]
}
