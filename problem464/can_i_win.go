package problem464

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	sum := (maxChoosableInteger + 1) / 2 * maxChoosableInteger
	if sum < desiredTotal {
		return false
	}

	// 记录已使用的数
	bits := make([]int, maxChoosableInteger+1)
	for i := maxChoosableInteger; i > 0; i-- {
		bits[i] = 1 << uint8(i)
	}

	// 记录路径
	dp := make(map[int]bool)

	return dfs(0, maxChoosableInteger, desiredTotal, bits, dp)
}

func dfs(usedBit, max, remain int, bits []int, dp map[int]bool) bool {
	if res, ok := dp[usedBit]; ok {
		return res
	}

	if max >= remain {
		for i := max; i >= remain; i-- {
			if usedBit&bits[i] != 0 {
				// 使用过 i
				continue
			}

			dp[usedBit] = true
			return true
		}
	}

	for i := max; i > 0; i-- {
		if usedBit&bits[i] != 0 {
			// 使用过 i
			continue
		}

		old := usedBit
		usedBit |= bits[i]
		isOpponentWin := dfs(usedBit, max, remain-i, bits, dp)
		usedBit = old

		if isOpponentWin == false {
			// 对手输了 就是我赢了
			dp[usedBit] = true
			return true
		}
	}

	dp[usedBit] = false
	return false
}
