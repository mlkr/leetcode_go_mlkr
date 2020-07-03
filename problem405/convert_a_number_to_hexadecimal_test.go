package problem405

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	num int
	ans string
}{
	{
		0,
		"0",
	},

	{
		26,
		"1a",
	},

	{
		-1,
		"ffffffff",
	},
}

func Test_toHex(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, toHex(q.num), "输入 %v", q.num)
	}
}

func Test_toHex2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, toHex2(q.num), "输入 %v", q.num)
	}
}

func Benchmark_toHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			toHex(q.num)
		}
	}
}

func Benchmark_toHex2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			toHex2(q.num)
		}
	}
}
