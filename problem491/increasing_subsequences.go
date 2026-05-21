package problem491

func findSubsequences(nums []int) [][]int {
	size := len(nums)
	list := make([]int, 0, size)
	res := [][]int{}
	helper(0, nums, list, &res)

	return res
}

func helper(index int, nums []int, list []int, res *[][]int) {
	// 已经没有数可以加入list
	if index == len(nums) {
		if len(list) > 1 { //list最少有两个元素
			temp := make([]int, len(list))
			copy(temp, list)
			*res = append(*res, temp)
		}
		return
	}

	// 大于等于list最后一个数的，加入
	if len(list) == 0 || list[len(list)-1] <= nums[index] {
		list = append(list, nums[index])
		helper(index+1, nums, list, res)
		list = list[:len(list)-1]
	}

	// 跳过一个数
	if len(list) == 0 ||
		list[len(list)-1] != nums[index] { // list[4,6,7,7，7] 这种后面有几个数相等的情况，要跳到最后 [4,6]，避免重复
		helper(index+1, nums, list, res)
	}

}
