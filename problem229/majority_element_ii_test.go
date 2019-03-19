package problem229

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para, ans []int
}{
	{
		[]int{2},
		[]int{2},
	},

	{
		[]int{2, 2},
		[]int{2},
	},

	{
		[]int{3, 2, 3},
		[]int{3},
	},

	{
		[]int{1, 1, 1, 3, 3, 2, 2, 2},
		[]int{1, 2},
	},
}

func Test_majorityElement(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, majorityElement(q.para))
	}
}

func Test_majorityElement2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, majorityElement2(q.para))
	}
}

func Benchmark_majorityElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			majorityElement(q.para)
		}
	}
}

func Benchmark_majorityElement2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			majorityElement2(q.para)
		}
	}
}
