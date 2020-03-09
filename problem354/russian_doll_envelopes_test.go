package problem354

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	envelopes [][]int
	ans       int
}{
	{
		[][]int{
			[]int{5, 4},
			[]int{6, 4},
			[]int{6, 7},
			[]int{2, 3},
		},
		3,
	},

	{
		[][]int{
			[]int{5, 4},
			[]int{6, 1},
			[]int{6, 7},
			[]int{2, 3},
		},
		3,
	},

	{
		[][]int{
			[]int{5, 4},
			[]int{6, 6},
			[]int{6, 7},
			[]int{2, 3},
		},
		3,
	},
}

func Test_maxEnvelopes(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, maxEnvelopes(q.envelopes), "输入 %v", q.envelopes)
	}
}

func Benchmark_maxEnvelopes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maxEnvelopes(q.envelopes)
		}
	}
}
