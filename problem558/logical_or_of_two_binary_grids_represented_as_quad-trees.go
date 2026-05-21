package problem558

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func intersect(q1 *Node, q2 *Node) *Node {
	if q1.IsLeaf {
		if q1.Val {
			return q1
		}

		return q2
	}

	if q2.IsLeaf {
		if q2.Val {
			return q2
		}

		return q1
	}

	tl := intersect(q1.TopLeft, q2.TopLeft)
	tr := intersect(q1.TopRight, q2.TopRight)
	bl := intersect(q1.BottomLeft, q2.BottomLeft)
	br := intersect(q1.BottomRight, q2.BottomRight)

	a, e := tl.IsLeaf, tl.Val
	b, f := tr.IsLeaf, tr.Val
	c, g := bl.IsLeaf, bl.Val
	d, h := br.IsLeaf, br.Val

	if a && b && c && d && (e == f && f == g && g == h) {
		return &Node{a, e, nil, nil, nil, nil}
	}

	return &Node{false, false, tl, tr, bl, br}
}
