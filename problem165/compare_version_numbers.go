package problem165

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	if version1 == version2 {
		return 0
	}

	ver1 := strings.Split(version1, ".")
	ver2 := strings.Split(version2, ".")

	len1 := len(ver1)
	len2 := len(ver2)

	size := 0
	res := 0
	remain := []string{}
	if len1 > len2 {
		size = len2
		res = 1
		remain = ver1[len2:]
	} else {
		size = len1
		res = -1
		remain = ver2[len1:]
	}

	for i := 0; i < size; i++ {
		if ver1[i] == ver2[i] {
			continue
		}

		num1, _ := strconv.Atoi(ver1[i])
		num2, _ := strconv.Atoi(ver2[i])
		if num1 > num2 {
			return 1
		}

		if num1 < num2 {
			return -1
		}

		return 0
	}

	for _, str := range remain {
		num, _ := strconv.Atoi(str)
		if num != 0 {
			return res
		}
	}

	return 0
}
