package problem131

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para string
	ans  [][]string
}{
	{
		"aab",
		[][]string{
			[]string{"a", "a", "b"},
			[]string{"aa", "b"},
		},
	},
}

func Test_partition(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(partition(q.para), q.ans, "输入 %v", q.para)
	}
}

func Benchmark_partition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			partition(q.para)
		}
	}
}
