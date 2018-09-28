package problem123

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	ast := assert.New(t)

	questions := []struct {
		prices []int
		profit int
	}{
		{
			[]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
			// 3,-2,5,-5,7,-9
			// 3,1,6,1,8,0
			// 3,3,6,6,8,8

			// 8,5,7,2,7,0
			// 8,7,7,7,7,0

			// 0,7,2,7,5,8
			// 8,5,7,2,7,0
			// 0,3,1,6,8,0
			//
			// 3,3,6,6,8,8
			//
			13,
		},

		{
			[]int{3, 3, 5, 0, 0, 3, 1, 4},
			6,
		},

		{
			[]int{1, 2, 3, 4, 5},
			4,
		},

		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	for _, q := range questions {
		ast.Equal(maxProfit(q.prices), q.profit, "输入 %v", q.prices)
		ast.Equal(maxProfit2(q.prices), q.profit, "输入 %v", q.prices)
	}
}
