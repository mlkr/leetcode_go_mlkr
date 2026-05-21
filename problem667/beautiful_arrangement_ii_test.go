package problem667

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, k int
	ans  []int
}{
	{
		3,
		1,
		[]int{1, 2, 3},
	},

	{
		3,
		2,
		[]int{3, 1, 2},
	},
}

func Test_constructArray(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, constructArray(q.n, q.k))
	}
}

func Benchmark_constructArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			constructArray(q.n, q.k)
		}
	}
}
