package problem44

func isMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)

	temp := [][]bool{}
	for i := 0; i < sLen+1; i++ {
		temp = append(temp, make([]bool, pLen+1))
	}

	temp[0][0] = true
	// temp[i][j] 表示 s[:i+1] 与 p[:i+1] 匹配

	for j := 1; j <= pLen; j++ {
		if p[j-1] == '*' {
			temp[0][j] = temp[0][j-1]
		}
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if p[j-1] != '*' {
				temp[i][j] = (s[i-1] == p[j-1] || p[j-1] == '?') && temp[i-1][j-1]
			} else {
				temp[i][j] = temp[i-1][j] || temp[i][j-1]
			}
		}
	}

	return temp[sLen][pLen]
}
