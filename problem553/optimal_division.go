package problem553

import "strconv"

func optimalDivision(nums []int) string {
	// n1/n2/n3/n4/...nt = n1/(n2*n3*n4*...nt)
	// n1/(n2/n3/n4/...nt) = (n1*n3*n4*...nt)/n2
	// n2怎么样都会被除，除了n2之外，后面的数都可以转变为乘积！这样的结果最大
	size := len(nums)
	strs := make([]string, size)
	for i := 0; i < size; i++ {
		strs[i] = strconv.Itoa(nums[i])
	}

	if size == 1 {
		return strs[0]
	}

	if size == 2 {
		return strs[0] + "/" + strs[1]
	}

	res := ""
	for i := 0; i < size; i++ {
		if i != 0 {
			res += "/"
		}

		if i == 1 {
			res += "("
		}

		res += strs[i]

		if i == size-1 {
			res += ")"
		}
	}

	return res
}
