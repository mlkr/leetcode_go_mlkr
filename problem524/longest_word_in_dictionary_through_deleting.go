package problem524

import "sort"

func findLongestWord(s string, dict []string) string {
	// 字符串长度从大到小排列
	sort.Slice(dict, func(i, j int) bool {
		if len(dict[i]) != len(dict[j]) {
			return len(dict[i]) > len(dict[j])
		}

		// 字符串长度相同，字典顺序排列
		return dict[i] < dict[j]
	})

	for _, str := range dict {
		if isSub(str, s) {
			return str
		}
	}

	return ""
}

// str1 是否为 str2 的子串
func isSub(str1, str2 string) bool {
	l1, l2 := len(str1), len(str2)
	if l1 > l2 {
		return false
	}

	i, j := 0, 0
	for i < l1 && j < l2 {
		if str1[i] == str2[j] {
			i++
		}

		j++
	}

	return i == l1
}

// 另一种解法
func findLongestWord2(s string, dict []string) string {
	m := make([][]int, 26)

	for i, letter := range s {
		m[letter-'a'] = append(m[letter-'a'], i)
	}

	res := ""
	for _, word := range dict {
		if len(word) < len(res) || (len(word) == len(res) && word >= res) {
			continue
		}

		sIndex := -1
		isValid := true
		for wIndex, letter := range word {
			if len(s)-sIndex-1 < len(word)-wIndex {
				isValid = false
				break
			}

			// 在 s 中查找 letter
			index := letter - 'a'
			i := sort.Search(len(m[index]), func(i int) bool {
				return m[index][i] > sIndex
			})

			if i == len(m[index]) {
				// 查找失败
				isValid = false
				break
			}

			// 查找成功
			sIndex = m[index][i]
		}

		if isValid {
			res = word
		}

	}

	return res
}
