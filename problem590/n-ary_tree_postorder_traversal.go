package problem590

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}

	if root.Children != nil {
		for _, c := range root.Children {
			res = append(res, postorder(c)...)
		}
	}

	res = append(res, root.Val)
	return res
}
