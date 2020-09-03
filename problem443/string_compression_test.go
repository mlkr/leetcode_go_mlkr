package problem443

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	chars      []byte
	ans        int
	ansInPlace []byte
}{
	{
		[]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'},
		6,
		[]byte{'a', '3', 'b', '2', 'a', '2'},
	},

	{
		[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
		6,
		[]byte{'a', '2', 'b', '2', 'c', '3'},
	},

	{
		[]byte{'a'},
		1,
		[]byte{'a'},
	},

	{
		[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
		4,
		[]byte{'a', 'b', '1', '2'},
	},
}

func Test_compress(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		size := compress(q.chars)
		ast.Equal(q.ansInPlace, q.chars[:size], "输入 %s", q.chars)
	}
}

func Benchmark_compress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			compress(q.chars)
		}
	}
}
