package problem638

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	price   []int
	special [][]int
	needs   []int
	ans     int
}{
	{
		[]int{2, 5},
		[][]int{
			[]int{3, 0, 5},
			[]int{1, 2, 10},
		},
		[]int{3, 2},
		14,
	},

	{
		[]int{2, 3, 4},
		[][]int{
			[]int{1, 1, 0, 4},
			[]int{2, 2, 1, 9},
		},
		[]int{1, 2, 1},
		11,
	},
}

func Test_shoppingOffers(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, shoppingOffers(q.price, q.special, q.needs))
	}
}

func Benchmark_shoppingOffers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			shoppingOffers(q.price, q.special, q.needs)
		}
	}
}
