package problem412

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n   int
	ans []string
}{
	{
		15,
		[]string{
			"1",
			"2",
			"Fizz",
			"4",
			"Buzz",
			"Fizz",
			"7",
			"8",
			"Fizz",
			"Buzz",
			"11",
			"Fizz",
			"13",
			"14",
			"FizzBuzz",
		},
	},
}

func Test_fizzBuzz(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, fizzBuzz(q.n), "输入 %v", q.n)
	}
}
