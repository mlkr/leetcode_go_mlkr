package problem113

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}

	var runPath func(root *TreeNode, count int, path []int)
	runPath = func(root *TreeNode, count int, path []int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			count += root.Val
			if count == sum {
				temp := make([]int, len(path)+1)
				temp[len(path)] = root.Val
				copy(temp[:len(path)], path)

				res = append(res, temp)
			}

			return
		}

		runPath(root.Left, count+root.Val, append(path, root.Val))
		runPath(root.Right, count+root.Val, append(path, root.Val))
	}

	runPath(root, 0, []int{})

	return res
}
