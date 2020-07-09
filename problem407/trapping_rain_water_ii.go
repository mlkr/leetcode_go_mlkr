package problem407

import (
	"container/heap"
)

func trapRainWater(heightMap [][]int) int {
	rows := len(heightMap)
	cols := len(heightMap[0])

	if rows < 3 || cols < 3 {
		return 0
	}

	cells := make([]cell, 0, rows*cols)
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// 添加四周的 cell
	for i := 0; i < rows; i++ {
		visited[i][0] = true
		visited[i][cols-1] = true
		cells = append(cells, cell{i, 0, heightMap[i][0]})
		cells = append(cells, cell{i, cols - 1, heightMap[i][cols-1]})
	}

	for i := 1; i < cols-1; i++ {
		visited[0][i] = true
		visited[rows-1][i] = true
		cells = append(cells, cell{0, i, heightMap[0][i]})
		cells = append(cells, cell{rows - 1, i, heightMap[rows-1][i]})

	}

	ch := cellHeap(cells)
	heap.Init(&ch)

	dirs := [][]int{[]int{-1, 0}, []int{0, -1}, []int{1, 0}, []int{0, 1}}

	vol := 0
	for ch.Len() > 0 {
		c := heap.Pop(&ch).(cell)
		for _, dir := range dirs {
			i := c.row + dir[0]
			j := c.col + dir[1]

			if 0 <= i && i < rows && 0 <= j && j < cols && !visited[i][j] {
				visited[i][j] = true
				vol += max(0, c.height-heightMap[i][j])
				heap.Push(&ch, cell{i, j, max(c.height, heightMap[i][j])})
			}
		}
	}

	return vol
}

type cell struct {
	row, col, height int
}

type cellHeap []cell

func (this cellHeap) Len() int           { return len(this) }
func (this cellHeap) Less(i, j int) bool { return this[i].height < this[j].height }
func (this cellHeap) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }

func (this *cellHeap) Push(x interface{}) {
	*this = append(*this, x.(cell))
}

func (this *cellHeap) Pop() interface{} {
	res := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
