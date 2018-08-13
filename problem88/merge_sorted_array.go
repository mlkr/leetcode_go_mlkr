package problem88

import "math/rand"

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	nums1 = append(nums1, nums2...)

	quickSort(nums1)
}

func quickSort(nums []int) {
	numsLen := len(nums)
	if numsLen <= 1 {
		return
	}

	p := rand.Intn(numsLen)
	i := part(nums, p)

	quickSort(nums[:i])
	quickSort(nums[i+1:])
}

func part(nums []int, part int) int {
	numsLen := len(nums)
	nums[0], nums[part] = nums[part], nums[0]

	i, j := 1, numsLen-1
	for i < j {
		for nums[i] <= nums[0] && i < numsLen-1 {
			i++
		}

		for nums[j] > nums[0] && j > 0 {
			j--
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	// 类似 [5,6] i,j = 1,1 这种情况
	if nums[0] > nums[j] {
		nums[0], nums[j] = nums[j], nums[0]
	}

	return j
}

// 另一解法
func merge2(nums1 []int, m int, nums2 []int, n int) {
	nums1Len := len(nums1)
	temp := make([]int, nums1Len)
	copy(temp, nums1)

	j, k := 0, 0
	for i := 0; i < nums1Len; i++ {
		if k >= n {
			nums1[i] = temp[j]
			j++
			continue
		}

		if j >= m {
			nums1[i] = nums2[k]
			k++
			continue
		}

		if temp[j] < nums2[k] {
			nums1[i] = temp[j]
			j++
		} else {
			nums1[i] = nums2[k]
			k++
		}
	}
}
