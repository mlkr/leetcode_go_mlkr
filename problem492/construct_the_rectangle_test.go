package problem492

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	area int
	ans  []int
}{
	{
		2,
		[]int{2, 1},
	},

	{
		4,
		[]int{2, 2},
	},

	{
		37,
		[]int{37, 1},
	},

	{
		122122,
		[]int{427, 286},
	},
}

func Test_constructRectangle(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, constructRectangle(q.area), "输入 %v", q.area)
	}
}

func Benchmark_constructRectangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			constructRectangle(q.area)
		}
	}
}
