package problem378

// 二分法求解
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	lo, hi := matrix[0][0], matrix[n-1][n-1]

	for lo < hi {
		mid := (lo+hi)>>1
		count := 0

		j := n-1
		for i:=0;i<n;i++{
			if matrix[i][0] > mid {
				break
			}

			for j>=0 && matrix[i][j] > mid {
				j--
			}
			count += j+1
		}

		if count < k {
			lo = mid+1
		}else{
			hi = mid
		}
	}

	return lo
}