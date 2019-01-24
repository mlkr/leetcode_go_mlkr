package problem210

func findOrder(numCourses int, prerequisites [][]int) []int {
	next := findNext(numCourses, prerequisites)
	preCount := findPreCount(next)
	return check(next, preCount)
}

func findNext(numCourses int, prerequisites [][]int) [][]int {
	next := make([][]int, numCourses)
	for _, pre := range prerequisites {
		next[pre[1]] = append(next[pre[1]], pre[0])
	}

	return next
}

func findPreCount(next [][]int) []int {
	preCount := make([]int, len(next))
	for _, v := range next {
		for _, vv := range v {
			preCount[vv]++
		}
	}

	return preCount
}

func check(next [][]int, preCount []int) []int {
	size := len(next)
	res := make([]int, 0, size)

	for i := 0; i < size; i++ {
		if preCount[i] != 0 {
			continue
		}
		preCount[i] = -1

		res = append(res, i)

		for _, n := range next[i] {
			preCount[n]--
		}

		i = -1
	}

	if len(res) != size {
		return []int{}
	}

	return res
}
