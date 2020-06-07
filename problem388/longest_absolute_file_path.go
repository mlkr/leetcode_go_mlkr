package problem388

import "strings"

func lengthLongestPath(input string) int {
	if !hasFile(input) {
		return 0
	}

	return llp("\n"+input) - 1
}

func llp(path string) int {
	if !hasFile(path) {
		return 0
	}

	if !hasDir(path) {
		return len(path)
	}

	dirLen := nextCR(path, -1)
	i := dirLen
	maxSub := 0
	for i < len(path) {
		j := nextCR(path, i)
		maxSub = max(maxSub, llp(strings.ReplaceAll(path[i+1:j], "\n\t", "\n")))
		i = j
	}

	return dirLen + 1 + maxSub
}

// 下一个回车的位置
func nextCR(path string, i int) int {
	size := len(path)
	for i++; i < size; i++ {
		if path[i:i+1] == "\n" && (i+1 == size || path[i+1:i+2] != "\t") {
			break
		}
	}

	return i
}

func hasFile(path string) bool {
	return strings.Contains(path, ".")
}

func hasDir(path string) bool {
	return strings.Contains(path, "\n")
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
