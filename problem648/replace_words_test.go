package problem648

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	dict []string
	sent string
	ans  string
}{
	{
		[]string{
			"a",
			"aa",
			"aaa",
			"aaas",
		},
		"a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa",
		"a a a a a a a a bbb baba a",
	},

	{
		[]string{
			"cat",
			"bat",
			"rat",
		},
		"the cattle was rattled by the battery",
		"the cat was rat by the bat",
	},

	{
		[]string{
			"a",
			"b",
			"c",
		},
		"aadsfasf absbs bbab cadsfafs",
		"a a b c",
	},
}

func Test_replaceWords(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, replaceWords(q.dict, q.sent))
	}
}

func Benchmark_replaceWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			replaceWords(q.dict, q.sent)
		}
	}
}
