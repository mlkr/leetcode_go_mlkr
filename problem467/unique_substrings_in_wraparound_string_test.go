package problem467

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	p   string
	ans int
}{
	{
		"a",
		1,
	},

	{
		"cac",
		2,
	},

	{
		"zab",
		6,
	},

	{
		"zabb",
		6,
	},
}

func Test_findSubstringInWraproundString(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findSubstringInWraproundString(q.p), "输入 %v", q.p)
	}
}

func Benchmark_findSubstringInWraproundString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findSubstringInWraproundString(q.p)
		}
	}
}
