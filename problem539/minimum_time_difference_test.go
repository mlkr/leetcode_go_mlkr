package problem539

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	timePoints []string
	ans        int
}{
	{
		[]string{"05:31", "22:08", "00:35"},
		147,
	},

	{
		[]string{"01:01", "02:01"},
		60,
	},

	{
		[]string{"23:59", "00:00"},
		1,
	},

	{
		[]string{"00:00", "23:59", "00:00"},
		0,
	},
}

func Test_findMinDifference(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findMinDifference(q.timePoints), "输入 %v", q.timePoints)
	}
}

func Benchmark_findMinDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findMinDifference(q.timePoints)
		}
	}
}
