package problem349

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums1, nums2, ans []int
}{
	{
		[]int{2, 6, 9, 1},
		[]int{7, 1},
		[]int{1},
	},

	{
		[]int{1, 2, 2, 1},
		[]int{2, 2},
		[]int{2},
	},

	{
		[]int{4, 9, 5},
		[]int{9, 4, 9, 8, 4},
		[]int{9, 4},
	},
}

func Test_intersection(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, intersection(q.nums1, q.nums2))
	}
}

func Benchmark_intersection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			intersection(q.nums1, q.nums2)
		}
	}
}
