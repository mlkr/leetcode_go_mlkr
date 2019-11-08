package problem334

func increasingTriplet(nums []int) bool {
	first := 1<<63 - 1
	second := 1<<63 - 1

	for _, num := range nums {
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}

	return false
}
