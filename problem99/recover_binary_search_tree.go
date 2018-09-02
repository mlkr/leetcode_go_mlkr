package problem99

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	// 得中序序列
	in := getInOrder(root)
	inLen := len(in)

	// 得要交换的两个 TreeNode
	swap := []*TreeNode{}
	for i := 0; i < inLen-1; i++ {
		if in[i].Val > in[i+1].Val {
			if len(swap) == 0 {
				swap = append(swap, in[i])
			} else if len(swap) == 1 {
				swap = append(swap, in[i+1])
			}

			// 3,2,1 这种情况
			if i+2 <= inLen-1 && in[i].Val > in[i+1].Val && in[i+1].Val > in[i+2].Val {
				swap = append(swap, in[i+2])
				break
			}

			continue
		}
	}

	// 1,3,2,4 这种
	if len(swap) == 1 {
		idx := indexOfTreeNode(swap[0], in)
		swap = append(swap, in[idx+1])
	}

	swap[0].Val, swap[1].Val = swap[1].Val, swap[0].Val

}

func indexOfTreeNode(n *TreeNode, nums []*TreeNode) int {
	res := -1
	for k, v := range nums {
		if v.Val == n.Val {
			return k
		}
	}

	return res
}

// 得二叉树的中序序列
func getInOrder(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	roots := []*TreeNode{}

	for root != nil || len(roots) != 0 {
		if root != nil {
			roots = append(roots, root)

			for root.Left != nil {
				root = root.Left
				roots = append(roots, root)
			}
		}

		node := roots[len(roots)-1]
		roots = roots[:len(roots)-1]

		res = append(res, node)

		root = node.Right
	}

	return res
}
