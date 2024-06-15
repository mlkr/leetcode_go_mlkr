package problem567

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s1, s2 string
	ans    bool
}{

	{
		"abc",
		"ac",
		false,
	},

	{
		"abc",
		"acb",
		true,
	},

	{
		"abc",
		"bbbca",
		true,
	},

	{
		"ab",
		"eidbaooo",
		true,
	},

	{
		"ab",
		"eidboaoo",
		false,
	},
}

func Test_checkInclusion(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, checkInclusion(q.s1, q.s2), "输入 %v", q.s1)
	}
}

func Benchmark_checkInclusion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			checkInclusion(q.s1, q.s2)
		}
	}
}
