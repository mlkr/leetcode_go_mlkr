package problem228

import (
	"fmt"
	"strconv"
)

type Bound struct {
	min, max int
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	bound := Bound{
		min: nums[0],
		max: nums[0],
	}

	res := make([]Bound, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i]-bound.max == 1 {
			bound.max = nums[i]
		} else {
			res = append(res, bound)
			bound = Bound{
				min: nums[i],
				max: nums[i],
			}
		}
	}
	res = append(res, bound)

	return printStrings(res)
}

func printStrings(bounds []Bound) []string {
	res := make([]string, 0, len(bounds))
	for _, bound := range bounds {
		if bound.min == bound.max {
			res = append(res, strconv.Itoa(bound.min))
		} else {
			res = append(res, strconv.Itoa(bound.min)+"->"+strconv.Itoa(bound.max))
		}
	}

	return res
}

// 另一解法
func summaryRanges2(nums []int) []string {
	size := len(nums)
	if size == 0 {
		return nil
	}

	res := make([]string, 0)
	str := ""
	begin := nums[0]
	for i := 0; i < size; i++ {
		if i == size-1 || nums[i]+1 != nums[i+1] {
			if begin == nums[i] {
				str = fmt.Sprintf("%d", begin)
			} else {
				str = fmt.Sprintf("%d->%d", begin, nums[i])
			}

			if i+1 < size {
				begin = nums[i+1]
			}

			res = append(res, str)
		}
	}

	return res
}
