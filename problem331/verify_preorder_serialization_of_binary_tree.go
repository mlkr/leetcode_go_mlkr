package problem331

import (
	"strings"
)

func isValidSerialization(preorder string) bool {
	if preorder == "#" {
		return true
	}

	str := strings.Split(preorder, ",")
	size := len(str)

	// 长度为奇数并且大于3
	if size%2 != 1 || size < 3 {
		return false
	}

	// 末尾必为 #
	if str[size-1] != "#" {
		return false
	}

	for i := size - 3; i >= 0; i-- {
		if i <= len(str)-3 && str[i] != "#" && str[i+1] == "#" && str[i+2] == "#" {
			temp := str[i+3:]
			str = append(str[:i+1], temp...)
			str[i] = "#"
		}
	}

	if len(str) == 1 && str[0] == "#" {
		return true
	}

	return false
}
