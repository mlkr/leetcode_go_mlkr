package problem204

// 参看 埃拉托斯特尼筛法
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	// 小于 n 的数,有一半是偶数, 偶数不是素数(除了2)
	count := n / 2

	// 合数, 素数的反义
	isComposite := make([]bool, n)

	// 检查每个奇数
	for i := 3; i*i < n; i += 2 {
		if isComposite[i] {
			continue
		}

		// j += 2 * i 保证 j 都是奇数
		for j := i * i; j < n; j += 2 * i {
			if !isComposite[j] {
				isComposite[j] = true
				count--
			}
		}
	}

	return count
}
