package problem600

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	num, ans int
}{
	{
		4,
		4,
	},

	{
		5,
		5,
	},

	{
		1,
		2,
	},

	{
		2,
		3,
	},
}

func Test_findIntegers(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findIntegers(q.num), "输入 %v", q.num)
	}
}

func Benchmark_findIntegers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findIntegers(q.num)
		}
	}
}
