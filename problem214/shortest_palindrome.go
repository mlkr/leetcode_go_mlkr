package problem214

func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	// 例子
	// "aacecaaa#aaacecaa"
	i := getIndex(s + "#" + reverse(s))
	return reverse(s[i:]) + s
}

// 例子
// aacecaaa
// 01234567
// -----
// 0 a
// 1 aa
// 0 aac
// 0 aace
// 0 aacec
// 1 aaceca
// 2 aacecaa
// 2 aacecaaa
func getIndex(s string) int {
	size := len(s)
	// table[i] 是 s[:i+1] 的前缀集与后缀集中，最长公共元素的长度
	// "abcd" 的前缀集是 ["", "a", "ab", "abc"]
	// "abcd" 的后缀集是 ["", "d", "cd", "bcd"]
	// "abcd" 的前缀集与后缀集的最长公共元素是"", 它的长度是 0
	table := make([]int, size)

	i, klen := 1, 0
	for i < size {
		if s[i] == s[klen] {
			klen++
			table[i] = klen
			i++
		} else {
			if klen == 0 {
				table[i] = 0
				i++
			} else {
				klen = table[klen-1]
			}
		}
	}

	return table[size-1]
}

func reverse(s string) string {
	size := len(s)
	res := make([]byte, size)
	for i := 0; i < size; i++ {
		res[size-1-i] = s[i]
	}

	return string(res)
}
