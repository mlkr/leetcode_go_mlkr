package problem477

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	// 0100
	// 0100
	// 1110
	{

		[]int{4, 14, 4},
		4,
	},

	// 0010
	// 0100
	// 1110
	{
		[]int{4, 14, 2},
		6,
	},
}

func Test_totalHammingDistance(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, totalHammingDistance(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_totalHammingDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			totalHammingDistance(q.nums)
		}
	}
}
