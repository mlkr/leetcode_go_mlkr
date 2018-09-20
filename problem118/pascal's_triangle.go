package problem118

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	if numRows == 1 {
		return [][]int{
			[]int{1},
		}
	}

	res := [][]int{
		[]int{1},
		[]int{1, 1},
	}

	if numRows == 2 {
		return res
	}

	n := 2
	pre := []int{1, 1}
	var getN func()
	getN = func() {
		n++
		next := make([]int, n)
		next[0], next[n-1] = 1, 1
		for i := 1; i < n-1; i++ {
			next[i] = pre[i-1] + pre[i]
		}

		res = append(res, next)
		pre = next

		if n == numRows {
			return
		}

		getN()
	}

	getN()

	return res
}
