package problem166

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para1 int
	para2 int
	ans   string
}{
	{
		2147483647,
		37,
		"58040098.(567)",
	},

	{
		10,
		3,
		"3.(3)",
	},

	{
		2,
		3,
		"0.(6)",
	},

	{
		1,
		6,
		"0.1(6)",
	},

	{
		4,
		333,
		"0.(012)",
	},

	{
		-50,
		8,
		"-6.25",
	},

	{
		-2147483648,
		1,
		"-2147483648",
	},

	{
		1,
		2,
		"0.5",
	},

	{
		2,
		1,
		"2",
	},
}

func Test_fractionToDecimal(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(fractionToDecimal(q.para1, q.para2), q.ans, "输入 %d, %d", q.para1, q.para2)
	}
}

func Benchmark_fractionToDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			fractionToDecimal(q.para1, q.para2)
		}
	}
}
