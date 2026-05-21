package problem506

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	score []int
	ans   []string
}{
	{
		[]int{5, 4, 3, 2, 1},
		[]string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
	},

	{
		[]int{10, 3, 8, 9, 4},
		[]string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
	},
}

func Test_findRelativeRanks(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findRelativeRanks(q.score), "输入 %v", q.score)
	}
}

func Benchmark_findRelativeRanks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findRelativeRanks(q.score)
		}

	}
}
