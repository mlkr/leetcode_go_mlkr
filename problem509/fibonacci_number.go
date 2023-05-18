package problem509

// func fib(n int) int {
// 	if n == 0 {
// 		return 0
// 	}

// 	if n == 1 {
// 		return 1
// 	}

// 	return fib(n-1) + fib(n-2)
// }

func fib(n int) int {
	nums := []int{0, 1}

	for i := 2; i <= n; i++ {
		nums = append(nums, nums[i-1]+nums[i-2])
	}

	return nums[n]
}
