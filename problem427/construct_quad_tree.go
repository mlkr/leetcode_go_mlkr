package problem427

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return helper(0, 0, len(grid), grid)
}

func helper(i, j, n int, grid [][]int) *Node {
	if n == 0 {
		return nil
	}

	if isLeaf(i, j, n, grid) {
		return &Node{grid[i][j] == 1, true, nil, nil, nil, nil}
	}

	n >>= 1
	return &Node{
		true,
		false,
		helper(i, j, n, grid),
		helper(i, j+n, n, grid),
		helper(i+n, j, n, grid),
		helper(i+n, j+n, n, grid),
	}
}

func isLeaf(i, j, n int, grid [][]int) bool {
	v := grid[i][j]

	for x := i; x < i+n; x++ {
		for y := j; y < j+n; y++ {
			if grid[x][y] != v {
				return false
			}
		}
	}

	return true
}
