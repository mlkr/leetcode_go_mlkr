package problem96

func numTrees(n int) int {
	if n == 0 {
		return 1
	}

	if n <= 2 {
		return n
	}

	if n == 3 {
		return 5
	}

	// 如果一个搜索二叉树有 3 个节点, 节点的分布可能如下
	// n = 3
	//  root     root     root
	//  /  \     /  \     /  \
	// 0    2   1    1   2    0
	// 上面的数字表示子树的节点数
	// n节点的一个可能的分布情况的二叉搜索树数目 = 左子树所有二叉搜索树数目 * 右子树所有二叉搜索树数目
	// n节点的二叉搜索树数目 = n节点的一个可能的分布情况的二叉搜索树数目 + n节点的一个可能的分布情况的二叉搜索树数目 +....
	count := 0

	// 情况重复, 只算一半
	// 例
	// n = 3
	//  root       root
	//  /  \       /  \
	// 0    2     2    0
	// 重复
	for i := 0; i < n/2; i++ {
		count += numTrees(i) * numTrees(n-1-i)
	}
	count *= 2

	// 如果 n 为单数, 算出中间的节点分布的二叉搜索树数目
	// 例
	// n = 3
	//  root
	//  /  \
	// 1    1
	if n%2 == 1 {
		midCount := numTrees(n / 2)
		midCount *= midCount
		count += midCount
	}

	return count
}
