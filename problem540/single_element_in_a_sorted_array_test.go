package problem540

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 1, 2, 3, 3, 4, 4, 8, 8},
		2,
	},

	{
		[]int{3, 3, 7, 7, 10, 11, 11},
		10,
	},
}

func Test_singleNonDuplicate(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, singleNonDuplicate(q.nums), "输入%v", q.nums)
	}
}

func Benchmark_singleNonDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			singleNonDuplicate(q.nums)
		}
	}
}
