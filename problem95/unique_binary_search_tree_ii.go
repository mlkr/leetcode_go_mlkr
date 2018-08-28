package problem95

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	var res []*TreeNode
	if n <= 0 {
		return res
	}

	// 所有的二叉搜索树的中序遍历都是一样的
	in := make([]int, n)
	for i := 1; i <= n; i++ {
		in[i-1] = i
	}

	// 根据中序遍历得所有可能的前序遍历
	pres := getPre(in)

	// 前序遍历+中序遍历生成二叉树
	for _, pre := range pres {
		root := prein2Tree(pre, in)
		res = append(res, root)
	}

	return res
}

// 根据中序遍历得所有可能的前序遍历
func getPre(in []int) [][]int {
	inLen := len(in)
	if inLen <= 1 {
		return [][]int{in}
	}

	if inLen == 2 {
		return [][]int{
			[]int{in[0], in[1]},
			[]int{in[1], in[0]},
		}
	}

	res := [][]int{}
	for i := 0; i < inLen; i++ {
		// in[i] 作为 root
		// 左子树的前序遍历
		lt := getPre(in[:i])
		// 右子树的前序遍历
		rt := getPre(in[i+1:])

		for _, l := range lt {
			for _, r := range rt {
				temp := make([]int, 1, inLen)
				temp[0] = in[i]
				temp = append(temp, l...)
				temp = append(temp, r...)

				res = append(res, temp)

			}
		}
	}

	return res
}

// 前序遍历+中序遍历生成二叉树
func prein2Tree(pre, in []int) *TreeNode {
	preLen, inLen := len(pre), len(in)
	if preLen != inLen {
		panic("前序,中序序列长度不一样!")
	}

	if preLen < 1 {
		return nil
	}

	root := &TreeNode{
		Val: pre[0],
	}

	if preLen == 1 {
		return root
	}

	// idx 表明 root 节点在 in 中的位置, 还表明左子树在 pre 中的位置
	// 例:
	//  (DLR) pre [1,2,3], (LDR) in [2,1,3], root 节点为 1,
	//  在 in 中的位置 idx = 1, 表明 in 中在 root 前还有 1 num 作为它的左子树,
	//  所以在 pre 中去除 root 节点后还有 1 num 划为左子树: pre[1:idx+1],
	//  剩下的为右子树: pre[idx+1:]
	idx := indexOf(pre[0], in)
	root.Left = prein2Tree(pre[1:idx+1], in[:idx])
	root.Right = prein2Tree(pre[idx+1:], in[idx+1:])

	return root
}

func indexOf(num int, nums []int) int {
	for idx, n := range nums {
		if n == num {
			return idx
		}
	}

	msg := fmt.Sprintf("%d 在 %v 中没有找到\n", num, nums)
	panic(msg)
}
