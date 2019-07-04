package problem306

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	input string
	ans   bool
}{
	{
		"000",
		true,
	},

	{
		"12012122436",
		true,
	},

	{
		"121224036",
		false,
	},

	{
		"120122436",
		false,
	},

	{
		"199111992",
		true,
	},

	{
		"101",
		true,
	},

	{
		"0235813",
		false,
	},

	{
		"211738",
		true,
	},

	{
		"1023",
		false,
	},

	{
		"11",
		false,
	},

	{
		"192358",
		false,
	},

	{
		"112358",
		true,
	},

	{
		"199100199",
		true,
	},
}

func Test_isAdditiveNumber(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, isAdditiveNumber(q.input), "输入 %s", q.input)
	}
}

func Test_isAdditiveNumber2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, isAdditiveNumber2(q.input), "输入 %s", q.input)
	}
}

func Benchmark_isAdditiveNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isAdditiveNumber(q.input)
		}
	}
}

func Benchmark_isAdditiveNumber2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isAdditiveNumber2(q.input)
		}
	}
}
