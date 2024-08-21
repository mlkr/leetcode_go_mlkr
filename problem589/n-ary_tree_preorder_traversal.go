package problem589

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}

	res = append(res, root.Val)

	for _, c := range root.Children {
		res = append(res, preorder(c)...)
	}

	return res
}
