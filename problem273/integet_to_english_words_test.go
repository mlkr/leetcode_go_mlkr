package problem273

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para int
	ans  string
}{
	{
		0,
		"Zero",
	},

	{
		1000000,
		"One Million",
	},

	{
		123,
		"One Hundred Twenty Three",
	},

	{
		12345,
		"Twelve Thousand Three Hundred Forty Five",
	},

	{
		1234567,
		"One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
	},

	{
		1234567891,
		"One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One",
	},
}

func Test_numberToWords(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(numberToWords(q.para), q.ans)
	}
}

func Benchmark_numberToWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			numberToWords(q.para)
		}
	}
}
