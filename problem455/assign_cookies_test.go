package problem455

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	g, s []int
	ans  int
}{
	{
		[]int{1, 2, 3},
		[]int{1, 1},
		1,
	},

	{
		[]int{1, 2},
		[]int{1, 2, 3},
		2,
	},
}

func Test_findContentChildren(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findContentChildren(q.g, q.s), "输入 %v", q.g)
	}
}

func Benchmark_findContentChildren(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findContentChildren(q.g, q.s)
		}
	}
}
