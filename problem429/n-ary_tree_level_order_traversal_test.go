package problem429

import (
	"testing"
)

// var questions = []struct {
// 	nums []int
// 	ans  [][]int
// }{
// 	{
// 		[]int{1, nil, 3, 2, 4, nil, 5, 6},
// 		[][]int{
// 			{1},
// 			{3, 2, 4},
// 			{5, 6},
// 		},
// 	},
// }

func Test_levelOrder(t *testing.T) {
	// ast := assert.New(t)

	// for _, q := range questions {

	// }
}

// func nums2Tree(nums []int) *Node {
// 	root := &Node{Val: nums[0]}

// 	queue := []*Node{}
// 	queue = append(queue, &Node{Val: nums[0]})

// 	for i := 2; i < len(nums); i++ {
// 		children := []*Node{}
// 		for j := i; j < len(nums); j++ {
// 			if nums[j] == nil {
// 				break
// 			}

// 			node := &Node{Val: nums[j]}
// 			children = append(children, node)
// 			queue = append(queue, node)
// 		}

// 		temp := queue[0]
// 		queue = queue[1:]

// 		if len(children) != 0 {
// 			temp.Children = children
// 		}

// 	}

// 	return root
// }
