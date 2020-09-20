package problem448

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums, ans []int
}{
	{
		[]int{4, 3, 2, 7, 8, 2, 3, 1},
		[]int{5, 6},
	},
}

func Test_findDisappearedNumbers(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		nums := make([]int, len(q.nums))
		copy(nums, q.nums)
		ast.Equal(q.ans, findDisappearedNumbers(nums), "输入 %v", q.nums)
	}
}

func Benchmark_findDisappearedNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			nums := make([]int, len(q.nums))
			copy(nums, q.nums)
			findDisappearedNumbers(nums)
		}
	}
}
