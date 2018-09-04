package problem100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil ||
		p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	ps, qs := []*TreeNode{}, []*TreeNode{}

	for (p != nil && q != nil) || (len(ps) != 0 && len(qs) != 0) {

		if p != nil && q != nil {
			if p.Val != q.Val {
				return false
			}

			ps = append(ps, p)
			qs = append(qs, q)

			for p.Left != nil && q.Left != nil {
				p = p.Left
				q = q.Left

				if p.Val != q.Val {
					return false
				}

				ps = append(ps, p)
				qs = append(qs, q)
			}

			if p.Left != nil {
				return false
			}

			if q.Left != nil {
				return false
			}
		}

		if p == nil && q != nil ||
			p != nil && q == nil {
			return false
		}

		pNode := ps[len(ps)-1]
		qNode := qs[len(qs)-1]

		ps = ps[:len(ps)-1]
		qs = qs[:len(qs)-1]

		p = pNode.Right
		q = qNode.Right
	}

	return true
}
