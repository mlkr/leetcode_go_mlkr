package problem363

import (
	"math"
	"sort"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
	r := len(matrix)
	c := len(matrix[0])

	res := math.MinInt64
	for l := 0; l < c; l++ {

		rowSum := make([]int, r)
		for b := l; b < c; b++ {

			// 计算矩形 acc 的大小
			acc := 0
			maxAcc := 0
			for i := 0; i < r; i++ {
				rowSum[i] += matrix[i][b]
				acc += rowSum[i]

				if acc <= k && acc > res {
					res = acc
				}

				if res == k {
					return res
				}

				if acc < 0 {
					acc = 0
				} else {
					maxAcc = max(maxAcc, acc)
				}
			}

			if maxAcc < k {
				continue
			}

			// 最大矩形大于 k, 上下剪切
			sums := []int{}
			acc = 0
			for i := 0; i < r; i++ {
				acc += rowSum[i]

				sort.Ints(sums)
				t := sort.SearchInts(sums, acc-k)
				if t < len(sums) {
					if acc-sums[t] > res {
						res = acc - sums[t]
						if res == k {
							return res
						}
					}
				}

				sums = append(sums, acc)
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
