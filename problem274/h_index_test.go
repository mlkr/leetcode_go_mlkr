package problem274

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	ans  int
}{
	{
		[]int{3, 0, 6, 1, 5},
		3,
	},
}

func Test_hIndex(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, hIndex(q.para))
	}
}

func Benchmark_hIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			hIndex(q.para)
		}
	}
}
