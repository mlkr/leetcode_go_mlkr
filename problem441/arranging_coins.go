package problem441

func arrangeCoins(n int) int {
	i:=0
	count := 0
	for count <= n{
		i++
		count += i
	}

	return i-1
}

// 二分法
func arrangeCoins2(n int) int {
	left, right := 0, n
	for left <= right {
		mid := (left+right)/2

		// 计算mid行硬币总数
		count := (mid+1)*mid/2
		if count == n {
			return mid
		}

		if count < n {
			left = mid + 1 
		}else{
			right = mid - 1
		}
	}

	return right
}