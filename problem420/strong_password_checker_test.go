package problem420

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	ans int
}{
	{
		"aaaaaa1234567890123Ubefg",
		4,
	},

	{
		"23VFG",
		1,
	},

	{
		"23VFGHerew",
		0,
	},

	{
		"aaaaaaaaaaaaaaaaaaaaa",
		7,
	},

	{
		"aaaaaaaaaaaaaaaaaaaaaa",
		8,
	},

	{
		"aaaaaaaaaaaaaaaaaaaaaaa",
		9,
	},
}

func Test_strongPasswordChecker(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, strongPasswordChecker(q.s), "输入 %s", q.s)
	}
}

func Benchmark_strongPasswordChecker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			strongPasswordChecker(q.s)
		}
	}
}
