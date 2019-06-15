package problem300

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 3, 6, 7, 9, 4, 10, 5, 6},
		6,
	},

	{
		[]int{},
		0,
	},

	{
		[]int{10, 9, 2, 5, 3, 7, 101, 18},
		4,
	},
}

func Test_lengthOfLIS(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, lengthOfLIS(q.nums))
	}
}

func Benchmark_lengthOfLIS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			lengthOfLIS(q.nums)
		}
	}
}
