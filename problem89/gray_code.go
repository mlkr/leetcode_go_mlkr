package problem89

func grayCode(n int) []int {
	return getGrayCode(n, 1, []int{0})
}

// 利用模拟法来产生格雷码的时候有一个镜面定理，
// 即当所有三位的格雷码可以由所有二位格雷码＋将二位所有格雷码从从最后一个到第一个依次最左边加１构成．
// 例如二位的格雷码依次是00, 01, 11, 10．
// 而三位的格雷码就是继承二位的格雷码＋逆序二位格雷码并在左边加１：
// 110, 111, 101, 100．
// 这样我们就可以得到所有的３位格雷为
// 00, 01, 11, 10, 110, 111, 101, 100.
func getGrayCode(n, base int, nums []int) []int {
	if n == 0 {
		return nums
	}

	length := len(nums)
	temp := make([]int, length)

	for i := range nums {
		temp[length-1-i] = nums[i] + base
	}

	return getGrayCode(n-1, base*2, append(nums, temp...))
}
