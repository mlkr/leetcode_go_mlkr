package problem119

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	if rowIndex == 1 {
		return []int{1, 1}
	}

	idx := 1
	pre := []int{1, 1}
	var get func()
	get = func() {
		if idx == rowIndex {
			return
		}

		idx++
		next := make([]int, idx+1)
		next[0], next[idx] = 1, 1
		for i := 1; i < idx; i++ {
			next[i] = pre[i] + pre[i-1]
		}

		pre = next
		get()
	}

	get()

	return pre
}

// 另一解法
func getRow2(rowIndex int) []int {
	res := make([]int, 1, rowIndex+1)
	res[0] = 1
	if rowIndex == 0 {
		return res
	}

	for i := 1; i <= rowIndex; i++ {
		res = append(res, 1)
		for j := len(res) - 2; j > 0; j-- {
			res[j] += res[j-1]
		}
	}

	return res
}
