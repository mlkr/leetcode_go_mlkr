package problem496

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums1, nums2, ans []int
}{
	{
		[]int{4, 1, 2},
		[]int{2, 1, 3, 4},
		[]int{-1, 3, 3},
	},

	{
		[]int{4, 1, 2},
		[]int{1, 3, 4, 2},
		[]int{-1, 3, -1},
	},

	{
		[]int{2, 4},
		[]int{1, 2, 3, 4},
		[]int{3, -1},
	},
}

func Test_nextGreaterElement(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, nextGreaterElement(q.nums1, q.nums2), "输入 %v", q.nums1)
	}
}

func Benchmark_nextGreaterElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			nextGreaterElement(q.nums1, q.nums2)
		}
	}
}
