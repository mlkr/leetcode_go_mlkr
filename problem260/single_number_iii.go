package problem260

func singleNumber(nums []int) []int {
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		if !m[v] {
			m[v] = true
		} else {
			delete(m, v)
		}
	}

	res := []int{}
	for k := range m {
		res = append(res, k)
	}

	return res
}

func singleNumber2(nums []int) []int {
	quickSort(nums)

	res := []int{}
	size := len(nums)
	for i := 1; i < size-1; i++ {
		if nums[i-1] != nums[i] && nums[i] != nums[i+1] {
			res = append(res, nums[i])
		}
	}

	if nums[0] != nums[1] {
		res = append(res, nums[0])
	}

	if nums[size-1] != nums[size-2] {
		res = append(res, nums[size-1])
	}

	return res
}

// 把所有的数字进行一次异或，得到的是只出现了一次的两个数字的异或.
// 这两个数字不等，因此他们的二进制必定至少1位不同，即异或结果中为1的那位（一个数字的该位为1，另个数字的该位为0）.
// 找出从右向左的第一个不同的位置（异或值为1的位置），给mask在该位置设置成1，mask的其余位置是0.
// mask存在的意义在于我们能通过该位置来分辨出两个只出现了一次的数字.
// 然后技巧性的来了：再进行一次异或操作.
// 每个数字都跟mask相与。通过与的结果为0和为1，即可区分出两个数字.
// 最佳
func singleNumber3(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}

	mask := xor & -xor
	a, b := 0, 0
	for _, v := range nums {
		if mask&v == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}

	return []int{a, b}
}

func quickSort(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}

	base := nums[0]
	left, right := 0, size-1
	for i := 1; i <= right; {
		if base < nums[i] {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			nums[i], nums[left] = nums[left], nums[i]
			i++
			left++
		}
	}

	quickSort(nums[:right])
	quickSort(nums[right+1:])
}
