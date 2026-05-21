package problem491

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  [][]int
}{
	{
		[]int{4, 6, 7, 7},
		[][]int{
			[]int{4, 6},
			[]int{4, 6, 7},
			[]int{4, 6, 7, 7},
			[]int{4, 7},
			[]int{4, 7, 7},
			[]int{6, 7},
			[]int{6, 7, 7},
			[]int{7, 7},
		},
	},

	{
		[]int{4, 4, 3, 2, 1},
		[][]int{
			[]int{4, 4},
		},
	},
}

func Test_findSubsequences(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, findSubsequences(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_findSubsequences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findSubsequences(q.nums)
		}
	}
}
