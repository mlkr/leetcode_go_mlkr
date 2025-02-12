package problem621

// 出现次数最多的字母记为A，次数为maxC
// A-------
// A-------
// A-------
// A-------
// A-------
// A-------
// A
// 运行时间为 (maxC-1)*(n+1)+1

// 如果B出现的次数和A相同
// AB------
// AB------
// AB------
// AB------
// AB------
// AB------
// AB
// 运行时间为 (maxC-1)*(n+1)+2

// 如果上面的运行时间不够运行所有的tasks, 则运行时间为len(tasks)
// ABCDFHKNR
// ABCDFHKNS
// ABCEGILOT
// ABCEGILO
// ABDEGJMP
// ABDEGJMQ
// AB

func leastInterval(tasks []byte, n int) int {
	count := [26]int{}
	for _, task := range tasks {
		count[task-'A']++
	}

	maxC := 0
	for _, c := range count {
		maxC = max(maxC, c)
	}

	res := 0
	for _, c := range count {
		if c == maxC {
			res++
		}
	}

	return max((maxC-1)*(n+1)+res, len(tasks))

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
