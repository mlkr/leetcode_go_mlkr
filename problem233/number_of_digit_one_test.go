package problem233

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para, ans int
}{
	{
		13,
		6,
	},
}

func Test_countDigitOne(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, countDigitOne(q.para))
	}
}

func Benchmark_countDigitOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			countDigitOne(q.para)
		}
	}
}
