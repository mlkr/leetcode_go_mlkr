package problem137

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	ans  int
}{
	{
		[]int{2, 2, 3, 2},
		3,
	},

	{
		[]int{0, 1, 0, 1, 0, 1, 99},
		99,
	},
}

func Test_singleNumber(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(singleNumber(q.para), q.ans)
	}
}

func Test_singleNumber2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(singleNumber2(q.para), q.ans)
	}
}

func Benchmark_singleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			singleNumber(q.para)
		}
	}
}

func Benchmark_singleNumber2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			singleNumber2(q.para)
		}
	}
}
