package problem476

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	num, ans int
}{

	{
		8,
		7,
	},

	{
		2,
		1,
	},

	{
		5,
		2,
	},

	{
		1,
		0,
	},
}

func Test_findComplement(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findComplement(q.num), "输入 %v", q.num)
	}
}

func Benchmark_findComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findComplement(q.num)
		}
	}
}
