package problem202

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para int
	ans  bool
}{
	{
		10,
		true,
	},

	{
		7,
		true,
	},

	{
		19,
		true,
	},
}

func Test_isHappy(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(isHappy(q.para), q.ans, "输入 %d", q.para)
	}
}

func Benchmark_isHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isHappy(q.para)
		}
	}
}
