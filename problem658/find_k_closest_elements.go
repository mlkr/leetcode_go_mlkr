package problem658

func findClosestElements(arr []int, k int, x int) []int {
	idx := getIndex(arr, x)
	if idx == 0 {
		return arr[:k]
	}

	size := len(arr)
	if idx == size-1 {
		return arr[size-k : size]
	}

	left := idx
	right := idx + 1
	for right-left < k {
		if left == 0 {
			right++
			continue
		}

		if right == size {
			left--
			continue
		}

		if x-arr[left-1] <= arr[right]-x {
			left--
		} else {
			right++
		}

	}

	return arr[left:right]
}

func getIndex(arr []int, x int) int {
	if arr[0] >= x {
		return 0
	}

	size := len(arr)
	for i := 1; i < size; i++ {
		if arr[i] == x {
			return i
		}

		if arr[i] > x {
			if x-arr[i-1] <= arr[i]-x {
				return i - 1
			}

			return i
		}
	}

	return size - 1
}
