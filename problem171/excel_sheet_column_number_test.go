package problem171

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para string
	ans  int
}{
	{
		"A",
		1,
	},

	{
		"AB",
		28,
	},

	{
		"ZY",
		701,
	},
}

func Test_titleToNumber(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(titleToNumber(q.para), q.ans)
	}
}

func Benchmark_titleToNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			titleToNumber(q.para)
		}
	}
}
