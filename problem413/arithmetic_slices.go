package problem413

// 参看
// 1 2 3 	    1
// 1 2 3 4	    2
// 1 2 3 4 5	3
// https://blog.csdn.net/camellhf/article/details/52824234
func numberOfArithmeticSlices(A []int) int {
	added := 0
	count := 0

	for i := 2; i < len(A); i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			added++
			count += added
		} else {
			added = 0
		}
	}

	return count
}
