package problem518

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	amount int
	coins  []int
	ans    int
}{
	{
		5,
		[]int{1, 2, 5},
		4,
	},

	{
		3,
		[]int{2},
		0,
	},

	{
		10,
		[]int{10},
		1,
	},
}

func Test_change(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, change(q.amount, q.coins), "输入 %v %v", q.amount, q.coins)
	}
}

func Benchmark_change(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			change(q.amount, q.coins)
		}
	}
}
