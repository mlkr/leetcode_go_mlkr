package problem463

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	grid [][]int
	ans  int
}{

	{
		[][]int{
			{1, 1, 1},
			{1, 0, 1},
		},
		12,
	},

	{
		[][]int{
			{0, 1, 1},
			{1, 1, 0},
		},
		10,
	},

	{
		[][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0},
		},
		16,
	},

	{
		[][]int{
			{1},
		},
		4,
	},

	{
		[][]int{
			{1},
			{0},
		},
		4,
	},

	{
		[][]int{
			{1, 0},
		},
		4,
	},
}

func Test_islandPerimeter(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, islandPerimeter(q.grid), "输入 %v", q.grid)
	}
}

func Test_islandPerimeter2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, islandPerimeter2(q.grid), "输入 %v", q.grid)
	}
}

func Benchmark_islandPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			islandPerimeter(q.grid)
		}
	}
}

func Benchmark_islandPerimeter2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			islandPerimeter2(q.grid)
		}
	}
}
