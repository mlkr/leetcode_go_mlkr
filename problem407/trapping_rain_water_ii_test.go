package problem407

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	heightMap [][]int
	ans       int
}{
	{
		[][]int{
			[]int{12, 13, 1, 12},
			[]int{13, 4, 13, 12},
			[]int{13, 8, 10, 12},
			[]int{12, 13, 12, 12},
			[]int{13, 13, 13, 13},
		},
		14,
	},

	{
		[][]int{
			[]int{1, 4, 3, 1, 3, 2},
			[]int{3, 2, 1, 3, 2, 4},
			[]int{2, 3, 3, 2, 3, 1},
		},
		4,
	},

	{
		[][]int{
			[]int{1},
		},
		0,
	},
}

func Test_trapRainWater(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, trapRainWater(q.heightMap), "输入 %v", q.heightMap)
	}
}

func Benchmark_trapRainWater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			trapRainWater(q.heightMap)
		}
	}
}
