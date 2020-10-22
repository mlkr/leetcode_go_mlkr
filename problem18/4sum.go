package problem18

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, len(nums)-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]

				switch {
				case sum > target:
					r--
				case sum < target:
					l++
				default:
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l, r = next(nums, l, r)
				}

			}
		}
	}

	return res
}

func next(nums []int, l, r int) (int, int) {
loop:
	for l < r {
		switch {
		case nums[l+1] == nums[l]:
			l++
		case nums[r-1] == nums[r]:
			r--
		default:
			l++
			r--
			break loop
		}
	}

	return l, r
}
