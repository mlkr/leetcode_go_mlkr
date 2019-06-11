package problem299

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	secret, guess, ans string
}{
	{
		"00112233445566778899",
		"16872590340158679432",
		"3A17B",
	},

	{
		"1122",
		"1222",
		"3A0B",
	},

	{
		"1122",
		"2211",
		"0A4B",
	},

	{
		"1807",
		"7810",
		"1A3B",
	},

	{
		"1123",
		"0111",
		"1A1B",
	},
}

func Test_getHint(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, getHint(q.secret, q.guess))
	}
}

func Test_getHint2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, getHint2(q.secret, q.guess))
	}
}

func Benchmark_getHint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			getHint(q.secret, q.guess)
		}
	}
}

func Benchmark_getHint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			getHint2(q.secret, q.guess)
		}
	}
}
