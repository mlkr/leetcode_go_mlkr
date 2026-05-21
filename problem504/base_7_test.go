package problem504

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
		100,
		"202",
	},

	{
		-7,
		"-10",
	},
}

func Test_convertToBase7(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, convertToBase7(q.num), "输入 %v", q.num)
	}
}

func Benchmark_convertToBase7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			convertToBase7(q.num)
		}
	}
}
