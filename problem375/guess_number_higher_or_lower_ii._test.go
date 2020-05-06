package problem375

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, amount int
}{
	{
		100,
		400,
	},

	{
		5,
		6,
	},

	{
		4,
		4,
	},

	{
		10,
		16,
	},

	{
		1,
		0,
	},
}

func Test_getMoneyAmount(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.amount, getMoneyAmount(q.n), "输入 %v", q.n)
	}
}

func Benchmark_getMoneyAmount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			getMoneyAmount(q.n)
		}
	}
}
