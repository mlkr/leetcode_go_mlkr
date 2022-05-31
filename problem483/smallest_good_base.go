package problem483

import (
	"math"
	"strconv"
)

func smallestGoodBase(n string) string {
	num, _ := strconv.Atoi(n)

	// k^m+k^(m-1)+...+k+1 = n
	// k小 m大，k=2时， m最大 mMax = log2 n
	mMax := int(math.Log2(float64(num)))

	for m := mMax; m > 1; m-- {
		// (k+1)^m 展开可得 (k+1)^m > k^m+k^(m-1)+...+k+1 = n
		// (k+1)^m > n > k^m  m>1时成立
		// k+1 > n^(1/m) > k
		// k = int(n^(1/m)) 取整数时，数的大小向下（向下取整）
		k := int(math.Pow(float64(num), 1.0/float64(m)))

		if isFound(k, m, num) {
			return strconv.Itoa(k)
		}
	}

	// 特别注意当 m=1时 (k+1)^m = k^m+1 = n 的情况
	return strconv.Itoa(num - 1)
}

func isFound(k, m, num int) bool {
	sum := 1
	for m > 0 {
		sum = sum*k + 1
		m--
	}

	return sum == num
}
