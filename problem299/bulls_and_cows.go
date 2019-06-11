package problem299

import "strconv"

func getHint(secret string, guess string) string {
	size := len(secret)
	had := make(map[byte]int, size)
	for i := 0; i < size; i++ {
		had[secret[i]]++
	}

	A, B := 0, 0
	for i := 0; i < size; i++ {
		if secret[i] == guess[i] {
			A++
			had[guess[i]]--
		}
	}

	for i := 0; i < size; i++ {
		if secret[i] != guess[i] && had[guess[i]] > 0 {
			B++
			had[guess[i]]--
		}
	}

	return strconv.Itoa(A) + "A" + strconv.Itoa(B) + "B"
}

// 解法二
func getHint2(secret string, guess string) string {
	countS, countG := make([]int, 10), make([]int, 10)

	size := len(secret)
	A, B := 0, 0
	for i := 0; i < size; i++ {
		if secret[i] == guess[i] {
			A++
		} else {
			ns, ng := secret[i]-'0', guess[i]-'0'

			if countS[ng] > 0 {
				B++
				countS[ng]--
			} else {
				countG[ng]++
			}

			if countG[ns] > 0 {
				B++
				countG[ns]--
			} else {
				countS[ns]++
			}
		}
	}

	return strconv.Itoa(A) + "A" + strconv.Itoa(B) + "B"
}
