package problem207

func canFinish(numCourses int, prerequisites [][]int) bool {
	next := getNext(numCourses, prerequisites)
	dc := getDependencyCount(next)
	return check(next, dc)
}

// 每个课程可以学习的下一个课程
func getNext(numCourses int, prerequisites [][]int) [][]int {
	next := make([][]int, numCourses)
	for _, pre := range prerequisites {
		next[pre[1]] = append(next[pre[1]], pre[0])
	}

	return next
}

// 每个课程依赖的课程总数
func getDependencyCount(next [][]int) []int {
	dc := make([]int, len(next))
	for _, n := range next {
		for _, c := range n {
			dc[c]++
		}
	}

	return dc
}

// 检查死锁
func check(next [][]int, dc []int) bool {
	size := len(next)
	i, j := 0, 0

	for i = 0; i < size; i++ {
		for j = 0; j < size; j++ {
			// 找出依赖项为0的课程
			if dc[j] == 0 {
				break
			}
		}

		if j == size {
			// 没有依赖项为0的课程, 返回false
			return false
		}

		// 排除
		dc[j] = -1

		for _, n := range next[j] {
			dc[n]--
		}

	}

	return true
}
