package problem463

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func islandPerimeter(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}

			res += 4

			for k := 0; k < 4; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if 0 <= x && x < rows && 0 <= y && y < cols &&
					grid[x][y] == 1 {
					res--
				}
			}
		}
	}

	return res
}

// 解法二
const (
	WATER = iota
	LAND
)

func islandPerimeter2(grid [][]int) int {
	dirs := []int{1, 0, -1, 0, 1}
	rows, cols := len(grid), len(grid[0])

	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}

			count := 0
			for k := 0; k < 4; k++ {
				x := i + dirs[k]
				y := j + dirs[k+1]

				if x < 0 || x >= rows || y < 0 || y >= cols {
					count++
					continue
				}

				if grid[x][y] == WATER {
					count++
				}
			}

			res += count
		}
	}

	return res
}
