package problem373

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums1, nums2 []int
	k            int
	ans          [][]int
}{
	{
		[]int{1, 7, 11},
		[]int{2, 4, 6},
		3,
		[][]int{
			[]int{1, 2},
			[]int{1, 4},
			[]int{1, 6},
		},
	},

	{
		[]int{1, 1, 2},
		[]int{1, 2, 3},
		2,
		[][]int{
			[]int{1, 1},
			[]int{1, 1},
		},
	},

	{
		[]int{1, 2},
		[]int{3},
		3,
		[][]int{
			[]int{1, 3},
			[]int{2, 3},
		},
	},

	{
		[]int{},
		[]int{},
		0,
		[][]int{},
	},
}

func Test_kSmallestPairs(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, kSmallestPairs(q.nums1, q.nums2, q.k))
	}
}

func Benchmark_kSmallestPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			kSmallestPairs(q.nums1, q.nums2, q.k)
		}
	}
}
