package problem395

import "strings"

func longestSubstring(s string, k int) int {
	size := len(s)
	if size == 0 {
		return 0
	}

	count := make([]int, 26)
	for i := 0; i < size; i++ {
		count[int(s[i]-'a')]++
	}

	badStr := ""
	for i := range count {
		if count[i] > 0 && count[i] < k {
			badStr = string('a' + i)
			break
		}
	}

	if badStr == "" {
		return size
	}

	res := 0
	strs := strings.Split(s, badStr)
	for _, str := range strs {
		res = max(res, longestSubstring(str, k))
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
