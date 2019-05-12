package problem275

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	ans  int
}{
	{
		[]int{0, 1},
		1,
	},

	{
		[]int{0, 1, 3, 5, 6},
		3,
	},

	{
		[]int{0, 1, 2, 5, 6},
		2,
	},
}

func Test_hIndex(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, hIndex(q.para))
	}
}

func Test_hIndex2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, hIndex2(q.para))
	}
}

func Benchmark_hIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			hIndex(q.para)
		}
	}
}

func Benchmark_hIndex2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			hIndex2(q.para)
		}
	}
}
