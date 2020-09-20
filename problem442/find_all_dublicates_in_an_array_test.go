package problem442

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums, ans []int
}{
	{
		[]int{4, 3, 2, 7, 8, 2, 3, 1},
		[]int{2, 3},
	},
}

func Test_findDuplicates(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		nums := make([]int, len(q.nums))
		copy(nums, q.nums)
		ast.Equal(q.ans, findDuplicates(nums), "输入 %v", q.nums)
	}
}

func Test_findDuplicates2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		nums := make([]int, len(q.nums))
		copy(nums, q.nums)
		ast.Equal(q.ans, findDuplicates2(nums), "输入 %v", q.nums)
	}
}

func Benchmark_findDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			nums := make([]int, len(q.nums))
			copy(nums, q.nums)
			findDuplicates(nums)
		}
	}
}

func Benchmark_findDuplicates2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			nums := make([]int, len(q.nums))
			copy(nums, q.nums)
			findDuplicates2(nums)
		}
	}
}
