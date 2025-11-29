package problem658

func findClosestElements(arr []int, k int, x int) []int {
	index := getIndex(arr, x)
	if index == -1 {
		return arr[:k]
	}

	if index == len(arr) {
		return arr[len(arr)-k : len(arr)]
	}

	left := index
	right := index + 1
	for right-left < k {
		if left == 0 {
			right++
			continue
		}

		if right == len(arr) {
			left--
			continue
		}

		if x-arr[left-1] <= arr[right]-x {
			left--
			continue
		}

		right++
	}

	return arr[left:right]
}

func getIndex(arr []int, x int) int {
	if x < arr[0] {
		return -1
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return i
		}

		if arr[i] > x && (i-1 >= 0 && x > arr[i-1]) {
			if x-arr[i-1] <= arr[i]-x {
				return i - 1
			}

			return i
		}
	}

	return len(arr)
}
