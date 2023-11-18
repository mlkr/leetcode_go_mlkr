package problem537

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	num1, num2, ans string
}{
	{
		"1+1i",
		"1+1i",
		"0+2i",
	},

	{
		"1+-1i",
		"1+-1i",
		"0+-2i",
	},
}

func Test_complexNumberMultiply(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, complexNumberMultiply(q.num1, q.num2))
	}
}

func Benchmark_complexNumberMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			complexNumberMultiply(q.num1, q.num2)
		}
	}
}
