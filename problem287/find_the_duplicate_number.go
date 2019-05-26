package problem287

// 查看 https://blog.csdn.net/wr339988/article/details/53617914
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}

	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}

	return slow
}
