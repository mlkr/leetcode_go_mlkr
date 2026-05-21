package problem662

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	indexes := []int{0}        // 索引队列
	nodes := []*TreeNode{root} // 节点队列

	res := 0
	for len(nodes) > 0 {
		res = max(res, getWidth(indexes))

		size := len(indexes)
		upperIndexes := indexes[:size]
		upperNodes := nodes[:size]

		indexes = indexes[size:]
		nodes = nodes[size:]

		for i := 0; i < size; i++ {
			if upperNodes[i].Left != nil {
				indexes = append(indexes, upperIndexes[i]*2)
				nodes = append(nodes, upperNodes[i].Left)
			}

			if upperNodes[i].Right != nil {
				indexes = append(indexes, upperIndexes[i]*2+1)
				nodes = append(nodes, upperNodes[i].Right)
			}
		}
	}

	return res
}

// 给每层节点编号
// 每层的宽度 = 最后节点的编号 - 第一个节点的编号 + 1
func getWidth(indexes []int) int {
	size := len(indexes)
	return indexes[size-1] - indexes[0] + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
