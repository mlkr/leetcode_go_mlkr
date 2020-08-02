package problem430

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	temp := root
	stack := []*Node{}
	for temp.Next != nil || temp.Child != nil || len(stack) > 0 {
		if temp.Child != nil {
			if temp.Next != nil {
				stack = append(stack, temp.Next)
			}

			temp.Next = temp.Child
			temp.Child.Prev = temp
			temp.Child = nil
		}

		if temp.Next == nil {
			temp.Next = stack[len(stack)-1]
			stack[len(stack)-1].Prev = temp
			stack = stack[:len(stack)-1]
		}

		temp = temp.Next
	}

	return root
}
