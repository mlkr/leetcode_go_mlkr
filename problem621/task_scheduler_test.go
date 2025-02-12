package problem621

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	tasks []byte
	n     int
	ans   int
}{
	{
		[]byte{'A', 'A', 'A', 'B', 'B', 'B'},
		2,
		8,
	},

	{
		[]byte{'A', 'C', 'A', 'B', 'D', 'B'},
		1,
		6,
	},

	{
		[]byte{'A', 'A', 'A', 'B', 'B', 'B'},
		3,
		10,
	},
}

func Test_leastInterval(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, leastInterval(q.tasks, q.n), "输入 %v, %v", q.tasks, q.n)
	}
}

func Benchmark_leastInterval(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			leastInterval(q.tasks, q.n)
		}
	}
}
