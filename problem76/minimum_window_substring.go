package problem76

func minWindow(s string, t string) string {
	sLen := len(s)
	tLen := len(t)

	need := [256]int{}
	has := [256]int{}

	for i := range t {
		need[t[i]]++
	}

	min := sLen + 1
	begin, end, count, windowStart, windowEnd := 0, 0, 0, 0, 0
	for ; end < sLen; end++ {
		if need[s[end]] == 0 {
			continue
		}

		if has[s[end]] < need[s[end]] {
			count++
		}

		has[s[end]]++

		if count == tLen {
			for has[s[begin]] > need[s[begin]] || need[s[begin]] == 0 {
				if has[s[begin]] > need[s[begin]] {
					has[s[begin]]--
				}

				begin++
			}

			temp := end - begin + 1
			if temp < min {
				min = temp
				windowStart = begin
				windowEnd = end
			}
		}
	}

	if count < tLen {
		return ""
	}

	return s[windowStart : windowEnd+1]
}
