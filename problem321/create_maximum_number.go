package problem321

// http://zhuixin8.com/2016/10/02/leetcode-321/
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	size1, size2 := len(nums1), len(nums2)

	res := make([]int, k)
	for i := 0; i <= size1 && i <= k; i++ {
		if k-i > size2 {
			continue
		}

		temp := combine(maxArray(nums1, i), maxArray(nums2, k-i))
		if isBig(temp, res) {
			copy(res, temp)
		}

	}

	return res
}

// 取 nums 中最大的 k 个值, 保持顺序
func maxArray(nums []int, k int) []int {
	size := len(nums)

	if k == 0 {
		return []int{}
	}

	if k == size {
		return nums
	}

	res := make([]int, 0, k)
	for i := 0; i < size; i++ {
		for len(res) > 0 && res[len(res)-1] < nums[i] && size-i > k-len(res) {
			res = res[:len(res)-1]
		}

		if len(res) < k {
			res = append(res, nums[i])
		}
	}

	return res
}

// 合并两个数组, 使合并后的数组最大
func combine(nums1, nums2 []int) []int {
	size1, size2 := len(nums1), len(nums2)
	res := make([]int, size1+size2)

	i, j, r := 0, 0, 0
	for ; i < size1 && j < size2; r++ {
		if nums1[i] > nums2[j] || isBig(nums1[i:], nums2[j:]) {
			res[r] = nums1[i]
			i++
		} else {
			res[r] = nums2[j]
			j++
		}
	}

	res = append(res[:r], nums1[i:]...)
	res = append(res[:len(res)], nums2[j:]...)

	return res
}

func isBig(nums1, nums2 []int) bool {
	size2 := len(nums2)

	for i := range nums1 {
		if i == size2 || nums1[i] > nums2[i] {
			return true
		}

		if nums1[i] < nums2[i] {
			return false
		}
	}

	return false
}
