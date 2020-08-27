package problem440

// 参看
// https://leetcode.flowerplayer.com/2019/04/10/leetcode-440-k-th-smallest-in-lexicographical-order%E8%A7%A3%E9%A2%98%E6%80%9D%E8%B7%AF%E5%88%86%E6%9E%90/
func findKthNumber(n int, k int) int {
	res := 1
	k--

	for k > 0 {
		begin, end := res, res+1
		count := 0
		for begin <= n {
			count += min(n+1, end) - begin
			begin *= 10
			end *= 10
		}

		if count <= k {
			k -= count
			res++
		} else {
			k--
			res *= 10
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
