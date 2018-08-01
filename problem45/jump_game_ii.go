package problem45

func jump(nums []int) int {
	i, end, count := 0, len(nums)-1, 0

	for i < end {
		if nums[i]+i > end {
			return count + 1
		}

		nextI, nextMaxI, maxI := i+1, nums[i]+i, 0

		for nextI <= nextMaxI {
			if nums[nextI]+nextI > maxI {
				maxI, i = nums[nextI]+nextI, nextI
			}

			nextI++
		}

		count++
	}

	return count
}
