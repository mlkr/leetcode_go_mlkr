package problem553

import "strconv"

func optimalDivision(nums []int) string {
	// n1/n2/n3/n4/...nt = n1/(n2*n3*n4*...nt)
	// n1/(n2/n3/n4/...nt) = (n1*n3*n4*...nt)/n2
	// n2怎么样都会被除，除了n2之外，后面的数都可以转变为乘积！这样的结果最大
	size := len(nums)
	if size == 1 {
		return strconv.Itoa(nums[0])
	}

	if size == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}

	res := ""
	for i := 0; i < size; i++ {
		if i == 0 {
			res += strconv.Itoa(nums[i]) + "/("
			continue
		}

		if i == size-1 {
			res += strconv.Itoa(nums[i]) + ")"
			break
		}

		res += strconv.Itoa(nums[i]) + "/"

	}

	return res
}
