package problem313

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n      int
	primes []int
	ans    int
}{
	{
		12,
		[]int{
			2, 7, 13, 19,
		},
		32,
	},

	{
		1,
		[]int{2},
		1,
	},
}

func Test_nthSuperUglyNumber(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, nthSuperUglyNumber(q.n, q.primes))
	}
}

func Benchmark_nthSuperUglyNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			nthSuperUglyNumber(q.n, q.primes)
		}
	}
}
