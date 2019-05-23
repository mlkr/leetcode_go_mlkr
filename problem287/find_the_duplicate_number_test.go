package problem287

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 3, 4, 2, 2},
		2,
	},

	{
		[]int{3, 1, 3, 4, 2},
		3,
	},
}

func Test_findDuplicate(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findDuplicate(q.nums))
	}
}

func Benchmark_findDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findDuplicate(q.nums)
		}
	}
}
