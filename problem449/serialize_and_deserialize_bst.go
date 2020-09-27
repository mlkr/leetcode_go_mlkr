package problem449

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 The tree can be built using preorder or postorder traversal and inorder traversal.
 since the tree is BST, we already know that inorder traversal can give sorted order.
serialize tree using preorder or postorder traversal
deserialize using preoder/postorder and inorder traversal
*/
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	nodes := []rune{}

	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		nodes = append(nodes, rune(root.Val))
		preOrder(root.Left)
		preOrder(root.Right)
	}

	preOrder(root)

	return string(nodes)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	nodes := []rune{}
	for _, str := range data {
		nodes = append(nodes, rune(str))
	}

	index := 0
	var buildTree func(min, max int) *TreeNode
	buildTree = func(min, max int) *TreeNode {
		if index == len(nodes) {
			return nil
		}

		val := int(nodes[index])
		if val > max || val < min {
			return nil
		}

		index++
		root := &TreeNode{Val: val}
		root.Left = buildTree(min, val)
		root.Right = buildTree(val, max)

		return root
	}

	return buildTree(math.MinInt32, math.MaxInt32)
}
