package problem169

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	ans  int
}{
	{
		[]int{3, 2, 3},
		3,
	},

	{
		[]int{2, 2, 1, 1, 1, 2, 2},
		2,
	},
}

func Test_majorityElement(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(majorityElement(q.para), q.ans)
	}
}

func Test_majorityElement2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(majorityElement2(q.para), q.ans)
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
