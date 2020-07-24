package problem421

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{3, 10, 5, 25, 2, 8},
		28,
	},

	// 00010
	// 00011
	// 00101
	// 01000
	// 01010
	// 11001

	// 11111	mask
	// 11100	max

}

func Test_findMaximumXOR(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findMaximumXOR(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_findMaximumXOR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findMaximumXOR(q.nums)
		}
	}
}
