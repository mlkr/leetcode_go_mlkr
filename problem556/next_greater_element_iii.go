package problem556

import (
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	s := []byte(strconv.Itoa(n))
	size := len(s)

	for i := size - 2; i >= 0; i-- {
		// 从右往左 找到第一个比右边小的数字 位置为i
		if s[i] < s[i+1] {

			// 从右往左 找到比第一个比s[i]的数字 位置为j
			j := size - 1
			for ; j > i; j-- {
				if s[j] > s[i] {
					break
				}
			}

			// i, k 互换
			s[i], s[j] = s[j], s[i]

			// 升序排列剩下的数
			// sort asc from i+1 till end
			sort.Slice(s[i+1:], func(m, n int) bool {
				return s[i+1:][m] < s[i+1:][n]
			})

			// check if overflow
			res, _ := strconv.Atoi(string(s))
			if res > 1<<31-1 {
				return -1
			}

			return res
		}
	}

	return -1
}
