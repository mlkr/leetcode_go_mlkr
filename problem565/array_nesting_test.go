package problem565

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{5, 4, 0, 3, 1, 6, 2},
		4,
	},

	{
		[]int{0, 1, 2},
		1,
	},
}

func Test_arrayNesting(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, arrayNesting(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_arrayNesting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			arrayNesting(q.nums)
		}
	}
}
