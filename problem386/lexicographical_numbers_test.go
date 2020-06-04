package problem386

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n   int
	ans []int
}{
	{
		13,
		[]int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
	},
}

func Test_lexicalOrder(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, lexicalOrder(q.n), "输入 %d", q.n)
	}
}

func Benchmark_lexicalOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			lexicalOrder(q.n)
		}
	}
}
