package problem179

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	size := len(nums)
	strs := make([]string, 0, size)
	strs = append(strs, strconv.Itoa(nums[0]))

	for i := 1; i < size; i++ {
		str := strconv.Itoa(nums[i])
		n := insert(str, strs, 0)

		strsc := make([]string, n)
		copy(strsc, strs[:n])
		strs = append(append(strsc, str), strs[n:]...)
	}

	res := ""
	for _, str := range strs {
		if res == "0" {
			res = ""
		}
		res += str
	}

	return res
}

// 从大到小排列
func insert(str string, strs []string, lastN int) int {
	size := len(strs)
	if size == 0 {
		return lastN
	}

	n := size / 2
	switch compare(str, strs[n]) {
	case 0:
		return n + lastN
	case 1:
		return insert(str, strs[:n], lastN+0)
	case -1:
		return insert(str, strs[n+1:], lastN+n+1)
	}

	return lastN
}

// 比较两个数
// str1 > str2 return 1
// str1 = str2 return 0
// str1 < str2 return -1
// 34 > 3 > 30
func compare(str1, str2 string) int {
	size := len(str1) + len(str2)
	str12 := str1 + str2
	str21 := str2 + str1

	for i := 0; i < size; i++ {
		if str12[i] > str21[i] {
			return 1
		}

		if str12[i] < str21[i] {
			return -1
		}
	}

	return 0
}

// 另一解法(最佳)
func largestNumber2(nums []int) string {
	size := len(nums)
	resSize := 0
	b := make(bytes, size)
	for i := 0; i < size; i++ {
		b[i] = []byte(strconv.Itoa(nums[i]))
		resSize += len(b[i])
	}

	sort.Sort(b)

	res := make([]byte, 0, resSize)
	for i := size - 1; i >= 0; i-- {
		res = append(res, b[i]...)
	}

	i := 0
	for i < resSize-1 && res[i] == '0' {
		i++
	}

	return string(res[i:])
}

type bytes [][]byte

func (b bytes) Len() int {
	return len(b)
}

func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b bytes) Less(i, j int) bool {
	size := len(b[i]) + len(b[j])
	bij := append(b[i], b[j]...)
	bji := append(b[j], b[i]...)

	for i := 0; i < size; i++ {
		if bij[i] < bji[i] {
			return true
		}

		if bij[i] > bji[i] {
			return false
		}
	}

	return false
}
