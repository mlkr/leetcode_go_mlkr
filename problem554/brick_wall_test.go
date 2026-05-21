package problem554

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	wall [][]int
	ans  int
}{
	{
		[][]int{
			[]int{1, 2, 2, 1},
			[]int{3, 1, 2},
			[]int{1, 3, 2},
			[]int{2, 4},
			[]int{3, 1, 2},
			[]int{1, 3, 1, 1},
		},
		2,
	},

	{
		[][]int{
			[]int{1},
			[]int{1},
			[]int{1},
		},
		3,
	},
}

func Test_leastBricks(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, leastBricks(q.wall), "输入 %v", q.wall)
	}
}

func Benchmark_leastBricks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			leastBricks(q.wall)
		}
	}
}
