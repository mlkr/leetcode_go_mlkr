package problem134

func canCompleteCircuit(gas []int, cost []int) int {
	// 剩余油量
	remains := 0
	// 开始位置
	start := 0
	// 欠油量
	debts := 0
	for k, v := range gas {
		remains += v - cost[k]
		if remains < 0 {
			debts += remains
			start = k + 1
			remains = 0
		}
	}

	if debts+remains < 0 {
		return -1
	}

	return start
}
