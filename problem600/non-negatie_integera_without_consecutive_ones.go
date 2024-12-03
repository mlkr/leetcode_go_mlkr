package problem600

func findIntegers(num int) int {
	// num 转二进制数 】逆序
	binary := []byte{}
	for num > 0 {
		binary = append(binary, '0'+byte(num%2))
		num /= 2
	}

	// 将其转换为节点值为 0,1 的二叉树
	n := len(binary)

	// a[i] 以0结尾高度为i 的二叉树的所有路径
	// b[i] 以1结尾高度为i 的二叉树的所有路径
	a, b := make([]int, n), make([]int, n)
	a[0] = 1
	b[0] = 1

	for i := 1; i < n; i++ {
		a[i] = a[i-1] + b[i-1]
		b[i] = a[i-1]
	}

	// 以0结尾高度为i 的二叉树的所有路径 加上
	// 以1结尾高度为i 的二叉树的所有路径 等于所有不含连续1的非负整数 数量
	res := a[n-1] + b[n-1]

	// 减去 >num 的数量
	// 请构造 0 1 二叉树帮助理解
	// https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/solutions/990719/bu-han-lian-xu-1de-fei-fu-zheng-shu-by-l-9l86/
	// 例 num = 4
	//    0        1
	//  0   1    0   1
	// 0 1 0 1  0 1 0 1
	for i := n - 2; i >= 0; i-- {
		if binary[i] == '1' && binary[i+1] == '1' {
			break
		}

		if binary[i] == '0' && binary[i+1] == '0' {
			res -= b[i]
		}
	}

	return res
}
