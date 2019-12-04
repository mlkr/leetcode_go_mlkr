package problem342

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}

	// 如果 num 不是2的N次方, 那么也不是4的N次方
	if num&(num-1) != 0 {
		return false
	}

	// 8421 8421
	// 0101 0101
	// 除以4 等于向右移动两位, 只有当num&0x55555555 > 0 才符合要求
	if num&0x55555555 == 0 {
		return false
	}

	return true
}
