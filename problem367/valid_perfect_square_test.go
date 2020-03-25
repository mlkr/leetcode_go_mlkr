package problem367

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	num int
	ans bool
}{
	{
		195,
		false,
	},

	{
		16,
		true,
	},

	{
		14,
		false,
	},
}

func Test_isPerfectSquare(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, isPerfectSquare(q.num), "输入 %v", q.num)
	}
}

func Test_isPerfectSquare2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, isPerfectSquare2(q.num), "输入 %v", q.num)
	}
}

func Benchmark_isPerfectSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isPerfectSquare(q.num)
		}
	}
}

func Benchmark_isPerfectSquare2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isPerfectSquare2(q.num)
		}
	}
}
