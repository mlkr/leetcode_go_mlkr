package problem433

func minMutation(start string, end string, bank []string) int {
	count := 0
	visited := make([]bool, len(bank))

	var bfs func(queue []string)
	bfs = func(queue []string) {
		if len(queue) == 0 {
			count = -1
			return
		}

		newQueue := []string{}
		for _, q := range queue {
			if q == end {
				return
			}

			for i, b := range bank {
				if visited[i] || !isMutation(q, b) {
					continue
				}

				visited[i] = true
				newQueue = append(newQueue, b)
			}

		}

		count++
		bfs(newQueue)
	}

	queue := []string{start}
	bfs(queue)

	return count
}

func isMutation(s1, s2 string) bool {
	count := 0
	for i := 0; i < 8; i++ {
		if count > 1 {
			return false
		}

		if s1[i] != s2[i] {
			count++
		}
	}

	return count == 1
}
