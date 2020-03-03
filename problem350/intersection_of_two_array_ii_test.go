package problem350

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums1, nums2, ans []int
}{
	{
		[]int{1, 2, 2, 1},
		[]int{2, 2},
		[]int{2, 2},
	},

	{
		[]int{4, 9, 5},
		[]int{9, 4, 9, 8, 4},
		[]int{4, 9},
	},
}

func Test_intersect(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, intersect(q.nums1, q.nums2), "输入 %q %q", q.nums1, q.nums2)
	}
}

func Benchmark_intersect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			intersect(q.nums1, q.nums2)
		}
	}
}
