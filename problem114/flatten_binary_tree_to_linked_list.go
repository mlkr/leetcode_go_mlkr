package problem114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 按先序遍历的顺序展平(不是从小到大的顺序)
	roots := preOrderNode(root)
	// quickSort(roots)
	root = makeTree(roots)
}

func makeTree(roots []*TreeNode) *TreeNode {
	root := roots[0]
	temp := root

	rootsLen := len(roots)
	for i := 1; i < rootsLen; i++ {
		root.Left = nil
		root.Right = roots[i]

		root = root.Right
	}

	root.Left = nil
	root.Right = nil

	return temp
}

// func quickSort(roots []*TreeNode) {
// 	if len(roots) == 0 {
// 		return
// 	}

// 	part(roots)
// }

// func part(roots []*TreeNode) {
// 	rootsLen := len(roots)
// 	m := rand.Intn(rootsLen)
// 	v := roots[m].Val

// 	roots[m], roots[0] = roots[0], roots[m]

// 	i, j := 1, rootsLen-1
// 	for i < j {
// 		for roots[i].Val <= v && i < rootsLen-1 {
// 			i++
// 		}

// 		for roots[j].Val > v && j > 0 {
// 			j--
// 		}

// 		if i < j {
// 			roots[i], roots[j] = roots[j], roots[i]
// 		}
// 	}

// 	if roots[0].Val > roots[j].Val {
// 		roots[0], roots[j] = roots[j], roots[0]
// 	}

// 	quickSort(roots[:j])
// 	quickSort(roots[j+1:])
// }

func preOrderNode(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	roots := []*TreeNode{}

	if root == nil {
		return res
	}

	for root != nil || len(roots) != 0 {
		if root != nil {
			roots = append(roots, root)
			res = append(res, root)

			root = root.Left
		} else {
			root = roots[len(roots)-1]
			roots = roots[:len(roots)-1]

			root = root.Right
		}

	}

	return res
}
