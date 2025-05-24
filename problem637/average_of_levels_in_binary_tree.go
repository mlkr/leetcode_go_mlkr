package problem637

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	nodes := make([]*TreeNode, 1, 1024)
	res := make([]float64, 0, 14)

	nodes[0] = root
	for len(nodes) > 0 {

		size := len(nodes)
		sum := 0
		for i := 0; i < size; i++ {
			sum += nodes[i].Val

			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}

			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}

		res = append(res, float64(sum)/float64(size))
		nodes = nodes[size:]
	}

	return res
}
