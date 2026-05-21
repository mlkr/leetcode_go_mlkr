package problem485

func findMaxConsecutiveOnes(nums []int) int {
	currentMax := 0
	maxOnes := 0

	for _, n := range nums {
		if n == 1 {
			currentMax++
			if currentMax > maxOnes {
				maxOnes = currentMax
			}
			continue
		}

		currentMax = 0
	}

	return maxOnes
}
