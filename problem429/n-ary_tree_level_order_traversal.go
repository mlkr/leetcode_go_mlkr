package problem429

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	cur_level := []*Node{}
	cur_level = append(cur_level, root)

	for len(cur_level) > 0 {
		temp := []int{}
		next_level := []*Node{}

		for _, node := range cur_level {
			temp = append(temp, node.Val)

			for _, children := range node.Children {
				next_level = append(next_level, children)
			}
		}

		cur_level = next_level

		res = append(res, temp)
	}

	return res
}
