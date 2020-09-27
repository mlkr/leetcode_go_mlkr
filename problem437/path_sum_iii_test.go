package problem437

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Int struct {
	num   int
	isNil bool
}

func getInt(num int) Int {
	return Int{num, false}
}

func getNilInt() Int {
	return Int{0, true}
}

var questions = []struct {
	ints     []Int
	sum, ans int
}{
	{
		[]Int{getInt(0), getInt(1), getInt(1)},
		1,
		4,
	},

	{
		[]Int{getInt(10), getInt(5), getInt(-3),
			getInt(3), getInt(2), getNilInt(),
			getInt(11), getInt(3), getInt(-2),
			getNilInt(), getInt(1),
		},
		8,
		3,
	},

	// prein2Tree 不足以构建二叉树
	// 3 3 -2	dlr
	// 3 3 -2	ldr
	// 3 -2 3	lrd
	// 多一个后序遍历 postorder, 如果 全是3 也不行
	// root = [10,5,-3,3,2,null,11,3,-2,null,1] 这种类型的才行
	// https://stackoom.com/question/3Y4Y2/%E5%9C%A8golang%E4%B8%AD%E5%8C%85%E5%90%AB%E6%95%B4%E6%95%B0%E5%92%8Cnil%E5%80%BC%E7%9A%84%E6%95%B0%E7%BB%84
}

func Test_pathSum(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := nodes2Tree(q.ints)
		ast.Equal(q.ans, pathSum(root, q.sum), "%v", q.sum)
	}
}

// 和堆排序有点像
//  0 1  2 3 4 5    6  7  8  9   10
// 10,5,-3,3,2,null,11,3,-2,null,1
// 2n+1
// 2(n+1) + 1 = 2n + 3
func nodes2Tree(ints []Int) *TreeNode {
	root := &TreeNode{Val: ints[0].num}

	helper(ints, 0, root)

	return root
}

func helper(ints []Int, n int, parent *TreeNode) {
	if parent == nil {
		return
	}

	if 2*n+1 > len(ints)-1 {
		return
	}

	if !ints[2*n+1].isNil {
		parent.Left = &TreeNode{Val: ints[2*n+1].num}

	}

	if 2*n+2 > len(ints)-1 {
		return
	}

	if !ints[2*n+2].isNil {
		parent.Right = &TreeNode{Val: ints[2*n+2].num}
	}

	helper(ints, 2*n+1, parent.Left)
	helper(ints, 2*n+2, parent.Right)
}
