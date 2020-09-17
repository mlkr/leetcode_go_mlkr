package problem447

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	points [][]int
	ans    int
}{
	{
		[][]int{
			{0, 0},
			{1, 0},
			{2, 0},
		},
		2,

		// 相等的边数与可以组成的回旋镖的数量的关系
		// 1 0
		// 2 1*2
		// 3 (2+1)*2
		// 4 (3+2+1)*2
		// (n-1+...+1)*2
		// (n/2) * (n-1) * 2 = n*(n-1)

	},

	{
		[][]int{
			{0, 0},
			{1, 0},
		},
		0,
	},
}

func Test_numberOfBoomerangs(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, numberOfBoomerangs(q.points), "输入 %v", q.points)
	}
}

func Benchmark_numberOfBoomerangs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			numberOfBoomerangs(q.points)
		}
	}
}
