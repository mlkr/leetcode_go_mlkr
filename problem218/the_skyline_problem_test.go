package problem218

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para, ans [][]int
}{
	{
		[][]int{
			[]int{2, 9, 10},
			[]int{3, 7, 15},
			[]int{5, 12, 12},
			[]int{15, 20, 10},
			[]int{19, 24, 8},
		},
		[][]int{
			[]int{2, 10},
			[]int{3, 15},
			[]int{7, 12},
			[]int{12, 0},
			[]int{15, 10},
			[]int{20, 8},
			[]int{24, 0},
		},
	},
}

func Test_getSkyline(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(getSkyline(q.para), q.ans, "输入 %v", q.para)
	}
}

func Benchmark_getSkyline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			getSkyline(q.para)
		}
	}
}
