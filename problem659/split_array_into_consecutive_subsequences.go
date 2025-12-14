package problem659

func isPossible(nums []int) bool {
	size := len(nums)

	// l1 以当前数字结尾 长度为1的序列数
	// l2 以当前数字结尾 长度为2的序列数
	// l3p 以当前数字结尾 长度为>=3的序列数
	// 一个数只能在一个序列中
	l1, l2, l3p := 0, 0, 0

	// 默认前值与现值相差2，使其不连续
	pre := nums[0] - 2

	for i := 0; i < size; i++ {
		cur := nums[i]
		count := 1
		for i+1 < size && nums[i] == nums[i+1] {
			i++
			count++
		}

		// 不连续
		if pre+1 != cur {
			if l1 > 0 || l2 > 0 {
				return false
			}

			// 重置状态
			l3p = 0
		}

		// 将cur 分配给l1 l2 序列
		count -= l1 + l2
		if count < 0 {
			// 不够分 出现空档
			return false
		}

		pre = cur

		if count <= l3p {
			// 剩余的cur 全部分配给l3p序列
			l3p = count + l2
			l2 = l1
			l1 = 0
			continue
		}

		// 剩余的cur的数量 比l3p数量多
		newl1 := count - l3p
		l3p += l2
		l2 = l1
		l1 = newl1
	}

	return l1 == 0 && l2 == 0
}
