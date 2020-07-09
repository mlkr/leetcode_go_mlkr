package problem409

func longestPalindrome(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}

	count := make([]int, 58)
	for i:=0;i<size;i++{
		count[s[i]-'A']++
	}

	res := 0
	isOdd := 0
	for i := range count {
		if count[i] & 1 == 0 {
			res += count[i]
		}else{
			res += count[i]-1
			isOdd = 1
		}
	}

	return res + isOdd
}

