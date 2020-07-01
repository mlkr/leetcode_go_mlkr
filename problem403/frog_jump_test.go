package problem403

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	stones []int
	ans    bool
}{
	{

		[]int{0, 1, 3, 6, 10, 15, 19, 21, 24, 26, 30, 33},
		true,
	},

	{
		[]int{0, 1, 3, 5, 6, 8, 12, 17},
		true,
	},

	{
		[]int{0, 1, 2, 3, 4, 8, 9, 11},
		false,
	},

	{
		[]int{0, 1},
		true,
	},
}

func Test_canCross(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, canCross(q.stones), "输入 %d", q.stones)
	}
}

func Benchmark_canCross(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			canCross(q.stones)
		}
	}
}
