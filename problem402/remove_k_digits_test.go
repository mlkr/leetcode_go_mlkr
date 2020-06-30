package problem402

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	num string
	k   int
	ans string
}{
	{
		"4321",
		2,
		"21",
	},

	{
		"1111111",
		3,
		"1111",
	},

	{
		"1432219",
		3,
		"1219",
	},

	{
		"10200",
		1,
		"200",
	},

	{
		"10200",
		2,
		"0",
	},

	{
		"10",
		2,
		"0",
	},
}

func Test_removeKdigits(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, removeKdigits(q.num, q.k), "输入 %s", q.num)
	}
}

func Benchmark_removeKdigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			removeKdigits(q.num, q.k)
		}
	}
}
