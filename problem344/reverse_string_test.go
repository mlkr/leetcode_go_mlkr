package problem344

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s, ans []byte
}{
	{
		[]byte{'h', 'e', 'l', 'l', 'o'},
		[]byte{'o', 'l', 'l', 'e', 'h'},
	},

	{
		[]byte{'H', 'a', 'n', 'n', 'a', 'h'},
		[]byte{'h', 'a', 'n', 'n', 'a', 'H'},
	},
}

func Test_reverseString(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		reverseString(q.s)
		ast.Equal(q.ans, q.s)
	}
}

func Benchmark_reverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			reverseString(q.s)
		}
	}
}
