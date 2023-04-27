package problem507

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	num int
	ans bool
}{
	{
		1,
		false,
	},

	{
		28,
		true,
	},

	{
		7,
		false,
	},
}

func Test_checkPerfectNumber(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, checkPerfectNumber(q.num), "输入 %v", q.num)
	}
}

func Benchmark_checkPerfectNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			checkPerfectNumber(q.num)
		}
	}
}
