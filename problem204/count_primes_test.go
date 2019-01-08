package problem204

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para int
	ans  int
}{
	{
		10,
		4,
	},
}

func Test_countPrimes(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(countPrimes(q.para), q.ans)
	}
}

func Benchmark_countPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			countPrimes(q.para)
		}
	}
}
