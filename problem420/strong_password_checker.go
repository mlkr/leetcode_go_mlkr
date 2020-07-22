package problem420

func strongPasswordChecker(s string) int {
	// 小写 大写字母 数字 至少各一个
	l, u, d, missingType := 1, 1, 1, 3
	for i := range s {
		switch {
		case l > 0 && 'a' <= s[i] && s[i] <= 'z':
			l--
			missingType--
		case u > 0 && 'A' <= s[i] && s[i] <= 'Z':
			u--
			missingType--
		case d > 0 && '0' <= s[i] && s[i] <= '9':
			d--
			missingType--
		}

		if missingType == 0 {
			break
		}
	}

	// 算出要替换的字数
	replace := 0
	ones, twos := 0, 0
	for i := 0; i+2 < len(s); i++ {
		if s[i] != s[i+1] || s[i+1] != s[i+2] {
			continue
		}

		replaceLen := 2
		for ; i+2 < len(s) && s[i] == s[i+2]; i++ {
			replaceLen++
		}

		// 每3个连续的字母替换一个为另外的字母
		replace += replaceLen / 3
		switch replaceLen % 3 {
		case 0:
			// 如果长度超过20, 减少1个字母 少替换1个字母
			ones++
		case 1:
			// 如果长度超过20, 减少2个字母 少替换1个字母
			twos++
		}
	}

	if len(s) < 6 {
		return max(missingType, 6-len(s))
	}

	if len(s) <= 20 {
		return max(missingType, replace)
	}

	delete := len(s) - 20

	replace -= min(ones, delete)
	replace -= min(max(delete-ones, 0), twos*2) / 2
	// 剩下的 每减少3个字母 少替换1个字母
	replace -= max(delete-ones-twos*2, 0) / 3

	return delete + max(missingType, replace)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
