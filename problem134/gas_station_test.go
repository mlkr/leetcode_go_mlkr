package problem134

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	gas  []int
	cost []int
	ans  int
}{
	{
		[]int{2},
		[]int{2},
		0,
	},

	{
		[]int{1, 2, 3, 4, 5},
		[]int{3, 4, 5, 1, 2},
		3,
	},

	{
		[]int{2, 3, 4},
		[]int{3, 4, 3},
		-1,
	},
}

func Test_canCompleteCircuit(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, canCompleteCircuit(q.gas, q.cost))
	}
}

func Benchmark_canCompleteCircuit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			canCompleteCircuit(q.gas, q.cost)
		}
	}
}
