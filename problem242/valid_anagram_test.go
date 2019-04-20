package problem242

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s, t string
	ans  bool
}{
	{
		"aacc",
		"cca",
		false,
	},

	{
		"aacc",
		"ccac",
		false,
	},

	{
		"anagram",
		"nagaram",
		true,
	},

	{
		"rat",
		"car",
		false,
	},
}

func Test_isAnagram(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, isAnagram(q.s, q.t), "输入 %v", q.s)
	}
}

func Benchmark_isAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isAnagram(q.s, q.t)
		}
	}
}
