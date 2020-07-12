package problem413

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	A   []int
	ans int
}{
	{
		[]int{1, 2, 3, 4},
		3,
	},

	{
		[]int{1, 2, 3, 4, 6},
		3,
	},

	{
		[]int{1},
		0,
	},
}

func Test_numberOfArithmeticSlices(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, numberOfArithmeticSlices(q.A), "输入 %v", q.A)
	}
}

func Benchmark_numberOfArithmeticSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			numberOfArithmeticSlices(q.A)
		}
	}
}
