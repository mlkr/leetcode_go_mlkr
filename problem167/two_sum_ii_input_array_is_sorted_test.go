package problem167

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	numbers []int
	target  int
	ans     []int
}{
	{
		[]int{2, 7, 11, 15},
		9,
		[]int{1, 2},
	},
}

func Test_twoSum(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(twoSum(q.numbers, q.target), q.ans)
	}
}

func Test_twoSum2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(twoSum2(q.numbers, q.target), q.ans)
	}
}

func Benchmark_twoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			twoSum(q.numbers, q.target)
		}
	}
}

func Benchmark_twoSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			twoSum2(q.numbers, q.target)
		}
	}
}
