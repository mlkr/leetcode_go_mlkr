package problem438

func findAnagrams(s string, p string) []int {
	var window, target [26]int
	for i := range p {
		target[p[i]-'a']++
	}

	res := []int{}
	pSize := len(p)
	for i := range s {
		window[s[i]-'a']++

		if i >= pSize {
			window[s[i-pSize]-'a']--
		}

		if window == target {
			res = append(res, i-pSize+1)
		}
	}

	return res
}

// 最优
func findAnagrams2(s string, p string) []int {
	freq := [26]int{}
	for _, v := range p {
		freq[v-'a']++
	}

	left, right := 0, 0
	count := len(p)
	res := []int{}

	for right < len(s) {
		// 减少
		ch := s[right]
		if freq[ch-'a'] > 0 {
			count--
		}

		freq[ch-'a']--
		right++

		if count == 0 {
			res = append(res, left)
		}

		// 恢复
		if right-left == len(p) {
			ch = s[left]
			if freq[ch-'a'] >= 0 {
				count++
			}

			freq[ch-'a']++
			left++
		}
	}

	return res
}
