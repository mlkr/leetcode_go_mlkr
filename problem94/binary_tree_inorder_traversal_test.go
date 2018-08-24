package problem94

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 两个序列决定一个二叉树
type para struct {
	// 前序
	pre []int

	// 中序
	in []int
}

type ans struct {
	// 中序
	in []int
}

type question struct {
	para
	ans
}

func TestInorderTraversal(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
			},
			ans{
				[]int{1, 3, 2},
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(inorderTraversal(prein2Tree(para.pre, para.in)), ans.in, "输入%v", para)
		ast.Equal(inorderTraversal2(prein2Tree(para.pre, para.in)), ans.in, "输入%v", para)
	}
}

// 前序,中序 转二叉树
func prein2Tree(pre, in []int) *TreeNode {
	preLen, inLen := len(pre), len(in)
	if preLen != inLen {
		panic("两个序列不相等!")
	}

	if preLen == 0 {
		return nil
	}

	root := &TreeNode{
		Val: pre[0],
	}

	if preLen == 1 {
		return root
	}

	idx := indexOf(in, pre[0])

	// idx 表明 root 节点在 in 中的位置, 还表明左子树在 pre 中的位置
	// 例:
	//  (DLR) pre [1,2,3], (LDR) in [2,1,3], root 节点为 1,
	//  在 in 中的位置 idx = 1, 表明 in 中在 root 前还有 1 num 作为它的左子树,
	//  所以在 pre 中去除 root 节点后还有 1 num 划为左子树: pre[1:idx+1],
	//  剩下的为右子树: pre[idx+1:]
	root.Left = prein2Tree(pre[1:idx+1], in[:idx])
	root.Right = prein2Tree(pre[idx+1:], in[idx+1:])

	return root
}

func indexOf(nums []int, n int) int {
	for k, v := range nums {
		if v == n {
			return k
		}
	}

	return -1
}
