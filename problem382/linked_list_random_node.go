package problem382

import "math/rand"

// 参看https://leetcode-cn.com/problems/linked-list-random-node/solution/xu-shui-chi-chou-yang-suan-fa-by-jackwener/
type ListNode struct{
	Val int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{head}
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	node := this.head
	val := this.head.Val
	count := 1

	for node.Next != nil {
		node = node.Next
		count++

		if rand.Float64() < (1.0/float64(count)) {
			val = node.Val
		}

	}

	return val
}