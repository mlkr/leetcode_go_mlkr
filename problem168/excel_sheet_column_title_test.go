package problem168

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para int
	ans  string
}{
	{
		1,
		"A",
	},

	{
		28,
		"AB",
	},

	{
		701,
		"ZY",
	},

	{
		703,
		"AAA",
	},
}

func Test_convertToTitle(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(convertToTitle(q.para), q.ans)
	}
}

func Test_convertToTitle2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(convertToTitle2(q.para), q.ans)
	}
}

func Benchmark_convertToTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			convertToTitle(q.para)
		}
	}
}

func Benchmark_convertToTitle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			convertToTitle2(q.para)
		}
	}
}
