package problem327

//   -2 5 -1
// 0 -2 3 2
// -2 0 2 3
var lo, up int
var temp []int

func countRangeSum(nums []int, lower int, upper int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	// 算出累计值
	// sum[i] 为 nums[:i] 累计值
	sums := make([]int, size+1)
	for i := 1; i < size+1; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	lo, up = lower, upper
	temp = make([]int, size+1)

	// 归并排序
	return mergeSort(sums)
}

func mergeSort(sums []int) int {
	size := len(sums)
	if size == 1 {
		return 0
	}

	mid := size / 2
	left := sums[:mid]
	right := sums[mid:]
	count := mergeSort(left) + mergeSort(right)

	sl, sr := 0, 0
	j, t := 0, 0
	rSize := len(right)

	for i := 0; i < mid; i++ {
		for sl < rSize && right[sl]-left[i] < lo {
			sl++
		}

		for sr < rSize && right[sr]-left[i] <= up {
			sr++
		}

		count += sr - sl

		for j < rSize && right[j] < left[i] {
			temp[t] = right[j]
			t++
			j++
		}

		temp[t] = left[i]
		t++
	}

	for j < rSize {
		temp[t] = right[j]
		t++
		j++
	}

	copy(sums, temp)

	return count
}
