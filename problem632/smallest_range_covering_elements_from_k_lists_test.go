package problem632

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums [][]int
	ans  []int
}{
	{
		[][]int{
			[]int{4, 10, 15, 24, 26},
			[]int{0, 9, 12, 20},
			[]int{5, 18, 22, 30},
		},

		[]int{20, 24},
	},
}

func Test_smallestRange(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(smallestRange(q.nums), q.ans, "输入 %v", q.nums)
	}
}

func Benchmark_smallestRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			smallestRange(q.nums)
		}
	}
}

func Test_smallestRange2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(smallestRange2(q.nums), q.ans, "输入 %v", q.nums)
	}
}

func Benchmark_smallestRange2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			smallestRange2(q.nums)
		}
	}
}
