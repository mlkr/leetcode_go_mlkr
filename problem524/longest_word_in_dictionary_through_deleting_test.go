package problem524

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s    string
	dict []string
	ans  string
}{
	{
		"bab",
		[]string{"ba", "ab", "a", "b", "a"},
		"ab",
	},

	{
		"abce",
		[]string{"abe", "abc"},
		"abc",
	},

	{
		"abpcplea",
		[]string{"ale", "apple", "monkey", "plea"},
		"apple",
	},

	{
		"abpcplea",
		[]string{"a", "b", "c"},
		"a",
	},

	{
		"d",
		[]string{"a", "b", "c", "abc"},
		"",
	},

	{
		"abpcplea",
		[]string{"ale", "apple", "abpcp", "monkey", "plea"},
		"abpcp",
	},
}

func Test_findLongestWord(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findLongestWord(q.s, q.dict), "输入为 %v %v", q.s, q.dict)
	}
}

func Benchmark_findLongestWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findLongestWord(q.s, q.dict)
		}
	}
}

func Test_findLongestWord2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findLongestWord2(q.s, q.dict), "输入为 %v %v", q.s, q.dict)
	}
}

func Benchmark_findLongestWord2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findLongestWord2(q.s, q.dict)
		}
	}
}
