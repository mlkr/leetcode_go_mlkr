package problem400

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, ans int
}{
	{
		10000, 7,
	},

	{
		1000, 3,
	},

	{
		10, 1,
	},

	{
		3, 3,
	},

	{
		11, 0,
	},

	// 9101112
	// 100 101 102
}

func Test_findNthDigit(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findNthDigit(q.n))
	}
}

func Benchmark_findNthDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findNthDigit(q.n)
		}
	}
}
