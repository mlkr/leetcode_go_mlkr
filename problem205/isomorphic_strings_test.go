package problem205

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s, t string
	ans  bool
}{
	{
		"egg",
		"add",
		true,
	},

	{
		"foo",
		"bar",
		false,
	},

	{
		"paper",
		"title",
		true,
	},
}

func Test_isIsomorphic(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(isIsomorphic(q.s, q.t), q.ans, "输入 %s", q.s)
	}
}

func Test_isIsomorphic2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(isIsomorphic2(q.s, q.t), q.ans, "输入 %s", q.s)
	}
}

func Benchmark_isIsomorphic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isIsomorphic(q.s, q.t)
		}
	}
}

func Benchmark_isIsomorphic2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isIsomorphic2(q.s, q.t)
		}
	}
}
