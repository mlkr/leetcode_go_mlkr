package problem49

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	var res [][]string

	for _, str := range strs {
		key := getKey(str)
		strMap[key] = append(strMap[key], str)
	}

	for _, v := range strMap {
		sort.Strings(v)
		res = append(res, v)
	}

	return res
}

func getKey(str string) string {
	char := []byte(str)
	charLen := len(char)
	temp := make([]int, charLen)
	for i := 0; i < charLen; i++ {
		temp[i] = int(char[i])
	}

	sort.Ints(temp)
	for k, v := range temp {
		char[k] = byte(v)
	}

	return string(char)
}
