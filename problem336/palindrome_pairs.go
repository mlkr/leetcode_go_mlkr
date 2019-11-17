package problem336

func palindromePairs(words []string) [][]int {

	size := len(words)
	res := make([][]int, 0, size)

	if size < 2 {
		return res
	}

	hash := make(map[string]int, size)
	for i, word := range words {
		hash[word] = i
	}

	for i := 0; i < size; i++ {
		for k := 0; k <= len(words[i]); k++ {
			left := words[i][:k]
			right := words[i][k:]

			// len(left) != 0 为了防止重复添加
			if len(left) != 0 && isPalindrome(left) {
				rightRev := reverse(right)
				j, ok := hash[rightRev]
				if ok == true && i != j {
					res = append(res, []int{j, i})
				}
			}

			if isPalindrome(right) {
				leftRev := reverse(left)
				j, ok := hash[leftRev]
				if ok == true && i != j {
					res = append(res, []int{i, j})
				}
			}
		}
	}

	return res
}

func isPalindrome(str string) bool {
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func reverse(str string) string {
	bytes := []byte(str)
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	return string(bytes)
}
