package problem457

import (
	"unsafe"
)

func circularArrayLoop(nums []int) bool {
	size := len(nums)
	for i, n := range nums {
		nums[i] = n % size
	}

	// int 类型所占 bit 数减一
	bits := unsafe.Sizeof(size) - 1
	for i, n := range nums {

		// (n>>bits | 1) 决定mark的正负
		// 每个循环 mark 唯一
		mark := (i + size) * (n>>bits | 1)
		for n > -size && n < size && n != 0 {
			nums[i] = mark

			i = (i + n + size) % size
			n = nums[i]
			if n == mark {
				return true
			}

			if n*mark < 0 {
				// 方向改变
				break
			}
		}
	}

	return false
}
