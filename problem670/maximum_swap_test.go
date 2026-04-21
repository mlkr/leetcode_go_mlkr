package problem670

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	num, ans int
}{
	{
		99088,
		99880,
	},

	{
		98368,
		98863,
	},

	{
		1993,
		9913,
	},

	{
		11,
		11,
	},

	{
		2736,
		7236,
	},

	{
		9973,
		9973,
	},

	{
		1234,
		4231,
	},

	{
		4231,
		4321,
	},

	{
		1023,
		3021,
	},

	{
		3021,
		3201,
	},

	{
		3201,
		3210,
	},
}

func Test_maximumSwap(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, maximumSwap(q.num), "输入 %v", q.num)
	}
}

func Benchmark_maximumSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maximumSwap(q.num)
		}
	}
}
