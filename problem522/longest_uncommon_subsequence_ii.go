package problem523

import "sort"

func findLUSlength(strs []string) int {
	// 记录字符串出现的次数
	count := make(map[string]int)
	for _, s := range strs {
		count[s]++
	}

	// 去掉重复的字符串
	strs = strs[:0]
	for s := range count {
		strs = append(strs, s)
	}

	// 字符长度从大到小排序
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})

	for i := 0; i < len(strs); i++ {
		if count[strs[i]] > 1 {
			continue
		}

		if !isSub(strs[i], strs[:i]) {
			return len(strs[i])
		}
	}

	return -1
}

func isSub(s string, strs []string) bool {
	for _, str := range strs {
		if isSubOf(s, str) {
			return true
		}
	}

	return false
}

// s1是否为s2的子串
func isSubOf(s1, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	i, j := 0, 0

	for i < l1 && j < l2 {
		if s1[i] == s2[j] {
			i++
		}

		j++
	}

	return i == l1
}
