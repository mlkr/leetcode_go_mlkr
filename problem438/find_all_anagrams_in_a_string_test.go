package problem438

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s, p string
	ans  []int
}{
	{
		"dcba",
		"abc",
		[]int{1},
	},

	{
		"cbaebabacd",
		"abc",
		[]int{0, 6},
	},

	{
		"abab",
		"ab",
		[]int{0, 1, 2},
	},
}

func Test_findAnagrams(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findAnagrams(q.s, q.p), "%v", q.s)
	}
}

func Test_findAnagrams2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findAnagrams2(q.s, q.p), "%v", q.s)
	}
}

func Benchmark_findAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findAnagrams(q.s, q.p)
		}
	}
}

func Benchmark_findAnagrams2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findAnagrams2(q.s, q.p)
		}
	}
}
