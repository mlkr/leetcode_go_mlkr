package problem121

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
			[]int{7, 3, 5, 1, 6, 4},
			5,
		},
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
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
