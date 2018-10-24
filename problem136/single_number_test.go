package problem136

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	ans  int
}{
	{
		[]int{2, 2, 1},
		1,
	},

	{
		[]int{4, 1, 2, 1, 2},
		4,
	},
}

func Test_singleNumber(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(singleNumber(q.para), q.ans)
	}
}

func Benchmark_singleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			singleNumber(q.para)
		}
	}
}
