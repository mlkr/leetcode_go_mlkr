package problem258

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para, ans int
}{
	{
		38,
		2,
	},
}

func Test_addDigits(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, addDigits(q.para))
	}
}

func Test_addDigits2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, addDigits2(q.para))
	}
}

func Benchmark_addDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			addDigits(q.para)
		}
	}
}

func Benchmark_addDigits2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			addDigits2(q.para)
		}
	}
}
