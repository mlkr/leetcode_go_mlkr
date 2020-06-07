package problem388

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	input string
	ans   int
}{
	{
		"a",
		0,
	},

	{
		"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
		32,
	},
}

func Test_lengthLongestPath(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, lengthLongestPath(q.input), "输入 %s", q.input)
	}
}

func Benchmark_lengthLongestPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			lengthLongestPath(q.input)
		}
	}
}
