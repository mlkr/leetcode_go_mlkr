package problem128

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	ast := assert.New(t)

	questions := []struct {
		nums   []int
		length int
	}{
		{
			[]int{4, 0, -4, -2, 2, 5, 2, 0, -8, -8, -8, -8, -1, 7, 4, 5, 5, -4, 6, 6, -3},
			// [-8 -8 -8 -8 -4 -4 -3 -2 -1 0 0 2 2 4 4 5 5 5 6 6 7]
			5,
		},

		{
			[]int{1, 2, 0, 1},
			// [0,1,1,2] 为连续序列, 1 只算一次
			3,
		},
		{
			[]int{0},
			1,
		},
		{
			[]int{100, 4, 200, 1, 3, 2},
			4,
		},
	}

	for _, q := range questions {
		ast.Equal(longestConsecutive(q.nums), q.length, "输入 %v", q)
		ast.Equal(longestConsecutive2(q.nums), q.length, "输入 %v", q)
		ast.Equal(longestConsecutive3(q.nums), q.length, "输入 %v", q)
	}
}
