package problem611

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{2, 2, 3, 4},
		3,
	},

	{
		[]int{4, 2, 3, 4},
		4,
	},
}

func Test_triangleNumber(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, triangleNumber(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_triangleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			triangleNumber(q.nums)
		}
	}
}
