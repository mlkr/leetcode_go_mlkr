package problem217

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	input []int
	ans   bool
}{
	{
		[]int{1, 2, 3, 1},
		true,
	},

	{
		[]int{1, 2, 3, 4},
		false,
	},

	{
		[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		true,
	},
}

func Test_containsDuplicate(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(containsDuplicate(q.input), q.ans, "输入 %v", q.input)
	}
}

func Benchmark_containsDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			containsDuplicate(q.input)
		}
	}
}
