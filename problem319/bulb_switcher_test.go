package problem319

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, ans int
}{
	{
		0,
		0,
	},

	{
		1,
		1,
	},

	{
		2,
		1,
	},

	{
		3,
		1,
	},

	{
		4,
		2,
	},

	{
		5,
		2,
	},
}

func Test_bulbSwitch(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, bulbSwitch(q.n), "输入 %v", q.n)
	}
}

func Benchmark_bulbSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			bulbSwitch(q.n)
		}
	}
}
