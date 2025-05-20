package problem636

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n    int
	logs []string
	ans  []int
}{
	{
		2,
		[]string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"},
		[]int{3, 4},
	},

	{
		1,
		[]string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"},
		[]int{8},
	},

	{
		2,
		[]string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"},
		[]int{7, 1},
	},
}

func Test_exclusiveTime(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, exclusiveTime(q.n, q.logs))
	}
}

func Benchmark_exclusiveTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			exclusiveTime(q.n, q.logs)
		}
	}
}
