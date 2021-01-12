package problem467

func findSubstringInWraproundString(p string) int {
	size := len(p)
	if size <= 1 {
		return size
	}

	cnt := 1
	count := make([]int, 26)
	for i := 0; i < size; i++ {
		// 计算以 p[i] 结尾的字符串个数
		if i > 0 &&
			(p[i-1]+1 == p[i] || p[i-1]-25 == p[i]) {
			cnt++
		} else {
			cnt = 1
		}

		count[p[i]-'a'] = max(count[p[i]-'a'], cnt)
	}

	res := 0
	for _, cnt := range count {
		res += cnt
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
