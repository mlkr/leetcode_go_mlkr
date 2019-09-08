package problem321

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums1, nums2 []int
	k            int
	ans          []int
}{
	{
		[]int{3, 4, 6, 5},
		[]int{9, 1, 2, 5, 8, 3},
		5,
		[]int{9, 8, 6, 5, 3},
	},

	{
		[]int{6, 7},
		[]int{6, 0, 4},
		5,
		[]int{6, 7, 6, 0, 4},
	},

	{
		[]int{3, 9},
		[]int{8, 9},
		3,
		[]int{9, 8, 9},
	},

	{
		[]int{8, 9},
		[]int{8, 9},
		3,
		[]int{9, 8, 9},
	},

	{
		[]int{9, 1, 2, 5, 8, 3},
		[]int{3, 4, 6, 5},
		5,
		[]int{9, 8, 6, 5, 3},
	},

	{
		[]int{5, 6, 8},
		[]int{6, 4, 0},
		3,
		[]int{8, 6, 4},
	},
}

func Test_maxNumber(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, maxNumber(q.nums1, q.nums2, q.k))
	}
}

func Benchmark_maxNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maxNumber(q.nums1, q.nums2, q.k)
		}
	}
}
