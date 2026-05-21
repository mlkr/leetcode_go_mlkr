package problem575

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	candyType []int
	ans       int
}{
	{
		[]int{1, 1, 2, 2, 3, 3},
		3,
	},

	{
		[]int{1, 1, 2, 3},
		2,
	},

	{
		[]int{6, 6, 6, 6},
		1,
	},
}

func Test_distributeCandies(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, distributeCandies(q.candyType), "输入 %v", q.candyType)
	}
}

func Benchmark_distributeCandies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			distributeCandies(q.candyType)
		}
	}
}
