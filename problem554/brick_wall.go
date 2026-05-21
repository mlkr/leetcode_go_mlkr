package problem554

func leastBricks(wall [][]int) int {
	n := len(wall)

	// 统计坐标相同的 edge 出现的次数
	count := make(map[int]int, n)

	for i := 0; i < n; i++ {
		sum := 0

		// 每行的总宽度相同,不用计算最后一个
		for j := 0; j < len(wall[i])-1; j++ {
			sum += wall[i][j]
			count[sum]++
		}
	}

	max := 0
	for _, edges := range count {
		if edges > max {
			max = edges
		}
	}

	return n - max
}
