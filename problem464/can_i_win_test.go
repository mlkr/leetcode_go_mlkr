package problem464

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	maxChoosableInteger, desiredTotal int
	ans                               bool
}{
	{
		10,
		11,
		false,
	},

	{
		6,
		11,
		true,
	},

	{
		10,
		0,
		true,
	},

	{
		10,
		1,
		true,
	},

	{
		20,
		160,
		true,
	},

	{
		20,
		1600,
		false,
	},
}

func Test_canIWin(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans,
			canIWin(q.maxChoosableInteger, q.desiredTotal),
			"输入 %v %v", q.maxChoosableInteger, q.desiredTotal)
	}
}

func Benchmark_canIWin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			canIWin(q.maxChoosableInteger, q.desiredTotal)
		}
	}
}
