package problem274

func hIndex(citations []int) int {

}

func quickSort(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}

	i, j := 1, size-1
	base := nums[0]
	for i <= j {
		if nums[i] >= base {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}

		if nums[i] < base {
			nums[i], nums[0] = nums[i], nums[0]
			i++
		}
	}
}
